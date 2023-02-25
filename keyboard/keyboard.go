package keyboard

import (
	"github.com/karalabe/hid"
	"time"
)

type Keyboard struct {
	VendorID  uint16
	ProductID uint16
	Device    *hid.Device
	RGBState  []byte
	Keymap    []*Key
	ErrorCh   chan<- error
}

func NewKeyboard(ch chan<- error) (*Keyboard, error) {
	vid := uint16(1046)  //TODO: Set vid
	pid := uint16(49989) //TODO: set pid

	keyboard := &Keyboard{
		VendorID:  vid,
		ProductID: pid,
		ErrorCh:   ch,
	}

	keyboard.RGBState = keyboard.getResetBytes() //set byte map

	keyboard.setKeymap()

	// open
	if err := keyboard.openDevice(); err != nil {
		return nil, err
	}

	if err := keyboard.ResetState(); err != nil {
		return nil, err
	}

	return keyboard, nil
}

func (k *Keyboard) Run() {
	for true {
		_, err := k.UpdateWithKeys()
		if err != nil {
			k.ErrorCh <- err
			return
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func (k *Keyboard) openDevice() error {
	devices := hid.Enumerate(k.VendorID, k.ProductID)
	d, err := devices[2].Open()
	if err != nil {
		return err
	}
	k.Device = d

	return nil
}

func (k *Keyboard) ResetState() error {
	for _, key := range k.Keymap {
		key.Reset()
	}

	return nil
}

func (k *Keyboard) Update() (int, error) {
	return k.Device.Write(k.RGBState)
}

func (k *Keyboard) UpdateWithKeys() (int, error) {
	for _, key := range k.Keymap {
		k.RGBState[key.GetRedIndex()] = key.GetRed()
		k.RGBState[key.GetGreenIndex()] = key.GetGreen()
		k.RGBState[key.GetBlueIndex()] = key.GetBlue()
	}

	return k.Update()
}

func (k *Keyboard) setKeymap() {
	k.Keymap = []*Key{
		NewKey("ESC", 7),
		NewKey("F1", 13),
		NewKey("F2", 16),
		NewKey("F3", 19),
		NewKey("F4", 22),
		NewKey("F5", 28),
		NewKey("F6", 31),
		NewKey("F7", 34),
		NewKey("F8", 37),
		NewKey("F9", 40),
		NewKey("F10", 43),
		NewKey("F11", 46),
		NewKey("F12", 49),
		NewKey("PRTSCR", 52),
		NewKey("SCROLL", 55),
		NewKey("PAUSE", 58),

		NewKey("~", 83),
		NewKey("1", 86),
		NewKey("2", 89),
		NewKey("3", 92),
		NewKey("4", 95),
		NewKey("5", 98),
		NewKey("6", 101),
		NewKey("7", 104),
		NewKey("8", 107),
		NewKey("9", 110),
		NewKey("0", 113),
		NewKey("_", 116),
		NewKey("+", 119),
		NewKey("BACKSPACE", 135),
		NewKey("INS", 138),
		NewKey("HOME", 141),
		NewKey("PGUP", 144),

		NewKey("TAB", 159),
		NewKey("Q", 162),
		NewKey("W", 165),
		NewKey("E", 168),
		NewKey("R", 171),
		NewKey("T", 174),
		NewKey("Y", 177),
		NewKey("U", 180),
		NewKey("I", 183),
		NewKey("O", 186),
		NewKey("P", 199),
		NewKey("{", 202),
		NewKey("}", 205),
		NewKey("|", 211),
		NewKey("DEL", 214),
		NewKey("END", 217),
		NewKey("PGDN", 220),

		NewKey("CAPS", 235),
		NewKey("A", 241),
		NewKey("S", 244),
		NewKey("D", 247),
		NewKey("F", 250),
		NewKey("G", 263),
		NewKey("H", 266),
		NewKey("J", 269),
		NewKey("K", 272),
		NewKey("L", 275),
		NewKey(":", 278),
		NewKey("\"", 281),
		NewKey("ENTER", 287),

		NewKey("LSHIFT", 311),
		NewKey("Z", 327),
		NewKey("X", 330),
		NewKey("C", 333),
		NewKey("V", 336),
		NewKey("B", 339),
		NewKey("N", 342),
		NewKey("M", 345),
		NewKey("<", 348),
		NewKey(">", 351),
		NewKey("?", 354),
		NewKey("RSHIFT", 363),
		NewKey("UP", 369),

		NewKey("LCTRL", 397),
		NewKey("WIN", 400),
		NewKey("LALT", 403),
		NewKey("SPACE", 415),
		NewKey("RALT", 427),
		NewKey("DP", 430),
		NewKey("PRINT", 433),
		NewKey("RCTRL", 436),
		NewKey("LEFT", 442),
		NewKey("DOWN", 455),
		NewKey("RIGHT", 458),
	}
}

func (k *Keyboard) getResetBytes() []byte {
	return []byte{
		0x01, 0x0f, 0x00, 0x00, 0x00, 0x36, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x01, 0x0f, 0x00, 0x00, 0x01, 0x36, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x01, 0x0f, 0x00, 0x00, 0x02, 0x36, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x01, 0x0f, 0x00, 0x00, 0x03, 0x36, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x01, 0x0f, 0x00, 0x00, 0x04, 0x36, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x01, 0x0f, 0x00, 0x00, 0x05, 0x36, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x01, 0x0f, 0x00, 0x00, 0x06, 0x36, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x01, 0x0f, 0x00, 0x00, 0x07, 0x12, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}
}
