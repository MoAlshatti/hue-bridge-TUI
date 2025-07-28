package view

import (
	"os"
	"strings"

	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FailedReadingLogMsg bridge.ErrMsg
type LogFileMsg string

// This function doesnt fit in the view nor bridge packages, but the view package is a better fit
func Fetch_log_file(fileName string) tea.Cmd {
	return func() tea.Msg {
		content, err := os.ReadFile(fileName)
		if err != nil {
			return FailedReadingLogMsg(bridge.ErrMsg{Err: err})
		}
		return LogFileMsg(string(content))
	}
}

func Render_log_title(elem string, width, height int) string {
	style := lipgloss.NewStyle().
		Bold(true).
		Width(get_logpanel_width(width))

	output := style.Render(elem)

	if lipgloss.Height(output) >= get_logpanel_height(height) {
		output_array := strings.Split(output, "\n")
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
