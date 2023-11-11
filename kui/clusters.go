package kui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"os"
	"termk/k8i"
)

func ShowListClusters(clusters []string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	list := widgets.NewList()
	list.Title = "[ Clusters ]"
	list.Rows = clusters
	list.TextStyle = ui.NewStyle(ui.ColorYellow)
	list.WrapText = false
	info := InfoK("no selected", "no selected", "no selected")
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
		case "<Enter>":
			clusterSelected := list.Rows[list.SelectedRow]
			client := k8i.SetClient(clusterSelected, k8i.GetPathKubeConfig())
			ShowListNamespaces(client, clusterSelected)
			return
		}
		RenderList(list, info)
		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

	}
}
