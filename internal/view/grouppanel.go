package view

import (
	"github.com/charmbracelet/lipgloss"
)

func Render_group_title(title string, selected bool, width, height int) string {
	defaultStyle := lipgloss.NewStyle().Width(get_grouppanel_width(width))
	selectedStyle := defaultStyle.Background(white).Foreground(navy)

	if selected {
		return selectedStyle.Render(title)
	}
	return defaultStyle.Render(title)
}
func Render_group_panel(elems []string, selected bool, cursor, width, height int) string {

	border := lipgloss.RoundedBorder()
	border.TopLeft = "2" // gotta find a better way to title borders

	defaultStyle := lipgloss.NewStyle().
		Border(border).
		Margin(0, 1).
		PaddingLeft(1).
		Height(get_grouppanel_height(height))
	selectedStyle := defaultStyle.BorderForeground(cyan)

	//consider making a function that does this shit
	if len(elems) > get_grouppanel_height(height) {
		pagesize := get_grouppanel_height(height)
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

	items := lipgloss.JoinVertical(lipgloss.Left, elems...)

	if selected {
		return selectedStyle.Render(items)
	}
	return defaultStyle.Render(items)
}
