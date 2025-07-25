package view

import "github.com/charmbracelet/lipgloss"

func Render_details_panel(elems []string) string {
	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Margin(0, 3).PaddingLeft(1).MarginTop(1).MarginRight(3)

	elems = apply_vertical_limit(elems, details_vertical_limit)

	items := lipgloss.JoinVertical(lipgloss.Left, elems...)

	return style.Render(items)
}
