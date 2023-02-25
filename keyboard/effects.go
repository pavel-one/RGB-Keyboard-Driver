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

	blackColor := colorful.Color{
		R: 0,
		G: 0,
		B: 0,
	}

	c2 := c1.BlendRgb(blackColor, 0.50)
	c3 := c1.BlendRgb(blackColor, 0.90)

	for index, key := range k.Keymap {
		if index != 0 {
			old := k.Keymap[index-1]
			old.Red = int(c2.R)
			old.Green = int(c2.G)
			old.Blue = int(c2.B)
		}

		if index >= 2 {
			old := k.Keymap[index-2]
			old.Red = int(c3.R)
			old.Green = int(c3.G)
			old.Blue = int(c3.B)
		}

		if index >= 3 {
			old := k.Keymap[index-3]
			old.Red = 0
			old.Green = 0
			old.Blue = 0
		}

		key.Red = int(c1.R)
		key.Green = int(c1.G)
		key.Blue = int(c1.B)

		time.Sleep(time.Millisecond * 100)
	}

	k.ResetState()

	k.PrintText("zov")
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

func (k *Keyboard) Fill(r int, g int, b int) error {
	for _, key := range k.Keymap {
		key.Red = r
		key.Green = g
		key.Blue = b
	}

	return nil
}
