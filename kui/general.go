package kui

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"time"
)

func InfoK(cluster string, namespace string, pod string) *widgets.Paragraph {
	infoK := widgets.NewParagraph()
	now := time.Now()
	infoK.WrapText = true
	infoK.Text = fmt.Sprintf(`
	Date: %s
	cluster: %s
	namespace: %s
	pod: %s
		
	keys
	g:top    G: buttom    j:Down    k:Up    esc:return    <C-d>:Page down    <C-u>:Page up    <C-q>:delete pod    q:exit
	`, now.Format("2006-01-02 15:04:05"), cluster, namespace, pod)
	infoK.Title = "Info Kubernetes"
	return infoK

}

func RenderList(list interface{}, info *widgets.Paragraph) *ui.Grid {
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		ui.NewRow(1.0/5,
			ui.NewCol(1.0/1, info),
		),
		ui.NewRow(1.0/1.4, list),
	)
	ui.Render(grid)
	return grid
}
