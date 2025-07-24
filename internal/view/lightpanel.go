package view

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func Render_light_title(title string, bri float64, selected bool) string {
	style := lipgloss.NewStyle()
	selectedStyle := style.Background(white).Foreground(navy)
	brightness := fmt.Sprint(int(bri), `%`)

	output := apply_horizontal_limit(title, default_horizontal_limit)

	//adjusting the paddings for the brightness
	output = output[:len(output)-len(brightness)]
	output = output + brightness

	if selected {
		return selectedStyle.Render(output)
	}
	return style.Render(output)
}

func Render_light_panel(elems []string, selected bool) string {
	border := lipgloss.RoundedBorder()
	border.TopLeft = "3"

	defaultStyle := lipgloss.NewStyle().Border(border).Margin(0, 1).PaddingLeft(1)
	selectedStyle := defaultStyle.BorderForeground(cyan)

	elems = apply_vertical_limit(elems, lights_vertical_limit)

	items := lipgloss.JoinVertical(lipgloss.Left, elems...)

	if selected {
		return selectedStyle.Render(items)
	}
	return defaultStyle.Render(items)

}
