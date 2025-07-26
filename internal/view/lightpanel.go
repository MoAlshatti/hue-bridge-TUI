package view

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func Render_light_title(title string, bri float64, on bool, selected bool, width, height int) string {

	status := ""
	if !on {
		status = "OFF "
	} else {
		status = fmt.Sprint(int(bri), "% ")
	}

	style := lipgloss.NewStyle().Width((get_lightpanel_width(width)) - len(status))
	selectedStyle := style.Background(white).Foreground(navy)

	statusStyle := lipgloss.NewStyle().Align(lipgloss.Right).Width(len(status))
	selectedStatusStyle := statusStyle.Background(white).Foreground(navy)

	if selected {
		return lipgloss.JoinHorizontal(lipgloss.Right, selectedStyle.Render(title), selectedStatusStyle.Render(status))
	}
	return lipgloss.JoinHorizontal(lipgloss.Right, style.Render(title), statusStyle.Render(status))
}

func Render_light_panel(elems []string, selected bool, cursor, width, height int) string {
	border := lipgloss.RoundedBorder()
	border.TopLeft = "3"

	defaultStyle := lipgloss.NewStyle().
		Border(border).
		Margin(0, 1).
		PaddingLeft(1).
		Height(get_lightpanel_height(height))

	selectedStyle := defaultStyle.BorderForeground(cyan)

	if len(elems) > get_lightpanel_height(height) {
		pageSize := get_lightpanel_height(height)
		if cursor%pageSize == 0 {
			if cursor+pageSize > len(elems) {
				elems = elems[cursor:]
			} else {
				elems = elems[cursor : cursor+pageSize]
			}
		} else {
			start := cursor - cursor%pageSize
			if start+pageSize > len(elems) {

				elems = elems[start:]
			} else {
				elems = elems[start : start+pageSize]
			}
		}
	}

	items := lipgloss.JoinVertical(lipgloss.Left, elems...)

	if selected {
		return selectedStyle.Render(items)
	}
	return defaultStyle.Render(items)

}
