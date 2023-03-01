package ui

import "github.com/gotk3/gotk3/gtk"

func constructUi(window *gtk.Window) error {
	//cw, _ := gtk.ColorChooserDialogNew("TEST COLOR", window)

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

	return nil
}

func createGrid() (*gtk.Grid, error) {
	g, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}

	g.SetColumnSpacing(20)
	g.SetRowSpacing(20)

	return g, nil
}

func createSidebar() (*gtk.Box, error) {
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

func createRightBox() (*gtk.Box, error) {
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
