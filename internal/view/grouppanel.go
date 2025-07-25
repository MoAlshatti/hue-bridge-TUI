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
func Render_group_panel(elems []string, selected bool, cursor int) string {

	border := lipgloss.RoundedBorder()
	border.TopLeft = "2" // gotta find a better way to title borders

	defaultStyle := lipgloss.NewStyle().Border(border).Margin(0, 1).PaddingLeft(1)
	selectedStyle := defaultStyle.BorderForeground(cyan)

	if len(elems) > max_groups_page_size {
		pagesize := min(max_groups_page_size, len(elems))
		if cursor%pagesize == 0 {
			if cursor+pagesize > len(elems) {
				elems = elems[cursor:]
			} else {
				elems = elems[cursor : cursor+pagesize]
			}
		} else {
			start := cursor - cursor%pagesize
			if start+pagesize > len(elems) {
				elems = elems[start:]
			} else {
				elems = elems[start : start+pagesize]
			}
		}
	}

	elems = apply_vertical_limit(elems, groups_vertical_limit)

	items := lipgloss.JoinVertical(lipgloss.Left, elems...)

	if selected {
		return selectedStyle.Render(items)
	}
	return defaultStyle.Render(items)
}
