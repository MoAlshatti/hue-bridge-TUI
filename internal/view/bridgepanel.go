package view

import (
	"fmt"

	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
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
		Height(get_bridgepanel_height(height)).
		MaxHeight(5)

	selectedStyle := defaultStyle.BorderForeground(cyan)

	if selected {
		return selectedStyle.Render(title)
	}
	return defaultStyle.Render(title)
}

func Render_bridge(b bridge.Bridge, p bridge.Panel, width, height int) string {

	title := Render_bridge_title("Hue Bridge", width, height)
	return Render_bridge_panel(title, p == bridge.BridgePanel, width, height)
}

func Render_bridge_details(b bridge.Bridge, width, height int) string {
	style := lipgloss.NewStyle().
		Italic(true).
		Bold(true).
		Width(get_detailspanel_width(width))

	id := style.Render(fmt.Sprintln("Bridge ID: ", b.ID))
	ip := style.Render(fmt.Sprintln("Bridge IP: ", b.Ip_addr))
	port := style.Render(fmt.Sprintln("Port: ", b.Port))

	return lipgloss.JoinVertical(lipgloss.Left, id, ip, port)

}
