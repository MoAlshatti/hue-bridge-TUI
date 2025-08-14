package view

import (
	"fmt"
	"strings"

	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

func Render_light_title(title string, bri float64, on bool, selected bool, width, height int) string {

	status := ""
	if !on {
		status = "OFF "
	} else {
		status = fmt.Sprint(int(bri), "% ")
	}

	style := lipgloss.NewStyle().Width((get_lightpanel_width(width)) - len(status))
	selectedStyle := style.Background(white).Foreground(navy)

	statusStyle := lipgloss.NewStyle().Align(lipgloss.Right).Width(len(status))
	selectedStatusStyle := statusStyle.Background(white).Foreground(navy)

	if selected {
		return lipgloss.JoinHorizontal(lipgloss.Right, selectedStyle.Render(title), selectedStatusStyle.Render(status))
	}
	return lipgloss.JoinHorizontal(lipgloss.Right, style.Render(title), statusStyle.Render(status))
}

func Render_light_panel(elems []string, selected bool, cursor, width, height int) string {
	border := lipgloss.RoundedBorder()
	border.TopLeft = "3"

	defaultStyle := lipgloss.NewStyle().
		Border(border).
		Margin(0, 1).
		PaddingLeft(1).
		Height(get_lightpanel_height(height))

	selectedStyle := defaultStyle.BorderForeground(cyan)

	if len(elems) > get_lightpanel_height(height) {
		pageSize := get_lightpanel_height(height)
		if cursor%pageSize == 0 {
			if cursor+pageSize > len(elems) {
				elems = elems[cursor:]
			} else {
				elems = elems[cursor : cursor+pageSize]
			}
		} else {
			start := cursor - cursor%pageSize
			if start+pageSize > len(elems) {

				elems = elems[start:]
			} else {
				elems = elems[start : start+pageSize]
			}
		}
	}
	items := lipgloss.JoinVertical(lipgloss.Left, elems...)

	if selected {
		return selectedStyle.Render(items)
	}
	return defaultStyle.Render(items)
}

func Render_lights(l bridge.Lights, g bridge.Groups, p bridge.Panel, width, height int) string {
	var lights []string

	if len(l.Items) == 0 {
		lights = append(lights, " ")
	}

	for i, v := range l.Items {
		lights = append(lights, Render_light_title(v.Metadata.Name,
			v.Dimming.Brightness,
			v.On, i == l.Cursor && p == bridge.LightPanel, width, height))
	}
	return Render_light_panel(lights, p == bridge.LightPanel, l.Cursor, width, height)
}

func Render_light_details(l bridge.Light, width, height int) string {
	style := lipgloss.NewStyle().
		Italic(true).
		Bold(true).
		Width(get_detailspanel_width(width))

	name := style.Render(fmt.Sprintln("Name: ", l.Metadata.Name))
	archtype := style.Render(fmt.Sprintln("Archtype: ", l.Metadata.Archetype))
	function := style.Render(fmt.Sprintln("Function: ", l.Metadata.Function))
	id := style.Render(fmt.Sprintln("ID: ", l.ID))
	var color []string

	col := colorful.Xyy(l.Color.X, l.Color.Y, 1)
	hexCol := col.Clamped().Hex()
	renderedCol := lipgloss.NewStyle().Background(lipgloss.Color(hexCol))
	color = append(color, style.Render("Color: "))
	color = append(color, style.Render(fmt.Sprint("  X: ", l.Color.X)))
	if l.Color.X == 0 && l.Color.Y == 0 {
		color = append(color, style.Render(fmt.Sprintln("  Y:", l.Color.Y)))
	} else {
		color = append(color, style.Render(fmt.Sprint("  Y: ", l.Color.Y)))
		color = append(color, "  "+renderedCol.Render("        ")+"\n")
	}

	var colortemp []string
	colortemp = append(colortemp, style.Render("Color Temperature: "))
	colortemp = append(colortemp, style.Render(fmt.Sprint("  Mirek : ", l.ColorTemp.Mirek)))
	colortemp = append(colortemp, style.Render(fmt.Sprintln("  Mirek Valid: ", l.ColorTemp.MirekValid)))
	colortemp = append(colortemp, style.Render("Mirek Schema: "))
	colortemp = append(colortemp, style.Render(fmt.Sprint("  Mirek Maximum: ", l.MirekMax)))
	colortemp = append(colortemp, style.Render(fmt.Sprintln("  Mirek Minimum:", l.MirekMin)))

	brightness := style.Render(fmt.Sprintln("Brightness: ", l.Dimming.Brightness))
	preset := style.Render(fmt.Sprintln("Preset: ", l.Preset))
	minDimLevel := style.Render(fmt.Sprintln("Minimum Dimming Level: ", l.Dimming.MinDimLevel))

	output := lipgloss.JoinVertical(lipgloss.Left, name,
		archtype,
		function,
		id,
		lipgloss.JoinVertical(lipgloss.Left, color...),
		lipgloss.JoinVertical(lipgloss.Left, colortemp...),
		brightness,
		preset,
		minDimLevel)

	if lipgloss.Height(output) >= get_detailspanel_height(height) {
		output_array := strings.Split(output, "\n")
		return lipgloss.JoinVertical(lipgloss.Left, output_array[:get_detailspanel_height(height)]...)
	}
	return output
}
