package bridge

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func set_header(req *http.Request, appkey string) {
	req.Header.Add("hue-application-key", appkey)
}

type LightsMsg []Light
type FailedFetchingLightsMsg ErrMsg

func Fetch_lights(b Bridge, appkey string, logger *LogFile) tea.Cmd {
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
			logger.Log_Print(v.Error.Description)
		}

		lights := make([]Light, 0, 15)
		for _, v := range apiLights.Data {
			var light Light
			light.ID = v.ID
			light.Type = v.Type
			light.Metadata = v.Metadata
			light.Dimming = v.Dimming
			light.Color = v.Color.Xy
			light.Preset = v.Powerup.Preset
			light.MirekMax = v.ColorTemperature.MirekSchema.MirekMaximum
			light.MirekMin = v.ColorTemperature.MirekSchema.MirekMinimum
			light.owner.Rid = v.Owner.Rid
			light.owner.Rtype = v.Owner.Rtype
			light.ColorTemp = v.ColorTemperature
			light.On = v.On.On
			light.Connected = true
			lights = append(lights, light)
		}

		return LightsMsg(lights)
	}
}

type GroupsMsg []Group
type FailedToFetchGroupsMsg ErrMsg

func Fetch_groups(b Bridge, appkey string, logger *LogFile) tea.Cmd {
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
				logger.Log_Print(err.Error.Description)
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

				for _, service := range group.Services {
					if service.Rtype == "grouped_light" {
						newGroup.GroupID = service.Rid
					}
				}
				groups = append(groups, newGroup)
			}

			// query each one for group_lights
			baseurl := fmt.Sprintf("https://%s/clip/v2/resource", b.Ip_addr)
			lightgroups, err := fetch_lightgroups(baseurl, appkey, groups)
			if err != nil {
				return FailedToFetchGroupsMsg(ErrMsg{err})
			}
			groups = lightgroups
		}
		return GroupsMsg(groups)
	}
}

func fetch_lightgroups(baseurl, appkey string, groups []Group) ([]Group, error) {
	for i := range groups {

		var serviceID string
		for _, service := range groups[i].Services {
			if service.Rtype == "grouped_light" {
				serviceID = service.Rid
			}
		}
		url := fmt.Sprintf("%s/grouped_light/%s", baseurl, serviceID)

		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
		set_header(req, appkey)
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("%s: failed to fetch grouped_light", resp.Status)
		}

		var apiLightGroup ApiGroupedLights
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&apiLightGroup)
		if err != nil {
			return nil, err
		}

		groups[i].On = apiLightGroup.Data[0].On.On
		groups[i].Brightness = apiLightGroup.Data[0].Dimming.Brightness
		groups[i].MinDim = apiLightGroup.Data[0].Dimming.MinDimLevel
	}
	return groups, nil
}

type ScenesMsg []Scene
type FailedToFetchScenesMsg ErrMsg

func Fetch_Scenes(b Bridge, appkey string, logger *LogFile) tea.Cmd {
	return func() tea.Msg {

		url := fmt.Sprintf("https://%s/clip/v2/resource/scene", b.Ip_addr)

		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return FailedToFetchScenesMsg(ErrMsg{err})
		}
		set_header(req, appkey)

		resp, err := client.Do(req)
		if err != nil {
			return FailedToFetchScenesMsg(ErrMsg{err})
		}

		if resp.StatusCode != http.StatusOK {
			return FailedToFetchScenesMsg(ErrMsg{fmt.Errorf("Failed To fetch scenes: %v", resp.Status)})
		}

		var ApiScenes ApiScene

		decoder := json.NewDecoder(resp.Body)
		defer resp.Body.Close()

		err = decoder.Decode(&ApiScenes)
		if err != nil {
			return FailedToFetchScenesMsg(ErrMsg{err})
		}

		for _, v := range ApiScenes.Errors {
			logger.Log_Print(v.Error.Description)

		}

		scenes := make([]Scene, 0, 10)

		for _, scene := range ApiScenes.Data {
			var newScene Scene

			newScene.ID = scene.ID
			switch scene.Status.Active {
			case "inactive":
				newScene.Active = false
			default:
				newScene.Active = true
			}
			newScene.LastRecall = scene.Status.LastRecall
			newScene.Name = scene.Metadata.Name
			newScene.Speed = scene.Speed
			newScene.Group_Rid = scene.Group.Rid
			newScene.Group_Rtype = scene.Group.Rtype

			scenes = append(scenes, newScene)
		}
		return ScenesMsg(scenes)
	}
}

type ConnectivityMsg []Connectivity

func Fetch_connectivity(b Bridge, appkey string) tea.Cmd {
	return func() tea.Msg {
		//
		url := fmt.Sprintf("https://%s/clip/v2/resource/zigbee_connectivity", b.Ip_addr)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return ResourceErrMsg(ErrMsg{err})
		}
		set_header(req, appkey)

		resp, err := client.Do(req)
		if err != nil {
			return ResourceErrMsg(ErrMsg{err})
		}

		if resp.StatusCode != http.StatusOK {
			return ResourceErrMsg(ErrMsg{fmt.Errorf("http error: %v", resp.Status)})
		}

		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		var apiConn ApiConnectivity
		err = decoder.Decode(&apiConn)
		if err != nil {
			return ResourceErrMsg(ErrMsg{err})
		}

		for _, v := range apiConn.Errors {
			log.Println(v.Error.Description)
		}
		var connectedDevices []Connectivity
		for _, v := range apiConn.Data {
			connectedDevices = append(connectedDevices, Connectivity{v.Owner.Rid, v.Status})
		}
		return ConnectivityMsg(connectedDevices)
	}
}

type ResourceErrMsg ErrMsg
type ResourceSuccessMsg string

func Change_light_state(b Bridge, light Light, on bool, appkey string) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf("https://%s/clip/v2/resource/light/%s", b.Ip_addr, light.ID)

		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()

		body := struct {
			On On `json:"on"`
		}{On: On{on}}

		var buff bytes.Buffer
		json.NewEncoder(&buff).Encode(&body)

		req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, &buff)
		if err != nil {
			return ResourceErrMsg(ErrMsg{err})
		}
		set_header(req, appkey)
		resp, err := client.Do(req)
		if err != nil {
			return ResourceErrMsg(ErrMsg{err})
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			//light.On = !light.On
			return ResourceSuccessMsg(fmt.Sprint(light.Metadata.Name, " state changed!"))
		}
		return ResourceErrMsg(ErrMsg{errors.New(resp.Status + ": Failed to change " + light.Metadata.Name + " state!")})
	}
}

func Change_light_brightness(b Bridge, light Light, bri float64, appkey string) tea.Cmd {
	return func() tea.Msg {

		//this function should not be called if brightness is invalid, but this is a double check
		if bri > 100.0 || bri < 0.0 {
			return ResourceErrMsg(ErrMsg{errors.New("Invalid brightness!")})
		}

		url := fmt.Sprintf("https://%s/clip/v2/resource/light/%s", b.Ip_addr, light.ID)

		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()

		body := struct {
			Dimming struct {
				Brightness float64 `json:"brightness"`
			} `json:"dimming"`
		}{
			Dimming: struct {
				Brightness float64 `json:"brightness"`
			}{bri},
		}

		var buff bytes.Buffer
		json.NewEncoder(&buff).Encode(&body)

		req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, &buff)
		if err != nil {
			return ResourceErrMsg(ErrMsg{err})
		}
		set_header(req, appkey)
		resp, err := client.Do(req)
		if err != nil {
			return ResourceErrMsg(ErrMsg{err})
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			return ResourceSuccessMsg(fmt.Sprint(light.Metadata.Name, " brightness changed!"))
		}
		return ResourceErrMsg(ErrMsg{errors.New(resp.Status + ": " + light.Metadata.Name + " Failed to change Brightness!")})
	}
}

func Change_light_color(b Bridge, light Light, color Color, appkey string) tea.Cmd {
	return func() tea.Msg {

		url := fmt.Sprintf("https://%s/clip/v2/resource/light/%s", b.Ip_addr, light.ID)
		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()
		x, y, _ := color.Val.Xyy()

		type Col struct {
			Xy XyColor `json:"xy"`
		}
		body := struct {
			Color Col `json:"color"`
		}{Color: Col{Xy: XyColor{X: x, Y: y}}}
		var buff bytes.Buffer
		json.NewEncoder(&buff).Encode(&body)

		req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, &buff)
		if err != nil {
			return ResourceErrMsg(ErrMsg{err})
		}
		set_header(req, appkey)
		resp, err := client.Do(req)
		if err != nil {
			return ResourceErrMsg(ErrMsg{err})
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			return ResourceSuccessMsg(fmt.Sprint(light.Metadata.Name, " color changed!"))
		}
		return ResourceErrMsg(ErrMsg{errors.New(resp.Status + ": " + light.Metadata.Name + " Failed to change Color!")})
	}

}

func Pick_scene(b Bridge, scene Scene, appkey string) tea.Cmd {
	return func() tea.Msg {

		url := fmt.Sprintf("https://%s/clip/v2/resource/scene/%s", b.Ip_addr, scene.ID)

		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()

		body := struct {
			Recall struct {
				Action string `json:"action"`
			} `json:"recall"`
		}{Recall: struct {
			Action string `json:"action"`
		}{"active"}}

		var buff bytes.Buffer
		json.NewEncoder(&buff).Encode(&body)

		req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, &buff)
		if err != nil {
			return ResourceErrMsg(ErrMsg{err})
		}
		set_header(req, appkey)

		resp, err := client.Do(req)
		if err != nil {
			return ResourceErrMsg(ErrMsg{err})
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			return ResourceSuccessMsg(scene.Name + " has been picked sucessfully!")
		}
		return ResourceErrMsg(ErrMsg{errors.New(resp.Status + ": Failed to pick " + scene.Name)})
	}
}

func Change_group_state(b Bridge, group Group, on bool, appkey string) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf("https://%s/clip/v2/resource/grouped_light/%s", b.Ip_addr, group.GroupID)

		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()

		body := struct {
			On On `json:"on"`
		}{On: On{on}}

		var buff bytes.Buffer
		json.NewEncoder(&buff).Encode(&body)

		req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, &buff)
		if err != nil {
			return ResourceErrMsg(ErrMsg{err})
		}
		set_header(req, appkey)

		resp, err := client.Do(req)
		if err != nil {
			return ResourceErrMsg(ErrMsg{err})
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			group.On = !group.On
			return ResourceSuccessMsg(group.Metadata.Name + " has been changed sucessfully!")
		}
		return ResourceErrMsg(ErrMsg{errors.New(resp.Status + ": Failed to change " + group.Metadata.Name)})
	}
}

func Change_group_brightness(b Bridge, group Group, bri float64, appkey string) tea.Cmd {
	return func() tea.Msg {

		url := fmt.Sprintf("https://%s/clip/v2/resource/grouped_light/%s", b.Ip_addr, group.GroupID)

		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()

		body := struct {
			Dimming struct {
				Brightness float64 `json:"brightness"`
			} `json:"dimming"`
		}{
			Dimming: struct {
				Brightness float64 `json:"brightness"`
			}{bri},
		}

		var buff bytes.Buffer
		json.NewEncoder(&buff).Encode(&body)

		req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, &buff)
		if err != nil {
			return ResourceErrMsg(ErrMsg{err})
		}
		set_header(req, appkey)
		resp, err := client.Do(req)
		if err != nil {
			return ResourceErrMsg(ErrMsg{err})
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			group.Brightness = bri
			return ResourceSuccessMsg(fmt.Sprint(group.Metadata.Name, " brightness changed!"))
		}
		return ResourceErrMsg(ErrMsg{errors.New(resp.Status + " Failed to change Brightness!")})
	}
}
