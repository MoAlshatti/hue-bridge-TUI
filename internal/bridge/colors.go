package bridge

import (
	"github.com/lucasb-eyer/go-colorful"
)

var (
	White        Color = Color{Name: "cool white", Val: colorful.Xyy(0.31, 0.33, 1)}
	NeutralWhite Color = Color{Name: "neutral white", Val: colorful.Xyy(0.3227, 0.329, 1)}
	WarmWhite    Color = Color{Name: "warm white", Val: colorful.Xyy(0.457, 0.41, 1)}

	Yellow Color = Color{Name: "yellow", Val: colorful.Xyy(0.4432, 0.5154, 1)}
	Amber  Color = Color{Name: "amber (golden yellow)", Val: colorful.Xyy(0.5, 0.45, 1)}
	Orange Color = Color{Name: "orange", Val: colorful.Xyy(0.5562, 0.4084, 1)}

	Red      Color = Color{Name: "red", Val: colorful.Xyy(0.675, 0.322, 1)}
	DeepRed  Color = Color{Name: "deep red", Val: colorful.Xyy(0.7, 0.3, 1)}
	RosePink Color = Color{Name: "rose pink", Val: colorful.Xyy(0.3824, 0.1601, 1)}
	Magenta  Color = Color{Name: "magenta", Val: colorful.Xyy(0.382, 0.16, 1)}

	Green     Color = Color{Name: "green", Val: colorful.Xyy(0.1724, 0.7468, 1)}
	LimeGreen Color = Color{Name: "lime green", Val: colorful.Xyy(0.408, 0.517, 1)}
	DarkGreen Color = Color{Name: "dark green", Val: colorful.Xyy(0.214, 0.709, 1)}

	Cyan      Color = Color{Name: "cyan (light blue)", Val: colorful.Xyy(0.168, 0.336, 1)}
	SkyBlue   Color = Color{Name: "sky blue", Val: colorful.Xyy(0.15, 0.2, 1)}
	RoyalBlue Color = Color{Name: "royal blue", Val: colorful.Xyy(0.1355, 0.0399, 1)}
	Navy      Color = Color{Name: "navy (charcoal blue)", Val: colorful.Xyy(0.16, 0.09, 1)}
	Aqua      Color = Color{Name: "aqua (turquoise)", Val: colorful.Xyy(0.17, 0.35, 1)}

	Violet   Color = Color{Name: "violet (purple)", Val: colorful.Xyy(0.2725, 0.1096, 1)}
	Indigo   Color = Color{Name: "indigo (dark purple)", Val: colorful.Xyy(0.245, 0.12, 1)}
	Lavender Color = Color{Name: "lavender (soft purple)", Val: colorful.Xyy(0.31, 0.2, 1)}
	Lilac    Color = Color{Name: "lilac (light purple)", Val: colorful.Xyy(0.33, 0.25, 1)}
)
