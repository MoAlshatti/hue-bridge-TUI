package view

import betagloss "github.com/charmbracelet/lipgloss/v2"

func render_bri_box(intputView string, valid bool, width, height int) string {
	boxStyle := betagloss.NewStyle().
		Border(betagloss.RoundedBorder()).
		Foreground(white).
		Height(get_brimodal_height(height)).
		Width(get_brimodal_width(width)).
		AlignHorizontal(betagloss.Center)
	textStyle := betagloss.NewStyle()
	validStyle := textStyle.Foreground(betagloss.Red)

	InvalidWarning := " "
	if !valid {
		InvalidWarning = validStyle.Render("Invalid Input!")
	}

	title := textStyle.Render("Enter a number between 0 and a 100!")
	output := betagloss.JoinVertical(betagloss.Center, title, InvalidWarning)
	return boxStyle.Render(betagloss.JoinVertical(betagloss.Left, output, intputView))
}

func Render_bri_modal(output, inputView string, valid bool, width, height int) string {
	layer1 := betagloss.NewLayer(output).X(0).Y(0)
	layer2 := betagloss.NewLayer(render_bri_box(inputView, valid, width, height))

	l1H := layer1.GetHeight()
	l1W := layer1.GetWidth()

	layer2 = layer2.X(l1W / 3).Y(l1H / 4).Z(1)

	canv := betagloss.NewCanvas(layer1, layer2)

	return canv.Render()
}
