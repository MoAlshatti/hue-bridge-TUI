package view

import "github.com/charmbracelet/lipgloss"

func Render_scene_title(title string, on, selected bool) string {
	style := lipgloss.NewStyle()
	selectedStyle := style.Background(white).Foreground(navy)

	status := ""

	if on {
		status = "Active "
	}

	output := apply_horizontal_limit(title, default_horizontal_limit)

	output = output[:len(output)-len(status)]
	output = output + status

	if selected {
		return selectedStyle.Render(output)
	}
	return style.Render(output)
}

func Render_scene_panel(elems []string, selected bool, cursor int) string {
	border := lipgloss.RoundedBorder()
	border.TopLeft = "4"

	defaultStyle := lipgloss.NewStyle().Border(border).Margin(0, 1).PaddingLeft(1)
	selectedStyle := defaultStyle.BorderForeground(cyan)

	if len(elems) > max_scenes_page_size {
		pageSize := min(max_scenes_page_size, len(elems))
		if cursor%pageSize == 0 {
			if cursor+pageSize > len(elems) {
				elems = elems[cursor:]
			} else {
				elems = elems[cursor : cursor+pageSize]
			}
		} else {
			start := cursor - cursor%pageSize
			if cursor+pageSize > len(elems) {
				elems = elems[start:]
			} else {
				elems = elems[start:(start + pageSize)]
			}
		}
	}

	new_elems := apply_vertical_limit(elems, scenes_vertical_limit)

	items := lipgloss.JoinVertical(lipgloss.Left, new_elems...)

	if selected {
		return selectedStyle.Render(items)
	}
	return defaultStyle.Render(items)

}
