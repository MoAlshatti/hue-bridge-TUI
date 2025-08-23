package bridge

import (
	"encoding/json"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea/v2"
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

type StateUpdate struct {
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
	Status sceneStatus
}
type SceneActiveUpdate struct {
	SseUpdate
	Active string `json:"active"`
}
type SceneRecallUpdate struct {
	SseUpdate
	LastRecall time.Time `json:"last_recall"`
}
type ZigbeeUpdate struct {
	SseUpdate
	DeviceID string
	Status   string
}

//color temp update in the future mayhaps (ez but mehh)

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
					case StateUpdate:
						p.Send(update)
					case BriUpdate:
						p.Send(update)
					case ColorUpdate:
						p.Send(update)
					case SceneStateUpdate:
						p.Send(update)
					case SceneActiveUpdate:
						p.Send(update)
					case SceneRecallUpdate:
						p.Send(update)
					case ZigbeeUpdate:
						p.Send(update)
					default:
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
	id, ok := obj["id"].(string)
	if ok {
		s.Id = id
	}
	objType, ok := obj["type"].(string)
	if ok {
		s.Type = objType
	}
	return s
}

func fetch_sse_update(obj map[string]any, sseupdate SseUpdate) any {
	if sseupdate.Type == "zigbee_connectivity" {
		status, ok := obj["status"]
		if ok {
			if ok {
				owner := obj["owner"]
				rid := owner.(map[string]any)["rid"]
				return ZigbeeUpdate{sseupdate, rid.(string), status.(string)}
			}
		}
		return "nothing"
	}
	if v, ok := obj["on"]; ok {
		on := v.(map[string]any)["on"]
		return StateUpdate{sseupdate, on.(bool)}
	} else if v, ok := obj["dimming"]; ok {
		bri := v.(map[string]any)["brightness"]
		return BriUpdate{sseupdate, bri.(float64)}

	} else if v, ok := obj["color"]; ok {
		xy := v.(map[string]any)["xy"]
		x := xy.(map[string]any)["x"]
		y := xy.(map[string]any)["y"]
		return ColorUpdate{sseupdate, XyColor{x.(float64), y.(float64)}}

	} else if v, ok := obj["status"]; ok {
		activeObj := v.(map[string]any)["active"]
		active, okActive := activeObj.(string)
		lastrecallObj := v.(map[string]any)["last_recall"]
		lastrecall, okRecall := lastrecallObj.(string)
		t, err := time.Parse(time.RFC3339, lastrecall)
		if okRecall && okActive && err == nil {
			return SceneStateUpdate{sseupdate, sceneStatus{active, t}}
		} else if okActive && !okRecall {
			return SceneActiveUpdate{sseupdate, active}
		} else if okRecall && !okActive && err == nil {
			return SceneRecallUpdate{sseupdate, t}
		}
	}
	return "None"
}

func find_light(lights []Light, lightID string) *Light {
	for i := range lights {
		if lights[i].ID == lightID {
			return &lights[i]
		}
	}
	return nil
}
func Update_light_status(lights []Light, status StateUpdate) {
	l := find_light(lights, status.Id)
	if l.ID == status.Id {
		l.On = status.On
		l.Connected = true
	}
}
func Update_group_status(groups []Group, status StateUpdate) {
	for i := range groups {
		if groups[i].GroupID == status.Id {
			groups[i].On = status.On
			break
		}
	}
}
func Update_light_brightness(lights []Light, status BriUpdate) {
	l := find_light(lights, status.Id)
	if l.ID == status.Id {
		l.Dimming.Brightness = status.Brightness
		l.Connected = true
	}
}
func Update_group_brightness(groups []Group, status BriUpdate) {
	for i := range groups {
		if groups[i].GroupID == status.Id {
			groups[i].Brightness = status.Brightness
			break
		}
	}
}
func Update_light_color(lights []Light, status ColorUpdate) {
	l := find_light(lights, status.Id)
	if l.ID == status.Id {
		l.Color = status.Color
	}
}
func find_scene(scenes []Scene, sceneID string) *Scene {
	for i := range scenes {
		if scenes[i].ID == sceneID {
			return &scenes[i]
		}
	}
	return nil
}
func Update_scene_status(scenes []Scene, status SceneStateUpdate) {
	s := find_scene(scenes, status.Id)
	if status.Status.Active == "inactive" {
		s.Active = false
	} else {
		s.Active = true
	}
	s.LastRecall = status.Status.LastRecall
}
func Update_Scene_active(scenes []Scene, activeUpdate SceneActiveUpdate) {
	s := find_scene(scenes, activeUpdate.Id)
	if activeUpdate.Active == "inactive" {
		s.Active = false
	} else {
		s.Active = true
	}

}
func Update_Scene_recall(scenes []Scene, recallUpdate SceneRecallUpdate) {
	s := find_scene(scenes, recallUpdate.Id)
	s.LastRecall = recallUpdate.LastRecall
}

func Update_light_connection(lights []Light, zig ZigbeeUpdate) {
	for i := range lights {
		if lights[i].owner.Rid == zig.DeviceID {
			switch zig.Status {
			case "connectivity_issue", "disconnected":
				lights[i].Connected = false
			case "connected":
				lights[i].Connected = true
			}
			break
		}
	}
}
