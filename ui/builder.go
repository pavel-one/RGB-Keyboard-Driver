package ui

import "github.com/gotk3/gotk3/gtk"

func createGrid() (*gtk.Grid, error) {
	g, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}

	g.SetHExpand(true)

	return g, nil
}

func createSidebar() (*gtk.Box, error) {
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

func createRightBox() (*gtk.Box, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	if err != nil {
		return nil, err
	}

	box.SetHExpand(true)

	text, err := gtk.TextViewNew()
	if err != nil {
		return nil, err
	}

	buf, err := text.GetBuffer()
	if err != nil {
		return nil, err
	}
	buf.SetText("\n/*\n * Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>\n *\n * This file originated from: http://opensource.conformal.com/\n *\n * Permission to use, copy, modify, and distribute this software for any\n * purpose with or without fee is hereby granted, provided that the above\n * copyright notice and this permission notice appear in all copies.\n *\n * THE SOFTWARE IS PROVIDED \"AS IS\" AND THE AUTHOR DISCLAIMS ALL WARRANTIES\n * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF\n * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR\n * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES\n * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN\n * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF\n * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.")

	btn, _ := gtk.ButtonNew()
	btn.SetLabel("TEST")

	box.Add(text)
	box.Add(btn)

	return box, nil
}
