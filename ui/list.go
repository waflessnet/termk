package ui

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"k8s.io/client-go/kubernetes"
	"log"
	"os"
	"strings"
	"termk/k8i"
)

func ShowListClusters(clusters []string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	l := widgets.NewList()
	l.Title = "[ Clusters ]"
	l.Rows = clusters
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 120, 40)

	/*	vs := ViewState{
			ItemSelect: 0,
		}
		l.Rows[vs.ItemSelect] = setSelectedColor(l.Rows[vs.ItemSelect])
	*/
	ui.Render(l)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			os.Exit(0)
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		case "<Enter>":
			// println(l.Rows[l.SelectedRow])
			//ui.Close()
			clusterSelected := l.Rows[l.SelectedRow]
			client := k8i.SetClient(clusterSelected, k8i.GetPathKubeConfig())
			ShowListNamespaces(client, clusterSelected)
			return
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		ui.Render(l)
	}
}

func ShowListNamespaces(client *kubernetes.Clientset, cluster string) {
	// cluster string is used  only show menu top
	namespaces := k8i.GetNameSpaces(client)

	l := widgets.NewList()
	l.Title = "Namespaces of [" + cluster + "]"
	l.Rows = namespaces
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 120, 40)

	/*	vs := ViewState{
			ItemSelect: 0,
		}
		l.Rows[vs.ItemSelect] = setSelectedColor(l.Rows[vs.ItemSelect])
	*/
	ui.Render(l)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			os.Exit(0)
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		case "<Enter>":
			namespaceSelected := l.Rows[l.SelectedRow]
			//ui.Close()
			ShowListPods(client, cluster, namespaceSelected)
			return
			//println(pods)
		case "<Backspace>", "<Escape>":
			// go back to clusters
			clusters := k8i.GetClusters()
			ShowListClusters(clusters)
			return
		}
		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		ui.Render(l)
	}
}

func ShowListPods(client *kubernetes.Clientset, cluster string, namespace string) {
	ui.Clear()
	pods := k8i.GetPods(client, namespace)
	// println(pods)

	l := widgets.NewList()
	l.Title = "Pods of Namespaces [" + namespace + "]"
	l.Rows = pods
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 120, 40)

	/*	vs := ViewState{
			ItemSelect: 0,
		}
		l.Rows[vs.ItemSelect] = setSelectedColor(l.Rows[vs.ItemSelect])
	*/
	ui.Render(l)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			os.Exit(0)
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		case "<Enter>":
			podSelected := l.Rows[l.SelectedRow]
			podSelected = strings.Trim(strings.Split(podSelected, "->")[0], " ")

			ShowLogPod(client, cluster, namespace, podSelected, 0)
			return
			//println(pods)
		case "<Backspace>", "<Escape>":
			// go back to clusters

			ShowListNamespaces(client, cluster)
			return
		case "<C-q>":
			podSelected := l.Rows[l.SelectedRow]
			podSelected = strings.Trim(strings.Split(podSelected, "->")[0], " ")
			err := k8i.DeletePod(client, namespace, podSelected)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error deleting pod: %v\n", err)
			}
			ShowListPods(client, cluster, namespace)
			return
		case "<d>":
			podSelected := l.Rows[l.SelectedRow]
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

		ui.Render(l)
	}
}

func ShowLogPod(client *kubernetes.Clientset, cluster string, namespace string, pod string, selectLine int) {
	logs := k8i.GetLogPod(client, namespace, pod)

	var rowLog = strings.Split(logs, "\n")
	// last is empy ""
	if len(rowLog) > 1 {
		rowLog = rowLog[:len(rowLog)-1]
	}

	var last []string
	last = rowLog
	if len(rowLog) > 500 {
		last = rowLog[len(rowLog)-500:]
	}

	//rowLog = append(rowLog, logs)

	l := widgets.NewList()
	l.Title = "[" + pod + "]"
	l.Rows = last
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 120, 40)
	l.SelectedRow = selectLine
	if selectLine == 0 {
		l.SelectedRow = len(last) - 1
	}
	ui.Render(l)

	//l.ScrollBottom()
	previousKey := ""
	uiEvents := ui.PollEvents()
	//l.ScrollBottom()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			os.Exit(0)
			return
		case "j", "<Down>":
			l.ScrollDown()

		case "k", "<Up>":
			l.ScrollUp()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<PageUp>":
			l.ScrollHalfPageUp()
		case "<PageDown>":
			l.ScrollHalfPageDown()
		case "e":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		case "m":
			// podSelected := l.Rows[l.SelectedRow]
			ui.Clear()
			//ui.Close()
			//print("wait.. get logs")
			//time.Sleep(3 * time.Second)
			ShowLogPod(client, cluster, namespace, pod, 0)
			//print(k8i.GetLogPod(client, namespace, podSelected))
			//println(pods)
		case "<Enter>":
			logSelected := l.Rows[l.SelectedRow]
			//ui.Clear()
			//ui.Close()
			ShowCompleteLog(client, cluster, namespace, pod, logSelected, l.SelectedRow)
		case "<Backspace>", "<Escape>":
			// go back to clusters
			ShowListPods(client, cluster, namespace)
			return
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}
		//print(e.ID)

		ui.Render(l)
	}
}

func ShowCompleteLog(client *kubernetes.Clientset, cluster string, namespace string, pod string, str string, selectedLine int) {

	// Texto largo
	// longText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus pharetra mauris at erat efficitur, in sollicitudin nulla mollis. Sed tincidunt libero in dignissim fringilla. Vestibulum euismod, orci in fringilla ultricies, urna nisi fringilla elit, id lacinia quam tortor id risus."

	// Crear el párrafo con desplazamiento horizontal habilitado
	p := widgets.NewParagraph()
	p.Text = str
	p.WrapText = true
	p.SetRect(0, 0, 120, 40)
	p.TextStyle.Fg = ui.ColorWhite
	p.BorderStyle.Fg = ui.ColorCyan
	p.Title = "LOG"

	// Renderizar el párrafo
	ui.Render(p)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			os.Exit(0)
			return
		case "<Backspace>", "<Escape>":
			// go back to clusters
			//ui.Close()
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
		ui.Render(p)
	}

}

func ShowLoading() {

	g0 := widgets.NewGauge()
	g0.Title = "Slim Gauge"
	g0.SetRect(20, 20, 30, 30)
	g0.Percent = 75
	g0.BarColor = ui.ColorRed
	g0.BorderStyle.Fg = ui.ColorWhite
	g0.TitleStyle.Fg = ui.ColorCyan

	g2 := widgets.NewGauge()
	g2.Title = "Slim Gauge"
	g2.SetRect(0, 3, 50, 6)
	g2.Percent = 60
	g2.BarColor = ui.ColorYellow
	g2.LabelStyle = ui.NewStyle(ui.ColorBlue)
	g2.BorderStyle.Fg = ui.ColorWhite

	g1 := widgets.NewGauge()
	g1.Title = "Big Gauge"
	g1.SetRect(0, 6, 50, 11)
	g1.Percent = 30
	g1.BarColor = ui.ColorGreen
	g1.LabelStyle = ui.NewStyle(ui.ColorYellow)
	g1.TitleStyle.Fg = ui.ColorMagenta
	g1.BorderStyle.Fg = ui.ColorWhite

	g3 := widgets.NewGauge()
	g3.Title = "Gauge with custom label"
	g3.SetRect(0, 11, 50, 14)
	g3.Percent = 50
	g3.Label = fmt.Sprintf("%v%% (100MBs free)", g3.Percent)

	g4 := widgets.NewGauge()
	g4.Title = "Gauge"
	g4.SetRect(0, 14, 50, 17)
	g4.Percent = 50
	g4.Label = "Gauge with custom highlighted label"
	g4.BarColor = ui.ColorGreen
	g4.LabelStyle = ui.NewStyle(ui.ColorYellow)

	ui.Render(g0, g1, g2, g3, g4)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			os.Exit(0)
			return
		}
	}
}
