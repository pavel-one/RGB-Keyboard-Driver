package keyboard

import (
	"github.com/lucasb-eyer/go-colorful"
	"strings"
	"time"
)

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

		time.Sleep(time.Millisecond * 80)
	}

	if err := k.ResetState(); err != nil {
		k.ErrorCh <- err
	}

	err := k.Fill(&colorful.Color{
		R: 0,
		G: 255,
		B: 255,
	})
	if err != nil {
		k.ErrorCh <- err
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

func (k *Keyboard) Fill(color *colorful.Color) error {
	for _, key := range k.Keymap {
		key.Fill(color)
	}

	return nil
}
