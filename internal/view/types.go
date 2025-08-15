package view

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	white = lipgloss.Color("#FFFFFF")
	navy  = lipgloss.Color("#000080")
	green = lipgloss.Color("#59ff85")
	cyan  = lipgloss.Color("#00FFFF")
	aqua  = lipgloss.Color("#A7D8DE")
)

func get_lightpanel_width(width int) int {
	return (width / 3) + (width / 14)
}
func get_bridgepanel_width(width int) int {
	return get_lightpanel_width(width)
}
func get_grouppanel_width(width int) int {
	return get_lightpanel_width(width)
}
func get_scenepanel_width(width int) int {
	return get_lightpanel_width(width)
}
func get_detailspanel_width(width int) int {
	return (width / 2) + (width / 21)
}
func get_logpanel_width(width int) int {
	return get_detailspanel_width(width)
}
func get_colormodal_width(width int) int {
	return width / 3
}

func get_bridgepanel_height(height int) int {
	return height / 25
}
func get_lightpanel_height(height int) int {
	return int(0.25 * float64(height))
}
func get_grouppanel_height(height int) int {
	return int(0.25 * float64(height))
}
func get_scenepanel_height(height int) int {
	return int(0.25 * float64(height))
}
func get_detailspanel_height(height int) int {
	return get_bridgepanel_height(height) + get_grouppanel_height(height) + get_lightpanel_height(height) + (height / 10)
}
func get_logpanel_height(height int) int {
	return get_scenepanel_height(height)
}
func get_colormodal_height(height int) int {
	return int(0.25 * float64(height))
}
