package kui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"k8s.io/client-go/kubernetes"
	"os"
	"termk/k8i"
)

func ShowListNamespaces(client *kubernetes.Clientset, cluster string) {
	ui.Clear()
	// cluster string is used  only show menu top
	namespaces := k8i.GetNameSpaces(client)

	list := widgets.NewList()
	list.Title = "Namespaces of [" + cluster + "]"
	list.Rows = namespaces
	list.TextStyle = ui.NewStyle(ui.ColorYellow)
	list.WrapText = false
	info := InfoK(cluster, "no selected", "no selected")
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
		case "G", "<End>":
			list.ScrollBottom()
		case "r":
			ShowListNamespaces(client, cluster)
			return
		case "<Enter>":
			namespaceSelected := list.Rows[list.SelectedRow]
			//kui.Close()
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

		RenderList(list, info)
	}
}
