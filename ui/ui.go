package ui

import (
	"KeyboardDriver/keyboard"
	"github.com/gotk3/gotk3/gtk"
)

type UI struct {
	Keyboard *keyboard.Keyboard
}

func NewUi(k *keyboard.Keyboard) *UI {
	gtk.Init(nil)

	window, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Example")
	//window.SetDefaultSize(600, 600)

	g, _ := gtk.GridNew()
	//g.SetVAlign(gtk.AlignCenter)
	//g.SetVExpand(true)
	//g.SetHExpand(true)
	g.SetColumnSpacing(20)
	g.SetRowSpacing(20)

	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	box.SetVAlign(gtk.ALIGN_CENTER)
	box.SetSpacing(10)
	b, _ := gtk.ButtonNewFromIconName("go-home", gtk.ICON_SIZE_LARGE_TOOLBAR)
	box.Add(b)
	b1, _ := gtk.ButtonNewFromIconName("starred", gtk.ICON_SIZE_LARGE_TOOLBAR)
	box.Add(b1)
	l, _ := gtk.LabelNew("TEST!")
	box.Add(l)

	box2, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 6)
	text, _ := gtk.TextViewNew()
	buf, _ := text.GetBuffer()
	buf.SetText("\n/*\n * Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>\n *\n * This file originated from: http://opensource.conformal.com/\n *\n * Permission to use, copy, modify, and distribute this software for any\n * purpose with or without fee is hereby granted, provided that the above\n * copyright notice and this permission notice appear in all copies.\n *\n * THE SOFTWARE IS PROVIDED \"AS IS\" AND THE AUTHOR DISCLAIMS ALL WARRANTIES\n * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF\n * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR\n * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES\n * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN\n * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF\n * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.")
	box2.Add(text)

	//cw, _ := gtk.ColorChooserDialogNew("TEST COLOR", window)

	g.Attach(box, 1, 1, 1, 1)
	g.Attach(box2, 2, 1, 1, 1)

	window.Add(g)

	window.ShowAll()

	gtk.Main()

	return &UI{
		Keyboard: k,
	}
}
