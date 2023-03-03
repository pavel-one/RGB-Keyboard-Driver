package keyboard

import (
	"github.com/lucasb-eyer/go-colorful"
	"time"
)

type Key struct {
	Address []int  `json:"address"`
	Name    string `json:"name"`
	Red     int    `json:"red"`
	Green   int    `json:"green"`
	Blue    int    `json:"blue"`
}

func NewKey(name string, firstByte int) *Key {
	return &Key{
		Address: []int{
			firstByte,
			firstByte + 1,
			firstByte + 2,
		},
		Name: name,
	}
}

// Fill set color button
func (k *Key) Fill(c *colorful.Color) {
	k.Red = int(c.R)
	k.Green = int(c.G)
	k.Blue = int(c.B)
}

// FillSmooth Coloring the button Smoothly
func (k *Key) FillSmooth(c *colorful.Color) {
	old := colorful.Color{R: float64(k.Red), G: float64(k.Green), B: float64(k.Blue)}
	c1 := c.BlendRgb(old, 0.7)
	c2 := c.BlendRgb(old, 0.3)

	k.Fill(&c1)
	time.Sleep(time.Millisecond * 50)
	k.Fill(&c2)
	time.Sleep(time.Millisecond * 50)
	k.Fill(c)
}

func (k *Key) Reset() {
	k.Fill(&colorful.Color{})
}

// GetRedIndex Get red byte index
func (k *Key) GetRedIndex() int {
	return k.Address[0] - 1
}

// GetGreenIndex Get green byte index
func (k *Key) GetGreenIndex() int {
	return k.Address[1] - 1
}

// GetBlueIndex Get blue byte index
func (k *Key) GetBlueIndex() int {
	return k.Address[2] - 1
}

// GetRed Get red color in byte
func (k *Key) GetRed() byte {
	return byte(k.Red)
}

// GetGreen Get green color in byte
func (k *Key) GetGreen() byte {
	return byte(k.Green)
}

// GetBlue Get blue color in byte
func (k *Key) GetBlue() byte {
	return byte(k.Blue)
}
