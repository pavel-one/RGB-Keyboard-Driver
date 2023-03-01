package ui

import "KeyboardDriver/keyboard"

func (u *UI) SetColor(name string) {
	for _, key := range u.Keyboard.Keymap {
		if key.Name == name {
			key.FillSmooth(keyboard.RandomHappyColor())
			return
		}
	}
}
