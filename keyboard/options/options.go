package options

import "github.com/lucasb-eyer/go-colorful"

type Options struct {
	Mode       uint16          `json:"mode"`
	Brightness uint16          `json:"bright"`
	Speed      uint16          `json:"speed"`
	BackColor  *colorful.Color `json:"back_color"`
	Rainbow    bool            `json:"rainbow"`
}

func NewOptionsDefault() *Options {
	o := &Options{}
	SetModeDefault(o)
	SetSpeedDefault(o)

	o.BackColor = &colorful.Color{
		R: 255,
		G: 0,
		B: 0,
	}

	return o
}

func (o *Options) GetBytes() []byte {
	b := []byte{
		0x01, 0x07, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x04, //7 byte - mode (max 16), 8 - bright (max - 4)
		0x03, 0xff, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, // 9 - speed, 13,14,15 - back color
		0x00, // 17 - byte type (rainbow or mono)
	}

	b[6] = byte(o.Mode)
	b[7] = byte(o.Brightness)
	b[8] = byte(o.Speed)
	b[12] = byte(o.BackColor.R)
	b[13] = byte(o.BackColor.G)
	b[14] = byte(o.BackColor.B)
	if o.Rainbow {
		b[16] = 0x01
	} else {
		b[16] = 0x00
	}

	return b
}
