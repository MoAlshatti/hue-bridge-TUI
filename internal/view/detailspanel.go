package view

import "github.com/charmbracelet/lipgloss"

func Render_details_panel(elems string, width, height int) string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		PaddingLeft(1).
		MarginTop(1).
		MarginRight(3).
		Height(get_detailspanel_height(height))

	return style.Render(elems)
}
