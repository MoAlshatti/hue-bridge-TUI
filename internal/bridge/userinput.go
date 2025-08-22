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

//The List component

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
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

func Initialize_list() list.Model {
	newlist := list.New([]list.Item{
		NeutralWhite, White, WarmWhite,
		Yellow, Amber, Orange,
		Red, DeepRed, RosePink, Magenta,
		Green, LimeGreen, DarkGreen,
		Cyan, SkyBlue, RoyalBlue, Navy, Aqua,
		Violet, Indigo, Lavender, Lilac,
	}, itemDelegate{}, 20, 20)
	return newlist
}
