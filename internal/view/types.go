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
	default_horizontal_limit = 50

	groups_vertical_limit = 6
	lights_vertical_limit = 8
	scenes_vertical_limit = 6

	max_groups_page_size = 6
	max_lights_page_size = 8
	max_scenes_page_size = 6
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
