package view

import (
	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
	"github.com/charmbracelet/lipgloss"
)

func Render_userpage_title(title string) string {
	style := lipgloss.NewStyle().
		Bold(true).
		Border(lipgloss.DoubleBorder()).
		Align(lipgloss.Center, lipgloss.Center).
		Margin(1, 3).
		Padding(0, 2)

	return style.Render(title)
}

func Render_userpage_options(option string, selected bool) string {
	unselectedStyle := lipgloss.NewStyle().
		Margin(0, 5).
		Padding(1).
		MarginTop(2)

	selectedStyle := unselectedStyle.Foreground(lipgloss.Color("#52eb34"))

	if selected {
		return selectedStyle.Render(option)
	} else {
		return unselectedStyle.Render(option)
	}
}

func Render_userpage_panel(title, optQuit, optPress string) string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Align(lipgloss.Center, lipgloss.Center).
		Margin(3).
		Padding(2, 5)

	opts := lipgloss.JoinHorizontal(lipgloss.Center, optPress, optQuit)
	userpage := lipgloss.JoinVertical(lipgloss.Center, title, opts)

	return style.Render(userpage)
}

func Render_userpage(u bridge.UserPage) string {
	leftmargin := lipgloss.NewStyle().MarginLeft(2)
	rightmargin := lipgloss.NewStyle().MarginRight(2)

	title := Render_userpage_title("Press the hue bridge button!")
	quitOpt := Render_userpage_options(u.Items[bridge.Quit], u.Cursor == bridge.Quit)
	pressOpt := Render_userpage_options(u.Items[bridge.PressTheButton], u.Cursor != bridge.Quit)

	quitOpt = leftmargin.Render(quitOpt)
	pressOpt = rightmargin.Render(pressOpt)

	return Render_userpage_panel(title, quitOpt, pressOpt)

}
