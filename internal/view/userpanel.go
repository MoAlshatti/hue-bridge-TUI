package view

import "github.com/charmbracelet/lipgloss"

func Render_userpage_title(title string) string {
	style := lipgloss.NewStyle().Bold(true).Border(lipgloss.DoubleBorder()).Align(lipgloss.Center, lipgloss.Center).Margin(1, 3).Padding(0, 2)

	return style.Render(title)
}

func Render_userpage_options(option string, selected bool) string {
	unselectedStyle := lipgloss.NewStyle().Margin(1, 5).Padding(1)

	selectedStyle := unselectedStyle.Foreground(lipgloss.Color("#52eb34"))

	if selected {
		return selectedStyle.Render(option)
	} else {
		return unselectedStyle.Render(option)
	}
}

func Render_userpage(title, optQuit, optPress string) string {
	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Align(lipgloss.Center, lipgloss.Center).Margin(3).Padding(3, 10)

	opts := lipgloss.JoinHorizontal(lipgloss.Center, optPress, optQuit)
	userpage := lipgloss.JoinVertical(lipgloss.Center, title, opts)

	return style.Render(userpage)
}
