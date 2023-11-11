package kui

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"k8s.io/client-go/kubernetes"
	"os"
	"strings"
	"termk/k8i"
)

func ShowListPods(client *kubernetes.Clientset, cluster string, namespace string) {
	ui.Clear()
	pods := k8i.GetPods(client, namespace)
	// println(pods)

	list := widgets.NewList()
	list.Title = "Pods of Namespaces [" + namespace + "]"
	list.Rows = pods
	list.TextStyle = ui.NewStyle(ui.ColorYellow)
	list.WrapText = false
	info := InfoK(cluster, namespace, "no selected")
	RenderList(list, info)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			ui.Clear()
			ui.Close()
			os.Exit(0)
			return
		case "j", "<Down>":
			list.ScrollDown()
		case "k", "<Up>":
			list.ScrollUp()
		case "<C-d>":
			list.ScrollHalfPageDown()
		case "<C-u>":
			list.ScrollHalfPageUp()
		case "<C-f>":
			list.ScrollPageDown()
		case "<C-b>":
			list.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				list.ScrollTop()
			}
		case "<Home>":
			list.ScrollTop()
		case "r":
			ShowListPods(client, cluster, namespace)
			return
		case "G", "<End>":
			list.ScrollBottom()
		case "<Enter>":
			podSelected := list.Rows[list.SelectedRow]
			podSelected = strings.Trim(strings.Split(podSelected, "->")[0], " ")

			ShowLogPod(client, cluster, namespace, podSelected, 0)
			return
		case "<Backspace>", "<Escape>":
			ShowListNamespaces(client, cluster)
			return
		case "<C-q>":
			podSelected := list.Rows[list.SelectedRow]
			podSelected = strings.Trim(strings.Split(podSelected, "->")[0], " ")
			err := k8i.DeletePod(client, namespace, podSelected)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error deleting pod: %v\n", err)
			}
			ShowListPods(client, cluster, namespace)
			return
		case "<d>":
			podSelected := list.Rows[list.SelectedRow]
			podSelected = strings.Trim(strings.Split(podSelected, "->")[0], " ")
			err := k8i.DeletePod(client, namespace, podSelected)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error deleting pod: %v\n", err)
			}
			ShowListPods(client, cluster, namespace)
			return
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}
		RenderList(list, info)
	}
}

func ShowCompleteLog(client *kubernetes.Clientset, cluster string, namespace string, pod string, str string, selectedLine int) {
	ui.Clear()

	p := widgets.NewParagraph()
	p.Text = str
	p.WrapText = true
	p.SetRect(0, 0, 120, 40)
	p.TextStyle.Fg = ui.ColorWhite
	p.BorderStyle.Fg = ui.ColorCyan
	p.Title = "LOG"

	info := InfoK(cluster, namespace, pod)
	RenderList(p, info)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			ui.Clear()
			ui.Close()
			os.Exit(0)
			return
		case "<Backspace>", "<Escape>":
			ShowLogPod(client, cluster, namespace, pod, selectedLine)
			return
		case "g":
			if previousKey == "g" {
				//p.Title
			}
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}
		info := InfoK(cluster, namespace, pod)
		RenderList(p, info)
	}

}

func ShowLogPod(client *kubernetes.Clientset, cluster string, namespace string, pod string, selectLine int) {
	ui.Clear()
	logs := k8i.GetLogPod(client, namespace, pod)
	var rowLog = strings.Split(logs, "\n")
	if len(rowLog) > 1 {
		rowLog = rowLog[:len(rowLog)-1]
	}

	var last []string
	last = rowLog
	if len(rowLog) > 500 {
		last = rowLog[len(rowLog)-500:]
	}

	list := widgets.NewList()
	list.Title = "[" + pod + "]"
	list.Rows = last
	list.TextStyle = ui.NewStyle(ui.ColorYellow)
	list.WrapText = false
	list.SelectedRow = selectLine
	if selectLine == 0 {
		list.SelectedRow = len(last) - 1
	}
	//
	info := InfoK(cluster, namespace, pod)
	RenderList(list, info)
	previousKey := ""
	uiEvents := ui.PollEvents()
	//list.ScrollBottom()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			ui.Clear()
			ui.Close()
			os.Exit(0)
			return
		case "j", "<Down>":
			list.ScrollDown()

		case "k", "<Up>":
			list.ScrollUp()
		case "<C-d>":
			list.ScrollHalfPageDown()
		case "<C-u>":
			list.ScrollHalfPageUp()
		case "<PageUp>":
			list.ScrollHalfPageUp()
		case "<PageDown>":
			list.ScrollHalfPageDown()
		case "e":
			list.ScrollPageDown()
		case "<C-b>":
			list.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				list.ScrollTop()
			}
		case "<Home>":
			list.ScrollTop()
		case "G", "<End>":
			list.ScrollBottom()
		case "r":
			ShowLogPod(client, cluster, namespace, pod, 0)
		case "<Enter>":
			logSelected := list.Rows[list.SelectedRow]
			ShowCompleteLog(client, cluster, namespace, pod, logSelected, list.SelectedRow)
		case "<Backspace>", "<Escape>":
			ShowListPods(client, cluster, namespace)
			return
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}
		RenderList(list, info)
	}
}
