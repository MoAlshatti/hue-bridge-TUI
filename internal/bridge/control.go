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
			light.metadata = v.Metadata
			light.Color = v.Color.Xy
			light.ColorTemp = v.ColorTemperature
			lights = append(lights, light)
		}

		return LightsMsg(lights)
	}
}
