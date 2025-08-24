package bridge

import (
	"errors"
	"fmt"
	"image/color"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/v2/list"
	"github.com/charmbracelet/bubbles/v2/textinput"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

// controlling color and brightness input
func (bm *BrightnessModal) Init() {
	ti := textinput.New()
	ti.CharLimit = 3
	ti.Prompt = "❯ "
	ti.VirtualCursor = true
	ti.Styles.Cursor.BlinkSpeed = time.Millisecond * 500
	ti.Validate = func(s string) error {
		num, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		if num > 100 || num < 0 {
			return errors.New("invalid input, brightness should be between 0 and a 100")
		}
		return nil
	}
	bm.Input = &ti
}

func (bm *BrightnessModal) Parse() (float64, error) {

	if bm.Input.Err != nil {
		return 0, bm.Input.Err
	}
	bri, err := strconv.Atoi(bm.Input.Value())
	if err != nil {
		return 0, err
	}
	return float64(bri), nil
}

func (bm *BrightnessModal) On() {
	bm.Input.Focus()
}

func (bm *BrightnessModal) Off() {
	bm.Input.Blur()
	bm.Input.Reset()
}

// =============================================================================================
//The Colors List component

type colorDelegate struct{}

func (d colorDelegate) Height() int                             { return 1 }
func (d colorDelegate) Spacing() int                            { return 0 }
func (d colorDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d colorDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Color)
	if !ok {
		return
	}

	// join the color with the name later
	col := listItem.(Color)

	str := fmt.Sprintf("%d. %s", index+1, i.Name)
	colString := "   "

	itemStyle := lipgloss.NewStyle().PaddingLeft(4).Width((m.Width() / 3))
	selectedItemStyle := itemStyle.Foreground(lipgloss.Color("#000080")).Background(color.White)

	// not used for now
	colorStyle := lipgloss.NewStyle().
		Border(lipgloss.HiddenBorder()).
		Width(len(colString)).
		Background(lipgloss.Color(col.Val.Clamped().Hex())).Align(lipgloss.Right)

	//remove later
	_ = colorStyle.String()

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render(strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func Init_color_list() list.Model {
	newlist := list.New([]list.Item{
		NeutralWhite, White, WarmWhite,
		Yellow, Amber, Orange,
		Red, DeepRed, RosePink, Magenta,
		Green, LimeGreen, DarkGreen,
		Cyan, SkyBlue, RoyalBlue, Navy, Aqua,
		Violet, Indigo, Lavender, Lilac,
	}, colorDelegate{}, 20, 20)
	return newlist
}

// =============================================================================================
// The Keybinds List component

type keybindDelegate struct{}

func (d keybindDelegate) Height() int                             { return 1 }
func (d keybindDelegate) Spacing() int                            { return 0 }
func (d keybindDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d keybindDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Keybind)
	if !ok {
		return
	}

	descStyle := lipgloss.NewStyle().Width((Get_Help_width(m.Width()))).Bold(true).PaddingLeft(1)
	selectedDescStyle := descStyle.Foreground(lipgloss.Color("#000080")).Background(color.White)
	keyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#4A90E2"))
	keySelectedStyle := keyStyle.PaddingLeft(1).Background(color.White)

	max_len := 10 - lipgloss.Width(i.Key)

	key := keyStyle.Render(i.Key + strings.Repeat(" ", max_len))

	str := key + descStyle.Render(i.Description)

	fn := descStyle.Render
	if index == m.Index() {
		key = keySelectedStyle.Render(i.Key + strings.Repeat(" ", max_len))
		str = key + selectedDescStyle.Render(i.Description)

		fn = func(s ...string) string {
			return strings.Join(s, " ")
		}
	}

	fmt.Fprint(w, fn(str))
}

func Get_Help_width(width int) int {
	return int(float64(width) / 2)

}

func Init_help_list() list.Model {
	newlist := list.New([]list.Item{
		GroupsPan, LightsPan, ScenesPan, Help,
	}, keybindDelegate{}, 20, 20)
	return newlist

}
func Update_help_list(l *list.Model, p Panel, e Event) tea.Cmd {
	switch p {
	case BridgePanel:
		return l.SetItems([]list.Item{GroupsPan, LightsPan, ScenesPan, Help})
	case GroupPanel:
		switch e {
		case DisplayingLights:
			return l.SetItems([]list.Item{Up, Down,
				BridgePan, LightsPan, ScenesPan,
				Off, Bri, Help,
				Decrease_bri, Increase_bri})
		case DisplayingBrightness:
			return l.SetItems([]list.Item{Cancel, Apply})
		}
	case LightPanel:
		switch e {
		case DisplayingLights:
			return l.SetItems([]list.Item{Up, Down,
				BridgePan, GroupsPan, ScenesPan,
				Off, Bri, Col, Help,
				Decrease_bri, Increase_bri})
		case DisplayingBrightness:
			return l.SetItems([]list.Item{Cancel, Apply})
		case DisplayingColors:
			return l.SetItems([]list.Item{Cancel,
				PrevPage, NextPage,
				Apply, Filter})
		}
	case ScenePanel:
		return l.SetItems([]list.Item{Up, Down,
			BridgePan, GroupsPan, LightsPan,
			Off, Help})
	}
	return nil
}
