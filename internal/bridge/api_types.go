package bridge

type AuthSuccess struct {
	Success struct {
		UserName  string `json:"username"`
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
	}
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

type ApiLights struct {
	Errors []ApiError `json:"errors"`
	Data   []struct {
		ID    string `json:"id"`
		IDV1  string `json:"id_v1"`
		Owner struct {
			Rid   string `json:"rid"`
			Rtype string `json:"rtype"`
		} `json:"owner"`
		Metadata    Metadata `json:"metadata"`
		ProductData struct {
			Function string `json:"function"`
		} `json:"product_data"`
		Identify struct {
		} `json:"identify"`
		ServiceID int `json:"service_id"`
		On        struct {
			On bool `json:"on"`
		} `json:"on"`
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
				On   struct {
					On bool `json:"on"`
				} `json:"on"`
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
