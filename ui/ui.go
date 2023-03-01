package ui

import (
	"KeyboardDriver/keyboard"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

const appID = "org.gtk.kd87a"
const flags = glib.APPLICATION_FLAGS_NONE

type UI struct {
	Keyboard    *keyboard.Keyboard
	Application *gtk.Application
}

func NewUI(k *keyboard.Keyboard, name string) (*UI, error) {
	application, err := gtk.ApplicationNew(appID, flags)
	if err != nil {
		return nil, err
	}

	ui := &UI{
		Keyboard:    k,
		Application: application,
	}

	application.Connect("activate", func() {
		if err := ui.constructUi(application, name); err != nil {
			log.Fatalf("Error construct UI: %s", err)
		}
	})

	return ui, err
}

func (u *UI) constructUi(app *gtk.Application, name string) error {
	window, err := gtk.ApplicationWindowNew(app)
	if err != nil {
		return err
	}
	window.SetTitle(name)
	window.SetDefaultSize(800, 600)

	grid, err := createGrid()
	if err != nil {
		return err
	}

	sidebar, err := createSidebar()
	if err != nil {
		return err
	}

	right, err := createRightBox()
	if err != nil {
		return err
	}

	grid.Attach(sidebar, 1, 1, 1, 1)
	grid.Attach(right, 2, 1, 1, 1)

	window.Add(grid)

	window.ShowAll()

	return nil
}
