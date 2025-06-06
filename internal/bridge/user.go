package bridge

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type UserFoundMsg string
type NoUserFoundMsg ErrMsg

const fileName = "userinfo.txt"

func Find_User(b Bridge) tea.Cmd {
	return func() tea.Msg {
		data, err := os.ReadFile(fileName)
		username := strings.TrimSpace(string(data))
		if err != nil {
			return NoUserFoundMsg(ErrMsg{err})
		}
		url := fmt.Sprintf("https://%s/clip/v2/resource/bridge", b.Ip_addr)

		req, cancel, err := create_finduser_req(url, string(username))
		if err != nil {
			return NoUserFoundMsg(ErrMsg{err})
		}
		defer cancel()
		resp, err := client.Do(req)
		if err != nil {
			return NoUserFoundMsg(ErrMsg{err})
		}
		if resp.StatusCode != http.StatusOK {
			return NoUserFoundMsg(ErrMsg{errors.New("invalid user!")})
		}
		return UserFoundMsg("Yayy user found")
	}
}
func create_finduser_req(url, username string) (*http.Request, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		cancel()
		return nil, nil, err
	}
	req.Header.Add("hue-application-key", username)
	req.Header.Add("Accept", "application/json")
	return req, cancel, nil
}

type UserCreatedMsg string
type UserCreationFailedMsg ErrMsg

func Create_User(b Bridge) tea.Cmd {
	return func() tea.Msg {
		return ""
	}
}
