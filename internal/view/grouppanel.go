package view

import (
	"github.com/charmbracelet/lipgloss"
)

func Render_group_title(title string, selected bool) string {
	defaultStyle := lipgloss.NewStyle()
	selectedStyle := defaultStyle.Background(white).Foreground(navy)

	if selected {
		return selectedStyle.Render(apply_horizontal_limit(title, default_horizontal_limit))
	}
	return defaultStyle.Render(apply_horizontal_limit(title, default_horizontal_limit))
}
func Render_group_panel(elems []string, selected bool) string {

	border := lipgloss.RoundedBorder()
	border.TopLeft = "2" // gotta find a better way to title borders

	defaultStyle := lipgloss.NewStyle().Border(border).Margin(0, 1).PaddingLeft(1)
	selectedStyle := defaultStyle.BorderForeground(cyan)

	elems = apply_vertical_limit(elems, groups_vertical_limit)

	items := lipgloss.JoinVertical(lipgloss.Left, elems...)

	if selected {
		return selectedStyle.Render(items)
	}
	return defaultStyle.Render(items)
}
