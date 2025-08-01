package view

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Render_log_title(elem string, width, height int) string {
	style := lipgloss.NewStyle().
		Width(get_logpanel_width(width))

	output := style.Render(elem)

	if lipgloss.Height(output) >= get_logpanel_height(height) {
		output_array := strings.Split(output, "\n")
		last_elem := output_array[len(output_array)-1]
		output_array[len(output_array)-1] = strings.TrimSuffix(last_elem, "\r")
		output_array[len(output_array)-1] = strings.TrimSuffix(last_elem, "\n")
		return lipgloss.JoinVertical(lipgloss.Left, output_array[lipgloss.Height(output)-get_logpanel_height(height):]...)
	}
	return output
}

func Render_log_panel(elem string, width, height int) string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		PaddingLeft(1).
		MarginRight(3).
		Height(get_logpanel_height(height))

	return style.Render(elem)
}
