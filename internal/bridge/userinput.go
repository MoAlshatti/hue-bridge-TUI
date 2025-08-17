package bridge

import (
	"errors"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/v2/textinput"
)

// controlling color and brightness input
func (bm *BrightnessModal) Init() {
	ti := textinput.New()
	ti.CharLimit = 3
	ti.Prompt = "❯ "
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

func (bm *BrightnessModal) On() {
	bm.Input.Focus()
}

func (bm *BrightnessModal) Off() {
	bm.Input.Blur()
	bm.Input.Reset()
}
