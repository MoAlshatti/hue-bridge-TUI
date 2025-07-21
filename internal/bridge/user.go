package bridge

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
		return UserFoundMsg(username)
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
type ButtonNotPressed string
type UserCreationFailedMsg ErrMsg

// This function should be called after Find_User
func Create_User(b Bridge) tea.Cmd {
	return func() tea.Msg {

		url := fmt.Sprintf("https://%s/api", b.Ip_addr)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		body := struct {
			Devicetype        string `json:"devicetype"`
			Generateclientkey bool   `json:"generateclientkey"`
		}{
			"hueApp#root",
			true,
		}
		var buff bytes.Buffer
		json.NewEncoder(&buff).Encode(&body)
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buff)
		if err != nil {
			return UserCreationFailedMsg(ErrMsg{err})
		}

		resp, err := client.Do(req)
		if err != nil {
			return UserCreationFailedMsg(ErrMsg{err})
		}
		defer resp.Body.Close()
		var apiErr []ApiError
		decoder := json.NewDecoder(resp.Body)
		decoder.DisallowUnknownFields()

		err = decoder.Decode(&apiErr)

		log.Println(apiErr[0].Error)
		if err == nil {
			if apiErr[0].Error.Type == 101 {
				return ButtonNotPressed("Error 101, button not pressed")
			}
			return UserCreationFailedMsg(ErrMsg{fmt.Errorf("error %v\n", apiErr[0].Error)})
		}
		var auth []AuthSuccess
		decoder = json.NewDecoder(resp.Body)
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&auth)
		for i, v := range auth {
			log.Println("item", i, " ", v)
		}
		if err != nil {
			return UserCreationFailedMsg(ErrMsg{err})
		}
		log.Println("user created! ", auth[0].Success.ClientKey)
		return UserCreatedMsg(auth[0].Success.ClientKey)
	}
}

type UserNotSaved ErrMsg

func Save_Username(username string) tea.Cmd {
	return func() tea.Msg {
		err := os.WriteFile(fileName, []byte(username), 0644)
		if err != nil {
			return UserNotSaved(ErrMsg{err})
		}
		return nil
	}
}
