package keyboard

import (
	"github.com/lucasb-eyer/go-colorful"
	"strings"
	"time"
)

// WelcomeEffect driver first run effect, check all diodes
func (k *Keyboard) WelcomeEffect() {
	c1 := colorful.Color{
		R: 255,
		G: 0,
		B: 255,
	}

	blackColor := colorful.Color{}

	c2 := c1.BlendRgb(blackColor, 0.50)
	c3 := c1.BlendRgb(blackColor, 0.90)

	for index, key := range k.Keymap {
		if index != 0 {
			old := k.Keymap[index-1]
			old.Fill(&c2)
		}

		if index >= 2 {
			old := k.Keymap[index-2]
			old.Fill(&c3)
		}

		if index >= 3 {
			old := k.Keymap[index-3]
			old.Reset()
		}

		key.Fill(&c1)

		time.Sleep(time.Millisecond * 25)
	}

	if err := k.ResetState(); err != nil {
		k.ErrorCh <- err
	}

	err := k.Fill(RandomHappyColor())
	if err != nil {
		k.ErrorCh <- err
		return
	}
}

func (k *Keyboard) PrintText(word string) {
	word = strings.ToUpper(word)
	wordArr := strings.Split(word, "")

	for _, char := range wordArr {
		for _, key := range k.Keymap {
			if key.Name == char {
				key.Red = 255
			}
		}
	}

}

// Fill set color full keyboard
func (k *Keyboard) Fill(color *colorful.Color) error {
	keymap := k.KeymapMatrix

	for _, row := range keymap {
		row := row
		// fill rows parallels
		go func() {
			for _, key := range row {
				if key == nil {
					FakeButton()
					continue
				}
				key.FillSmooth(color)
			}
		}()
	}

	return nil
}
