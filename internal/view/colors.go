package view

import (
	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
	"github.com/lucasb-eyer/go-colorful"
)

var (
	White        bridge.Color = bridge.Color{Name: "cool white", Val: colorful.Xyy(0.31, 0.33, 1)}
	NeutralWhite bridge.Color = bridge.Color{Name: "neutral white", Val: colorful.Xyy(0.3227, 0.329, 1)}
	WarmWhite    bridge.Color = bridge.Color{Name: "warm white", Val: colorful.Xyy(0.457, 0.41, 1)}
	Yellow       bridge.Color = bridge.Color{Name: "yellow", Val: colorful.Xyy(0.4432, 0.5154, 1)}
	Amber        bridge.Color = bridge.Color{Name: "amber (golden yellow)", Val: colorful.Xyy(0.5, 0.45, 1)}
	Orange       bridge.Color = bridge.Color{Name: "orange", Val: colorful.Xyy(0.5562, 0.4084, 1)}
	Red          bridge.Color = bridge.Color{Name: "red", Val: colorful.Xyy(0.675, 0.322, 1)}
	DeepRed      bridge.Color = bridge.Color{Name: "deep red", Val: colorful.Xyy(0.7, 0.3, 1)}
	RosePink     bridge.Color = bridge.Color{Name: "rose pink", Val: colorful.Xyy(0.3824, 0.1601, 1)}
	Magenta      bridge.Color = bridge.Color{Name: "magenta", Val: colorful.Xyy(0.382, 0.16, 1)}
	Green        bridge.Color = bridge.Color{Name: "green", Val: colorful.Xyy(0.1724, 0.7468, 1)}
	LimeGreen    bridge.Color = bridge.Color{Name: "lime green", Val: colorful.Xyy(0.408, 0.517, 1)}
	DarkGreen    bridge.Color = bridge.Color{Name: "dark green", Val: colorful.Xyy(0.214, 0.709, 1)}
	Cyan         bridge.Color = bridge.Color{Name: "cyan (light blue)", Val: colorful.Xyy(0.168, 0.336, 1)}
	SkyBlue      bridge.Color = bridge.Color{Name: "sky blue", Val: colorful.Xyy(0.15, 0.2, 1)}
	RoyalBlue    bridge.Color = bridge.Color{Name: "royal blue", Val: colorful.Xyy(0.1355, 0.0399, 1)}
	Navy         bridge.Color = bridge.Color{Name: "navy (charcoal blue)", Val: colorful.Xyy(0.16, 0.09, 1)}
	Aqua         bridge.Color = bridge.Color{Name: "aqua (turquoise)", Val: colorful.Xyy(0.17, 0.35, 1)}
	Violet       bridge.Color = bridge.Color{Name: "violet (purple)", Val: colorful.Xyy(0.2725, 0.1096, 1)}
	Indigo       bridge.Color = bridge.Color{Name: "indigo (dark purple)", Val: colorful.Xyy(0.245, 0.12, 1)}
	Lavender     bridge.Color = bridge.Color{Name: "lavender (soft purple)", Val: colorful.Xyy(0.31, 0.2, 1)}
	Lilac        bridge.Color = bridge.Color{Name: "lilac (light purple)", Val: colorful.Xyy(0.33, 0.25, 1)}
)
