package bridge

import "time"

type Event int

const (
	_ Event = iota
	//listt all events here..
	FindingBridge
	FindingUser
	RequestPressButton
	CreateUser //do we need it?
	FetchingLights
	DisplayingLights
)

// Light represents ..... TODO
type Light struct {
	ID        string
	Type      string
	Metadata  Metadata
	Color     XyColor
	Dimming   Dimming
	Preset    string
	MirekMax  int
	MirekMin  int
	On        bool
	ColorTemp ColorTemperature
}

// Light is a panel type which incluedes an array of lights and a panel cursor...
type Lights struct {
	Items    []Light
	Selected bool
	Cursor   int
}

// LightID .... TODO
type LightID string

// Group is either a zone or a room.... TODO
type Group struct { // consider making it lowercase if you aint finna use it in main
	ID       string
	Children []Child
	Services []Service
	Metadata struct {
		Name      string
		Archetype string
	}
	Type string
}

// Groups is a panel type which includes an array of groups, a panel cursor, and a map from the lights to the groups
type Groups struct {
	Items    []Group
	LookUp   map[LightID]Group
	Selected bool
	Cursor   int
}

const (
	Quit           = 0
	PressTheButton = 1
)

type UserPage struct {
	Items  [2]string //its gonna be 2, quit and buttonPressed
	Cursor int
}

// Bridge has all the needed bridge info ....
type Bridge struct {
	ID       string `json:"id"`
	Ip_addr  string `json:"internalipaddress"`
	Port     int    `json:"port"`
	Selected bool
	Info     []string
}

type Scene struct {
	ID          string
	Active      bool
	LastRecall  time.Time
	Group_Rid   string
	Group_Rtype string
	Name        string
	Speed       float64
}

type Scenes struct {
	Items    []Scene
	Selected bool
	Cursor   int
}

// represents the user uesd for the API
type User struct {
	Username string `json:"Key"`
}

type ErrMsg struct {
	err error
}

func (e ErrMsg) Error() string { return e.err.Error() }
