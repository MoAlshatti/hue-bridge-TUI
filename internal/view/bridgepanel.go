package view

import (
	"github.com/charmbracelet/lipgloss"
)

func Render_bridge_title(title string, width, height int) string {
	style := lipgloss.NewStyle().Width(get_bridgepanel_width(width))

	return style.Render(title)

}
func Render_bridge_panel(title string, selected bool, width, height int) string {
	border := lipgloss.RoundedBorder()
	border.TopLeft = "1"
	defaultStyle := lipgloss.NewStyle().
		MarginLeft(1).
		MarginTop(1).
		Border(border).
		PaddingLeft(1).
		Height(get_bridgepanel_height(height)). //dont do shit for now, maybe useful later
		MaxHeight(5)

	selectedStyle := defaultStyle.BorderForeground(cyan)

	if selected {
		return selectedStyle.Render(title)
	}
	return defaultStyle.Render(title)
}

func Render_bridge_details(prefix, title string) string {
	style := lipgloss.NewStyle().Italic(true)

	return style.Render(apply_horizontal_limit(prefix+title, details_horizontal_limit))
}
