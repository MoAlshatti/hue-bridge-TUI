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
	Dimming   Dimming
	Preset    string
	MirekMax  int
	MirekMin  int
	On        bool
	ColorTemp ColorTemperature
}

// Light is a panel type which incluedes an array of lights and a panel cursor...
type Lights struct {
	Items  []Light
	Cursor int
}

func (l *Lights) Increment() {
	if l.Cursor < len(l.Items)-1 {
		l.Cursor++
	}
}
func (l *Lights) Decrement() {
	if l.Cursor > 0 {
		l.Cursor--
	}
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
	Items  []Group
	LookUp map[LightID]Group
	//Selected bool
	Cursor int
}

func (g *Groups) Increment() {
	if g.Cursor < len(g.Items)-1 {
		g.Cursor++
	}
}
func (g *Groups) Decrement() {
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
	Items  [2]string //its gonna be 2, quit and buttonPressed
	Cursor int
}

func (u *UserPage) Increment() {
	if u.Cursor < len(u.Items)-1 {
		u.Cursor++
	}
}
func (u *UserPage) Decrement() {
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
	Items  []Scene
	Cursor int
}

func (s *Scenes) Increment() {
	if s.Cursor < len(s.Items)-1 {
		s.Cursor++
	}
}

func (s *Scenes) Decrement() {
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
