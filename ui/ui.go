package ui

import (
	"KeyboardDriver/keyboard"
	"github.com/gotk3/gotk3/gtk"
)

type UI struct {
	Keyboard *keyboard.Keyboard
	Window   *gtk.Window
}

func NewUi(k *keyboard.Keyboard, name string) (*UI, error) {
	gtk.Init(nil)

	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}

	window.SetTitle(name)

	return &UI{
		Keyboard: k,
		Window:   window,
	}, nil
}

func (u *UI) Run() error {
	//cw, _ := gtk.ColorChooserDialogNew("TEST COLOR", window)

	grid, err := u.getGrid()
	if err != nil {
		return err
	}

	sidebar, err := u.getSidebar()
	if err != nil {
		return err
	}

	right, err := u.getRightBox()
	if err != nil {
		return err
	}

	grid.Attach(sidebar, 1, 1, 1, 1)
	grid.Attach(right, 2, 1, 1, 1)

	u.Window.Add(grid)

	u.Window.ShowAll()

	gtk.Main()

	return nil
}

func (u *UI) getGrid() (*gtk.Grid, error) {
	g, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}

	g.SetColumnSpacing(20)
	g.SetRowSpacing(20)

	return g, nil
}

func (u *UI) getSidebar() (*gtk.Box, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	if err != nil {
		return nil, err
	}

	box.SetVAlign(gtk.ALIGN_CENTER)
	box.SetSpacing(10)
	b, err := gtk.ButtonNewFromIconName("go-home", gtk.ICON_SIZE_LARGE_TOOLBAR)
	if err != nil {
		return nil, err
	}
	box.Add(b)

	b1, err := gtk.ButtonNewFromIconName("starred", gtk.ICON_SIZE_LARGE_TOOLBAR)
	if err != nil {
		return nil, err
	}
	box.Add(b1)

	l, err := gtk.LabelNew("TEST!")
	if err != nil {
		return nil, err
	}
	box.Add(l)

	return box, nil
}

func (u *UI) getRightBox() (*gtk.Box, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 6)
	if err != nil {
		return nil, err
	}

	text, err := gtk.TextViewNew()
	if err != nil {
		return nil, err
	}

	buf, err := text.GetBuffer()
	if err != nil {
		return nil, err
	}
	buf.SetText("\n/*\n * Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>\n *\n * This file originated from: http://opensource.conformal.com/\n *\n * Permission to use, copy, modify, and distribute this software for any\n * purpose with or without fee is hereby granted, provided that the above\n * copyright notice and this permission notice appear in all copies.\n *\n * THE SOFTWARE IS PROVIDED \"AS IS\" AND THE AUTHOR DISCLAIMS ALL WARRANTIES\n * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF\n * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR\n * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES\n * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN\n * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF\n * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.")
	box.Add(text)

	return box, nil
}
