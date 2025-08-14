package bridge

import (
	"log"
	"slices"
)

// Filters lights based on the selected group
func Filter_lights(l *Lights, g Groups) {

	var lights []*Light

	if g.Cursor == 0 {

		for i := range l.AllItems {
			lights = append(lights, &l.AllItems[i])
		}
		l.Items = lights
		return
	}

	for i := range l.AllItems {
		for _, child := range g.Items[g.Cursor].Children {
			if l.AllItems[i].ID == child.Rid || l.AllItems[i].owner.Rid == child.Rid {
				lights = append(lights, &l.AllItems[i])
			}
		}
	}
	l.Items = lights
	l.Cursor = 0
}

func Filter_scenes(s *Scenes, g Groups) {
	var scenes []*Scene

	if g.Cursor == 0 {
		for i := range s.AllItems {
			scenes = append(scenes, &s.AllItems[i])
		}
		s.Items = scenes
		return
	}

	for i := range s.AllItems {
		if s.AllItems[i].Group_Rid == g.Items[g.Cursor].ID {
			scenes = append(scenes, &s.AllItems[i])
		}
	}
	s.Items = scenes
	s.Cursor = 0
}

func Sort_Connectivity(l *Lights, connDevices []Connectivity) {
	for _, v := range connDevices {
		for i := range l.AllItems {
			if v.ID == l.AllItems[i].owner.Rid {
				log.Println("ID MATCH!")
				switch v.Status {
				case "connectivity_issue":
					l.AllItems[i].Connected = false
				case "connected":
					l.AllItems[i].Connected = true
				}
			}
		}
	}
	slices.SortStableFunc(l.AllItems, func(l1, l2 Light) int {
		if l1.Connected && !l2.Connected {
			return -1
		} else if !l1.Connected && l2.Connected {
			return 1
		}
		return 0
	})
}
