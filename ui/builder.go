package ui

import (
	"github.com/gotk3/gotk3/gtk"
)

func (u *UI) createGrid() (*gtk.Grid, error) {
	g, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}

	g.SetHExpand(true)

	return g, nil
}

func (u *UI) createSidebar() (*gtk.Box, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		return nil, err
	}

	box.SetVAlign(gtk.ALIGN_CENTER)
	box.SetVExpand(true)
	box.SetMarginEnd(10)
	box.SetMarginStart(10)

	b, err := gtk.ButtonNewFromIconName("go-home", gtk.ICON_SIZE_LARGE_TOOLBAR)
	if err != nil {
		return nil, err
	}
	b.SetLabel("Test 1")
	b.SetAlwaysShowImage(true)
	box.Add(b)

	b1, err := gtk.ButtonNewFromIconName("starred", gtk.ICON_SIZE_LARGE_TOOLBAR)
	if err != nil {
		return nil, err
	}
	b1.SetLabel("Test 2")
	b1.SetAlwaysShowImage(true)
	box.Add(b1)

	l, err := gtk.LabelNew("TEST!")
	if err != nil {
		return nil, err
	}
	box.Add(l)

	return box, nil
}

func (u *UI) createRightBox() (*gtk.Box, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	if err != nil {
		return nil, err
	}

	box.SetHExpand(true)
	box.SetVAlign(gtk.ALIGN_CENTER)
	box.SetHAlign(gtk.ALIGN_CENTER)

	keymap := u.Keyboard.KeymapMatrix
	for _, row := range keymap {
		lb, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
		if err != nil {
			return nil, err
		}

		for _, key := range row {
			btn, err := gtk.ButtonNew()
			if err != nil {
				return nil, err
			}
			btn.SetSizeRequest(100, 30)
			if key != nil {
				btn.SetLabel(key.Name)
				btn.Connect("clicked", func() {
					l, _ := btn.GetLabel()
					u.SetColor(l)
				})
			} else {
				btn.SetOpacity(0.0)
			}

			lb.Add(btn)
		}

		box.Add(lb)
	}

	return box, nil
}
