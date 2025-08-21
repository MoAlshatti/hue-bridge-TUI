package view

import (
	"github.com/charmbracelet/bubbles/v2/list"
	betagloss "github.com/charmbracelet/lipgloss/v2"
)

func Render_color_modal(output, listView string, width, height int) string {

	boxStyle := betagloss.NewStyle().Border(betagloss.RoundedBorder())
	layer1 := betagloss.NewLayer(output).X(0).Y(0)
	layer2 := betagloss.NewLayer(boxStyle.Render(listView))

	l1H := layer1.GetHeight()
	l1W := layer1.GetWidth()

	layer2 = layer2.X(l1W / 3).Y(l1H / 4).Z(1)

	canv := betagloss.NewCanvas(layer1, layer2)

	return canv.Render()

}

func Apply_list_style(l *list.Model) {

	titleStyle := betagloss.NewStyle().MarginLeft(2).Bold(true)
	paginationStyle := list.DefaultStyles(true).PaginationStyle.PaddingLeft(4)

	l.Title = "Select a color from the list!"
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)
	l.SetFilteringEnabled(true)

	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle

}
func Update_list_size(l *list.Model, width, height int) {
	l.SetHeight(Get_colormodal_height(height))
	l.SetWidth(width)
	l.FilterInput.SetWidth(Get_colormodal_width(width))
}
