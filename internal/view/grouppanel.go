package view

import (
	"fmt"
	"strings"

	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
	"github.com/charmbracelet/lipgloss"
)

func Render_group_title(title string, selected bool, width, height int) string {
	defaultStyle := lipgloss.NewStyle().Width(get_grouppanel_width(width))
	selectedStyle := defaultStyle.Background(aqua).Foreground(navy)

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

func Render_group(g bridge.Groups, p bridge.Panel, width, height int) string {
	var groups []string
	for i, v := range g.Items {
		groups = append(groups, Render_group_title(v.Metadata.Name, i == g.Cursor, width, height))
	}
	return Render_group_panel(groups, p == bridge.GroupPanel, g.Cursor, width, height)
}

func Render_group_details(group bridge.Group, width, height int) string {
	style := lipgloss.NewStyle().
		Italic(true).
		Bold(true).
		Width(get_detailspanel_width(width))

	name := style.Render(fmt.Sprintln("Group Name: ", group.Metadata.Name))
	id := style.Render(fmt.Sprintln("ID: ", group.ID))
	on := style.Render(fmt.Sprintln("On: ", group.On))
	bri := style.Render(fmt.Sprintf("Brightness: %.2f \n", group.Brightness))

	group_type := style.Render(fmt.Sprintln("Group Type: ", group.Type))
	archtype := style.Render(fmt.Sprintln("Archtype: ", group.Metadata.Archetype))

	var children []string
	children = append(children, style.Render("Children: "))
	for i, v := range group.Children {
		children = append(children, style.Render(fmt.Sprint("Child ", i+1)))
		children = append(children, style.Render(fmt.Sprint("    Rid: ", v.Rid)))
		children = append(children, style.Render(fmt.Sprintln("    Rtype: ", v.Rtype)))
	}
	var services []string
	services = append(services, style.Render("Group Services: "))
	for i, v := range group.Services {
		services = append(services, style.Render(fmt.Sprint("Service ", i+1)))
		services = append(services, style.Render(fmt.Sprint("    Rid: ", v.Rid)))
		services = append(services, style.Render(fmt.Sprintln("    Rtype: ", v.Rtype)))
	}
	output := lipgloss.JoinVertical(lipgloss.Left, name,
		id,
		on,
		bri,
		group_type,
		archtype,
		lipgloss.JoinVertical(lipgloss.Left, children...),
		lipgloss.JoinVertical(lipgloss.Left, services...))

	if group.ID == "None" {
		return style.Render(" No Group Chosen ! ")
	}

	if lipgloss.Height(output) >= get_detailspanel_height(height) {
		output_array := strings.Split(output, "\n")
		return lipgloss.JoinVertical(lipgloss.Left, output_array[:get_detailspanel_height(height)]...)
	}
	return output
}
