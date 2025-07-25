package view

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func Render_light_title(title string, bri float64, on bool, selected bool) string {
	style := lipgloss.NewStyle()
	selectedStyle := style.Background(white).Foreground(navy)

	brightness := fmt.Sprint(int(bri), `% `)
	status := "OFF "

	output := apply_horizontal_limit(title, default_horizontal_limit)

	if !on {
		output = output[:len(output)-len(status)]
		output = output + status
	} else {
		//adjusting the paddings for the brightness
		output = output[:len(output)-len(brightness)]
		output = output + brightness

	}

	if selected {
		return selectedStyle.Render(output)
	}
	return style.Render(output)
}

func Render_light_panel(elems []string, selected bool, cursor int) string {
	border := lipgloss.RoundedBorder()
	border.TopLeft = "3"

	defaultStyle := lipgloss.NewStyle().Border(border).Margin(0, 1).PaddingLeft(1)
	selectedStyle := defaultStyle.BorderForeground(cyan)

	if len(elems) > max_lights_page_size {
		pageSize := min(max_lights_page_size, len(elems))
		if cursor%pageSize == 0 {
			elems = elems[cursor : cursor+pageSize]
		} else {
			start := cursor - cursor%pageSize
			elems = elems[start:(start + pageSize)]
		}

	}

	new_elems := apply_vertical_limit(elems, lights_vertical_limit)

	items := lipgloss.JoinVertical(lipgloss.Left, new_elems...)

	if selected {
		return selectedStyle.Render(items)
	}
	return defaultStyle.Render(items)

}
