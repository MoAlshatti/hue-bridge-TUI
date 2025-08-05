package bridge

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
