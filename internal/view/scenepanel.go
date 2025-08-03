package view

import (
	"fmt"
	"strings"
	"time"

	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
	"github.com/charmbracelet/lipgloss"
)

func Render_scene_title(title string, on, selected bool, width, height int) string {

	status := ""

	if on {
		status = "Active "
	}

	style := lipgloss.NewStyle().Width(get_scenepanel_width(width) - len(status))
	selectedStyle := style.Background(white).Foreground(navy)

	statusStyle := lipgloss.NewStyle().Align(lipgloss.Right).Width(len(status))
	selectedStatusStyle := statusStyle.Background(white).Foreground(navy)

	if selected {
		return lipgloss.JoinHorizontal(lipgloss.Right, selectedStyle.Render(title), selectedStatusStyle.Render(status))
	}
	return lipgloss.JoinHorizontal(lipgloss.Right, style.Render(title), statusStyle.Render(status))
}

func Render_scene_panel(elems []string, selected bool, cursor, width, height int) string {
	border := lipgloss.RoundedBorder()
	border.TopLeft = "4"

	defaultStyle := lipgloss.NewStyle().
		Border(border).
		Margin(0, 1).
		PaddingLeft(1).
		Height(get_scenepanel_height(height))
	selectedStyle := defaultStyle.BorderForeground(cyan)

	if len(elems) > get_scenepanel_height(height) {
		pageSize := get_scenepanel_height(height)
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

	items := lipgloss.JoinVertical(lipgloss.Left, elems...)

	if selected {
		return selectedStyle.Render(items)
	}
	return defaultStyle.Render(items)

}

func Render_scenes(s bridge.Scenes, p bridge.Panel, width, height int) string {
	var scenes []string
	for i, v := range s.Items {
		scenes = append(scenes, Render_scene_title(v.Name,
			v.Active,
			i == s.Cursor && p == bridge.ScenePanel, width, height))
	}
	return Render_scene_panel(scenes, p == bridge.ScenePanel, s.Cursor, width, height)
}

func Render_scene_details(s bridge.Scene, width, height int) string {
	style := lipgloss.NewStyle().
		Italic(true).
		Bold(true).
		Width(get_detailspanel_width(width))

	name := style.Render(fmt.Sprintln("Name: ", s.Name))
	id := style.Render(fmt.Sprintln("ID: ", s.ID))
	active := style.Render(fmt.Sprintln("Active: ", s.Active))
	speed := style.Render(fmt.Sprintln("Speed: ", s.Speed))
	lastrecall := style.Render(fmt.Sprintln("Last Recall: ", s.LastRecall.Format(time.RFC850)))
	var group []string
	group = append(group, style.Render(fmt.Sprint("Group Information: ")))
	group = append(group, style.Render(fmt.Sprint("Rid: ", s.Group_Rid)))
	group = append(group, style.Render(fmt.Sprintln("Rtype: ", s.Group_Rtype)))

	output := lipgloss.JoinVertical(lipgloss.Left,
		name,
		id,
		active,
		lipgloss.JoinVertical(lipgloss.Left, group...),
		speed,
		lastrecall)

	if lipgloss.Height(output) >= get_detailspanel_height(height) {
		output_array := strings.Split(output, "\n")
		return lipgloss.JoinVertical(lipgloss.Left, output_array[:get_detailspanel_height(height)]...)
	}
	return output
}
