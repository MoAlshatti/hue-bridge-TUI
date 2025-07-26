package view

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	white = lipgloss.Color("#FFFFFF")
	navy  = lipgloss.Color("#000080")
	green = lipgloss.Color("#59ff85")
	cyan  = lipgloss.Color("#00FFFF")
)

var (
	details_horizontal_limit = 80
	details_vertical_limit   = 23
)

// returns the element with spaces appended to it,
// the number of spaces is determined by using the difference
// between the length and the max_horizontal_limit
// used as an alternative to lipgloss's padding
func apply_horizontal_limit(item string, lim int) string {

	limit := lim - len(item)

	spaces := strings.Repeat(" ", limit)

	return item + spaces
}

func apply_vertical_limit(list []string, lim int) []string {
	limit := lim - len(list)

	for range limit {
		list = append(list, " ")
	}
	return list
}

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
