package keyboard

import "github.com/lucasb-eyer/go-colorful"

type Key struct {
	Address []int
	Name    string
	Red     int
	Green   int
	Blue    int
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

func (k *Key) Fill(c *colorful.Color) {
	k.Red = int(c.R)
	k.Green = int(c.G)
	k.Blue = int(c.B)
}

func (k *Key) Reset() {
	k.Fill(&colorful.Color{})
}

func (k *Key) GetRedIndex() int {
	return k.Address[0] - 1
}

func (k *Key) GetGreenIndex() int {
	return k.Address[1] - 1
}

func (k *Key) GetBlueIndex() int {
	return k.Address[2] - 1
}

func (k *Key) GetRed() byte {
	return byte(k.Red)
}

func (k *Key) GetGreen() byte {
	return byte(k.Green)
}

func (k *Key) GetBlue() byte {
	return byte(k.Blue)
}
