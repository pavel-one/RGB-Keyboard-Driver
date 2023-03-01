package keyboard

import (
	"github.com/lucasb-eyer/go-colorful"
	"time"
)

// RandomHappyColor Create random happy color
func RandomHappyColor() *colorful.Color {
	r, g, b := colorful.FastHappyColor().RGB255()
	return &colorful.Color{R: float64(r), G: float64(g), B: float64(b)}
}

// FakeButton this func simulation button painting
func FakeButton() {
	time.Sleep(time.Millisecond * 100) // FillSmooth method sleep 100ms
}
