package bridge

import "time"

type AuthSuccess struct {
	Success struct {
		UserName string `json:"username"`
		//[Ignore] dont use this guy
		ClientKey string `json:"clientkey"`
	} `json:"success"`
}

type ApiError struct {
	Error struct {
		Type        int    `json:"type"`
		Address     string `json:"address"`
		Description string `json:"description"`
	} `json:"error"`
}

// These structs will be used elsewhere hence they have been extracted out of ApiLights
type Dimming struct {
	Brightness  float64 `json:"brightness"`
	MinDimLevel float64 `json:"min_dim_level"`
}
type ColorTemperature struct {
	Mirek       any  `json:"mirek"`
	MirekValid  bool `json:"mirek_valid"`
	MirekSchema struct {
		MirekMinimum int `json:"mirek_minimum"`
		MirekMaximum int `json:"mirek_maximum"`
	} `json:"mirek_schema"`
}
type XyColor struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
type Metadata struct {
	Name      string `json:"name"`
	Archetype string `json:"archetype"`
	Function  string `json:"function"`
}

type Owner struct {
	Rid   string `json:"rid"`
	Rtype string `json:"rtype"`
}

type On struct {
	On bool `json:"on"`
}

type ApiLights struct {
	Errors []ApiError `json:"errors"`
	Data   []struct {
		ID          string   `json:"id"`
		IDV1        string   `json:"id_v1"`
		Owner       Owner    `json:"owner"`
		Metadata    Metadata `json:"metadata"`
		ProductData struct {
			Function string `json:"function"`
		} `json:"product_data"`
		Identify struct {
		} `json:"identify"`
		ServiceID    int     `json:"service_id"`
		On           On      `json:"on"`
		Dimming      Dimming `json:"dimming"`
		DimmingDelta struct {
		} `json:"dimming_delta"`
		ColorTemperature      ColorTemperature `json:"color_temperature"`
		ColorTemperatureDelta struct {
		} `json:"color_temperature_delta,omitempty"`
		Color struct {
			Xy    XyColor `json:"xy"`
			Gamut struct {
				Red struct {
					X float64 `json:"x"`
					Y float64 `json:"y"`
				} `json:"red"`
				Green struct {
					X float64 `json:"x"`
					Y float64 `json:"y"`
				} `json:"green"`
				Blue struct {
					X float64 `json:"x"`
					Y float64 `json:"y"`
				} `json:"blue"`
			} `json:"gamut"`
			GamutType string `json:"gamut_type"`
		} `json:"color,omitempty"`
		Dynamics struct {
			Status       string   `json:"status"`
			StatusValues []string `json:"status_values"`
			Speed        float64  `json:"speed"`
			SpeedValid   bool     `json:"speed_valid"`
		} `json:"dynamics"`
		Alert struct {
			ActionValues []string `json:"action_values"`
		} `json:"alert"`
		Signaling struct {
			SignalValues []string `json:"signal_values"`
		} `json:"signaling"`
		Mode    string `json:"mode"`
		Effects struct {
			StatusValues []string `json:"status_values"`
			Status       string   `json:"status"`
			EffectValues []string `json:"effect_values"`
		} `json:"effects"`
		EffectsV2 struct {
			Action struct {
				EffectValues []string `json:"effect_values"`
			} `json:"action"`
			Status struct {
				Effect       string   `json:"effect"`
				EffectValues []string `json:"effect_values"`
			} `json:"status"`
		} `json:"effects_v2"`
		TimedEffects struct {
			StatusValues []string `json:"status_values"`
			Status       string   `json:"status"`
			EffectValues []string `json:"effect_values"`
		} `json:"timed_effects,omitempty"`
		Powerup struct {
			Preset     string `json:"preset"`
			Configured bool   `json:"configured"`
			On         struct {
				Mode string `json:"mode"`
				On   On     `json:"on"`
			} `json:"on"`
			Dimming struct {
				Mode string `json:"mode"`
			} `json:"dimming"`
			Color struct {
				Mode string `json:"mode"`
			} `json:"color"`
		} `json:"powerup"`
		Type string `json:"type"`
	} `json:"data"`
}
type sceneStatus struct {
	Active     string    `json:"active"`
	LastRecall time.Time `json:"last_recall"`
}

type ApiScene struct {
	Errors []ApiError `json:"errors"`
	Data   []struct {
		ID      string `json:"id"`
		IDV1    string `json:"id_v1"`
		Actions []struct {
			Target struct {
				Rid   string `json:"rid"`
				Rtype string `json:"rtype"`
			} `json:"target"`
			Action struct {
				On      On `json:"on"`
				Dimming struct {
					Brightness float64 `json:"brightness"`
				} `json:"dimming"`
				ColorTemperature struct {
					Mirek int `json:"mirek"`
				} `json:"color_temperature"`
			} `json:"action"`
		} `json:"actions"`
		Palette struct {
			Color   []any `json:"color"`
			Dimming []struct {
				Brightness float64 `json:"brightness"`
			} `json:"dimming"`
			ColorTemperature []struct {
				ColorTemperature struct {
					Mirek int `json:"mirek"`
				} `json:"color_temperature"`
				Dimming struct {
					Brightness float64 `json:"brightness"`
				} `json:"dimming"`
			} `json:"color_temperature"`
			Effects   []any `json:"effects"`
			EffectsV2 []any `json:"effects_v2"`
		} `json:"palette"`
		Recall struct {
		} `json:"recall"`
		Metadata struct {
			Name  string `json:"name"`
			Image struct {
				Rid   string `json:"rid"`
				Rtype string `json:"rtype"`
			} `json:"image"`
		} `json:"metadata"`
		Group struct {
			Rid   string `json:"rid"`
			Rtype string `json:"rtype"`
		} `json:"group"`
		Speed       float64     `json:"speed"`
		AutoDynamic bool        `json:"auto_dynamic"`
		Status      sceneStatus `json:"status"`
		Type        string      `json:"type"`
	} `json:"data"`
}

type Child struct {
	Rid   string `json:"rid"`
	Rtype string `json:"rtype"`
}
type Service struct {
	Rid   string `json:"rid"`
	Rtype string `json:"rtype"`
}
type ApiGroup struct {
	Errors []ApiError `json:"errors"`
	Data   []struct {
		ID       string    `json:"id"`
		IDV1     string    `json:"id_v1"`
		Children []Child   `json:"children"`
		Services []Service `json:"services"`
		Metadata struct {
			Name      string `json:"name"`
			Archetype string `json:"archetype"`
		} `json:"metadata"`
		Type string `json:"type"`
	} `json:"data"`
}

type ApiGroupedLights struct {
	Errors []ApiError `json:"errors"`
	Data   []struct {
		ID           string  `json:"id"`
		IDV1         string  `json:"id_v1"`
		Owner        Owner   `json:"owner"`
		On           On      `json:"on"`
		Dimming      Dimming `json:"dimming"`
		DimmingDelta struct {
		} `json:"dimming_delta"`
		ColorTemperature struct {
		} `json:"color_temperature"`
		ColorTemperatureDelta struct {
		} `json:"color_temperature_delta"`
		Color struct {
		} `json:"color"`
		Alert struct {
			ActionValues []string `json:"action_values"`
		} `json:"alert"`
		Signaling struct {
			SignalValues []string `json:"signal_values"`
		} `json:"signaling"`
		Dynamics struct {
		} `json:"dynamics"`
		Type string `json:"type"`
	} `json:"data"`
}
