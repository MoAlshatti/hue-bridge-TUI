package bridge

// Filters lights based on the selected group
func Filter_lights(l *Lights, g Groups) {

	if g.Cursor == 0 {
		l.Items = l.AllItems
		return
	}
	var lights []Light
	for _, light := range l.AllItems {
		for _, child := range g.Items[g.Cursor].Children {
			if light.ID == child.Rid || light.owner.Rid == child.Rid {
				lights = append(lights, light)
			}
		}
	}
	l.Items = lights
	l.Cursor = 0
}

func Filter_scenes(s *Scenes, g Groups) {
	if g.Cursor == 0 {
		s.Items = s.AllItems
		return
	}
	var scenes []Scene
	for _, scene := range s.AllItems {
		if scene.Group_Rid == g.Items[g.Cursor].ID {
			scenes = append(scenes, scene)
		}
	}
	s.Items = scenes
	s.Cursor = 0
}
