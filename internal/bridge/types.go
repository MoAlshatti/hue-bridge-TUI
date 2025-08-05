package bridge

import (
	"fmt"
	"log"
	"time"
)

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

type Panel int

const (
	_ Panel = iota
	BridgePanel
	GroupPanel
	LightPanel
	ScenePanel
)

// Light represents ..... TODO
type Light struct {
	ID        string
	Type      string
	Metadata  Metadata
	Color     XyColor
	owner     Owner
	Dimming   Dimming
	Preset    string
	MirekMax  int
	MirekMin  int
	On        bool
	ColorTemp ColorTemperature
}

// Light is a panel type which includes two array of lights(filtered/not filtered) and a panel cursor...
type Lights struct {
	//filtered lights showed based on the selected group
	Items []*Light
	// all lights no filters
	AllItems []Light
	Cursor   int
}

func (l *Lights) increment() {
	if l.Cursor < len(l.Items)-1 {
		l.Cursor++
	}
}
func (l *Lights) decrement() {
	if l.Cursor > 0 {
		l.Cursor--
	}
}

// Group is either a zone or a room.... TODO
type Group struct { // consider making it lowercase if you aint finna use it in main
	ID       string
	Children []Child
	Services []Service
	Metadata struct {
		Name      string
		Archetype string
	}
	GroupID    string
	On         bool
	Brightness float64
	MinDim     float64
	Type       string
}

// Groups is a panel type which includes an array of groups, a panel cursor, and a map from the lights to the groups
type Groups struct {
	Items  []Group
	Cursor int
}

func (g *Groups) increment() {
	if g.Cursor < len(g.Items)-1 {
		g.Cursor++
	}
}
func (g *Groups) decrement() {
	if g.Cursor > 0 {
		g.Cursor--
	}
}

const (
	//shitty naming conventions,well..
	Quit           = 0
	PressTheButton = 1
)

type UserPage struct {
	Items  [2]string //quit and buttonPressed
	Cursor int
}

func (u *UserPage) increment() {
	if u.Cursor < len(u.Items)-1 {
		u.Cursor++
	}
}
func (u *UserPage) decrement() {
	if u.Cursor > 0 {
		u.Cursor--
	}
}

// Bridge has all the needed bridge info ....
type Bridge struct {
	ID      string `json:"id"`
	Ip_addr string `json:"internalipaddress"`
	Port    int    `json:"port"`
	Info    []string
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
	AllItems []Scene
	Items    []*Scene
	Cursor   int
}

func (s *Scenes) increment() {
	if s.Cursor < len(s.Items)-1 {
		s.Cursor++
	}
}

func (s *Scenes) decrement() {
	if s.Cursor > 0 {
		s.Cursor--
	}
}

// represents the user uesd for the API
type User struct {
	Username string `json:"Key"`
}

type ErrMsg struct {
	Err error
}

func (e ErrMsg) Error() string { return e.Err.Error() }

type LogFile struct {
	Content string
}

func (l *LogFile) Log_Print(v ...any) {
	log.Println(v...)
	l.Content = l.Content + fmt.Sprint(time.Now().Format(time.DateTime)) + " " + fmt.Sprintln(v...)
}
