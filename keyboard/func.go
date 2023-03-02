package keyboard

import (
	"errors"
	"github.com/lucasb-eyer/go-colorful"
	"os/exec"
	"strconv"
	"strings"
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

// FindKeyboard find vid/pid keyboard
func FindKeyboard() (uint16, uint16, error) {
	cmd := exec.Command("lsusb")

	b, err := cmd.CombinedOutput()
	if err != nil {
		return 0, 0, err
	}
	out := string(b)

	devicesStrings := strings.Split(out, "\n")

	index := -1
	for idx, str := range devicesStrings {
		if strings.Contains(str, "Winbond Electronics Corp. Gaming Keyboard") {
			index = idx
			break
		}
	}

	if index == -1 {
		return 0, 0, errors.New("keyboard not found")
	}

	spdev := strings.Split(devicesStrings[index], " ")

	ids := strings.Split(spdev[5], ":")
	if len(ids) < 2 {
		return 0, 0, errors.New("keyboard not found")
	}

	vid, err := strconv.ParseUint(ids[0], 16, 64)
	if err != nil {
		return 0, 0, err
	}

	pid, err := strconv.ParseUint(ids[1], 16, 64)
	if err != nil {
		return 0, 0, err
	}

	return uint16(vid), uint16(pid), nil
}
