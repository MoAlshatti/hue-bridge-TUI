package bridge

import (
	"encoding/json"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tmaxmax/go-sse"
)

// Server sent events

type sseResponse struct {
	Creationtime time.Time        `json:"creationtime"`
	Data         []map[string]any `json:"data"`
	ID           string           `json:"id"`
	Type         string           `json:"type"`
}

type SseUpdate struct {
	Id   string
	Type string
}

type LightStateUpdate struct {
	SseUpdate
	On bool
}

type BriUpdate struct {
	SseUpdate
	Brightness float64
}
type ColorUpdate struct {
	SseUpdate
	Color XyColor
}
type SceneStateUpdate struct {
	SseUpdate
	Active string
}

type SseFailedMsg ErrMsg

// Maybe i should run this outside the update function, as a concurrent function in main
// and inject msgs
func Initiate_sse(b Bridge, appkey string, p *tea.Program) tea.Cmd {
	return func() tea.Msg {

		sseClient := sse.DefaultClient
		sseClient.HTTPClient = client

		url := "https://" + b.Ip_addr + "/eventstream/clip/v2"

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			// consider using ResourceErr
			return SseFailedMsg(ErrMsg{err})
		}
		set_header(req, appkey)
		conn := sseClient.NewConnection(req)

		_ = conn.SubscribeToAll(func(e sse.Event) {
			//
			var resps []sseResponse
			json.Unmarshal([]byte(e.Data), &resps)

			// go through each response, and deal with it
			for _, resp := range resps {
				//
				for _, v := range resp.Data {
					// 1- find the id and type
					sseUpdate := find_sse_update(v)
					// 2- check the update type
					update := fetch_sse_update(v, sseUpdate)
					switch update := update.(type) {
					case LightStateUpdate:
						p.Send(update)
					case BriUpdate:
						p.Send(update)
					case ColorUpdate:
						p.Send(update)
					default:
						// do nothing
					}
				}
			}

		})

		err = conn.Connect()
		if err != nil {
			return SseFailedMsg(ErrMsg{err})
		}
		return ""
	}
}

func find_sse_update(obj map[string]any) (s SseUpdate) {
	switch id := obj["id"].(type) {
	case string:
		s.Id = id
	}
	switch objType := obj["type"].(type) {
	case string:
		s.Type = objType
	}
	return s
}

func fetch_sse_update(obj map[string]any, sseupdate SseUpdate) any {
	if v, ok := obj["on"]; ok {
		on := v.(map[string]any)["on"]
		return LightStateUpdate{sseupdate, on.(bool)}
	} else if v, ok := obj["dimming"]; ok {
		bri := v.(map[string]any)["brightness"]
		return BriUpdate{sseupdate, bri.(float64)}

	} else if v, ok := obj["color"]; ok {
		xy := v.(map[string]any)["xy"]
		x := xy.(map[string]any)["x"]
		y := xy.(map[string]any)["y"]
		return ColorUpdate{sseupdate, XyColor{x.(float64), y.(float64)}}

	} else if _, ok := obj["status"]; ok {

	}
	return nil
}

func Update_light_status(lights *[]Light, status LightStateUpdate) {
	for i := range *lights {
		if (*lights)[i].ID == status.Id {
			(*lights)[i].On = status.On
			break
		}
	}
}
func Update_light_brightness(lights *[]Light, status BriUpdate) {
	for i := range *lights {
		if (*lights)[i].ID == status.Id {
			(*lights)[i].Dimming.Brightness = status.Brightness
			break
		}
	}
}
func Update_group_brightness(groups *[]Group, status BriUpdate) {
	for i := range *groups {
		if (*groups)[i].GroupID == status.Id {
			(*groups)[i].Brightness = status.Brightness
			break
		}
	}
}
