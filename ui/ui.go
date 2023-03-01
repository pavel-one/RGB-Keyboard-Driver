package ui

import (
	"KeyboardDriver/keyboard"
	"github.com/gotk3/gotk3/gtk"
)

type UI struct {
	Keyboard *keyboard.Keyboard
	Window   *gtk.Window
}

func CreateUi(k *keyboard.Keyboard, name string) (*UI, error) {
	gtk.Init(nil)

	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}

	window.SetTitle(name)
	window.SetDefaultSize(600, 600)

	return &UI{
		Keyboard: k,
		Window:   window,
	}, nil
}

func (u *UI) Run() error {
	if err := constructUi(u.Window); err != nil {
		return err
	}

	u.Window.ShowAll()

	gtk.Main()

	return nil
}
