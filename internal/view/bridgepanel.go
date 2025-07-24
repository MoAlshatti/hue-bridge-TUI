package view

import (
	"github.com/charmbracelet/lipgloss"
)

func Render_bridge_title(title string) string {
	style := lipgloss.NewStyle()

	return style.Render(apply_horizontal_limit(title, default_horizontal_limit))
}
func Render_bridge_panel(title string, selected bool) string {
	border := lipgloss.RoundedBorder()
	border.TopLeft = "1"
	defaultStyle := lipgloss.NewStyle().MarginLeft(1).MarginTop(1).Border(border).PaddingLeft(1)

	selectedStyle := defaultStyle.BorderForeground(cyan)

	if selected {
		return selectedStyle.Render(title)
	}
	return defaultStyle.Render(title)
}
