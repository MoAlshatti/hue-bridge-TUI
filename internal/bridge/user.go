package bridge

import tea "github.com/charmbracelet/bubbletea"

type UserFoundMsg string
type NoUserFoundMsg ErrMsg

func Find_User(b Bridge) tea.Cmd {
	return func() tea.Msg {

	}
}

type UserCreatedMsg string
type UserCreationFailedMsg ErrMsg

func Create_User(b Bridge) tea.Cmd {
	return func() tea.Msg {

	}
}
