package bridge

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func set_header(req *http.Request, appkey string) {
	req.Header.Add("hue-application-key", appkey)
}

type LightsMsg []Light
type FailedFetchingLightsMsg ErrMsg

func Fetch_lights(b Bridge, appkey string) tea.Cmd {
	return func() tea.Msg {

		url := fmt.Sprintf("https://%s/clip/v2/resource/light", b.Ip_addr)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return FailedFetchingLightsMsg(ErrMsg{err})
		}
		set_header(req, appkey)

		resp, err := client.Do(req)
		if err != nil {
			return FailedFetchingLightsMsg(ErrMsg{err})
		}

		if resp.StatusCode != http.StatusOK {
			return FailedFetchingLightsMsg(ErrMsg{fmt.Errorf("http error: %v", resp.Status)})
		}

		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		var apiLights ApiLights

		err = decoder.Decode(&apiLights)
		if err != nil {
			return FailedFetchingLightsMsg(ErrMsg{err})
		}

		for _, v := range apiLights.Errors {
			log.Println(v.Error.Description)
		}

		lights := make([]Light, 0, 15)
		for _, v := range apiLights.Data {
			var light Light
			light.ID = v.ID
			light.Type = v.Type
			light.Metadata = v.Metadata
			light.Dimming = v.Dimming
			light.Color = v.Color.Xy
			light.ColorTemp = v.ColorTemperature
			lights = append(lights, light)
		}

		return LightsMsg(lights)
	}
}

type GroupsMsg []Group
type FailedToFetchGroupsMsg ErrMsg

func Fetch_groups(b Bridge, appkey string) tea.Cmd {
	return func() tea.Msg {
		var urls []string
		url := fmt.Sprintf("https://%s/clip/v2/resource/zone", b.Ip_addr)
		url2 := fmt.Sprintf("https://%s/clip/v2/resource/room", b.Ip_addr)
		urls = append(urls, url, url2)

		groups := make([]Group, 0, 10)

		for _, v := range urls {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, v, nil)
			if err != nil {
				return FailedToFetchGroupsMsg(ErrMsg{err})
			}
			set_header(req, appkey)

			resp, err := client.Do(req)

			if resp.StatusCode != http.StatusOK {
				return FailedToFetchGroupsMsg(ErrMsg{fmt.Errorf("failed to fetch groups: %v", resp.Status)})
			}

			defer resp.Body.Close()

			decoder := json.NewDecoder(resp.Body)
			var apiGroups ApiGroup

			err = decoder.Decode(&apiGroups)
			if err != nil {
				return FailedToFetchGroupsMsg(ErrMsg{err})
			}

			for _, err := range apiGroups.Errors {
				log.Println(err.Error.Description)
			}

			for _, group := range apiGroups.Data {
				//
				var newGroup Group
				newGroup.ID = group.ID
				newGroup.Type = group.Type
				newGroup.Children = group.Children
				newGroup.Services = group.Services
				newGroup.Metadata.Archetype = group.Metadata.Archetype
				newGroup.Metadata.Name = group.Metadata.Name
				groups = append(groups, newGroup)
			}
		}
		return GroupsMsg(groups)
	}
}
