package bridge

type Incrementable interface {
	increment()
	decrement()
}

// The only reason this is used is because i fancy using interfaces
func Increment_cursor(i Incrementable) {
	i.increment()
}
func Decrement_cusror(i Incrementable) {
	i.decrement()
}

type Keybind struct {
	Key         string
	Name        string
	Description string
}

const (
	Left, Right = "←/h", "→/l"
	Enter, Esc  = "<enter>", "<esc>"
	QMark       = "?"
	Sep         = " | "
)

var (
	Up           = Keybind{"↑/k", "Up", "Navigate to the previous item"}
	Down         = Keybind{"↓/j", "Down", "Navigate to the next item"}
	Help         = Keybind{QMark, "Keybinds", "Display All Available keybinds"}
	BridgePan    = Keybind{"[1]", "Bridge", "Selects the bridge panel"}
	GroupsPan    = Keybind{"[2]", "Groups", "Selects the Groups panel"}
	LightsPan    = Keybind{"[3]", "Lights", "Selects the Lights panel"}
	ScenesPan    = Keybind{"[4]", "Scenes", "Selects the Scenes panel"}
	Bri          = Keybind{"b", "Change bright", "Changes the brightness to the desired value"}
	Increase_bri = Keybind{Right, "Increase bright", "Increases brightness by 20%"}
	Decrease_bri = Keybind{Left, "Decrease bright", "Decreases brightness by 20%"}
	Apply        = Keybind{Enter, "Apply", "Apply the selected value to the item"}
	Col          = Keybind{"c", "Change color", "Choose from the list of available colors"}
	NextPage     = Keybind{Right, "PgNext", "Navigate to the next page of the list"}
	PrevPage     = Keybind{Left, "PgPrev", "Navigate to the previous page of the list"}
	Filter       = Keybind{"/", "Filter", "Filter items by search value"}
	CancelBri    = Keybind{Esc + "/b", "Cancel", "Cancel brightness mode"}
	Cancel       = Keybind{Esc, "Cancel", "Close the current modal window"}
	Off          = Keybind{Enter, "on/off,", "Toggle On/Off"}
	MoveLeft     = Keybind{Left, "Navigate Left", ""}
	MoveRight    = Keybind{Right, "Navigate Right", ""}
)
