package view

import (
	betagloss "github.com/charmbracelet/lipgloss/v2"
)

func Render_color_box(elems []string, cursor, width, height int) string {
	style := betagloss.NewStyle().
		Border(betagloss.RoundedBorder()).
		Height(get_colormodal_height(height)).
		Width(get_colormodal_width(width))

	items := betagloss.JoinVertical(betagloss.Left, elems...)

	return style.Render(items)
}

func Render_color_modal(output string, width, height int) string {
	layer1 := betagloss.NewLayer(output).X(0).Y(0)
	style := betagloss.NewStyle().Width(get_colormodal_width(width))
	elems := []string{style.Render("elem 1"), style.Render("elem 2"), style.Render("elem 3"), style.Render("elem 4")}
	layer2 := betagloss.NewLayer(Render_color_box(elems, 0, width, height))

	l1H := layer1.GetHeight()
	l1W := layer1.GetWidth()

	layer2 = layer2.X(l1W / 3).Y(l1H / 4).Z(1)

	canv := betagloss.NewCanvas(layer1, layer2)

	return canv.Render()

}
