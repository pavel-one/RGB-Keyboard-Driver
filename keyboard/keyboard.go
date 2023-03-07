package keyboard

import (
	"errors"
	"github.com/karalabe/hid"
	"sync"
	"time"
)

type Keyboard struct {
	VendorID     uint16
	ProductID    uint16
	Device       *hid.Device
	RGBState     []byte
	Keymap       []*Key
	KeymapMatrix [][]*Key
	ErrorCh      chan<- error
	Mu           sync.Mutex
	Connected    bool
}

func NewKeyboard(ch chan<- error) *Keyboard {
	//vid, pid, err := FindKeyboard()
	//if err != nil {
	//	ch <- err
	//	return nil
	//}

	keyboard := &Keyboard{
		VendorID:  uint16(1046),
		ProductID: uint16(49989),
		ErrorCh:   ch,
		Mu:        sync.Mutex{},
	}

	keyboard.RGBState = keyboard.getColorBytes() //set byte map

	keyboard.setKeymap()
	keyboard.setMatrix()

	// open
	//if err := keyboard.OpenDevice(); err != nil {
	//	ch <- err
	//	return nil
	//}

	//wi, err := keyboard.SetDriverMode()
	//if err != nil {
	//	ch <- err
	//	return nil
	//}

	if err := keyboard.ResetState(); err != nil {
		ch <- err
		return nil
	}

	return keyboard
}

// SetDriverMode Send change mode keyboard
func (k *Keyboard) SetDriverMode() (int, error) {
	// MODES:
	// 1 - simple color change with dark
	// 2 - rainbow
	// 3 - simple color change smooth
	// 4 - circle
	// 5 - wtf ?
	// 6 - click color
	// 7 - click line
	// 8 - click feel
	// 9 - random color
	// 10 - static manual colors (driver mode)
	// 11 - click plus
	// 12 - equalizer clicked
	// 13 - equalizer 2 (?) music
	// 14 - auto diagonal feel
	// 15 - auto lines feel
	//16 - click circle

	if !k.Connected {
		return 0, errors.New("device not connectd")
	}

	return k.Device.Write([]byte{
		0x01, 0x07, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x04, //7 byte - mode (max 16), 8 - bright (max - 4)
		0x03, 0xff, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, // 9 - speed, 13,14,15 - back color
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // 17 - byte type (rainbow or mono)
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	})
}

// Run update state every 10ms
func (k *Keyboard) Run() {
	for true {
		if !k.Connected {
			time.Sleep(time.Second * 5)
			continue
		}

		if _, err := k.UpdateWithKeys(); err != nil {
			k.ErrorCh <- err
			return
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func (k *Keyboard) OpenDevice() error {
	devices := hid.Enumerate(k.VendorID, k.ProductID)
	if len(devices) == 0 {
		return errors.New("not found devices")
	}
	d, err := devices[2].Open()
	if err != nil {
		return err
	}
	k.Device = d

	k.Connected = true
	return nil
}

// ResetState Shutdown all keys
func (k *Keyboard) ResetState() error {
	for _, key := range k.Keymap {
		key.Reset()
	}

	return nil
}

func (k *Keyboard) write() (int, error) {
	if k.Device == nil {
		return 0, nil
	}

	k.Mu.Lock()
	write, err := k.Device.Write(k.RGBState)
	k.Mu.Unlock()
	return write, err
}

// UpdateWithKeys Update state keys on keyboard
func (k *Keyboard) UpdateWithKeys() (int, error) {
	for _, key := range k.Keymap {
		k.RGBState[key.GetRedIndex()] = key.GetRed()
		k.RGBState[key.GetGreenIndex()] = key.GetGreen()
		k.RGBState[key.GetBlueIndex()] = key.GetBlue()
	}

	return k.write()
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

func (k *Keyboard) setMatrix() {
	k.KeymapMatrix = k.getSliceKeymap()
}

// GetSliceKeymap Get button matrix
func (k *Keyboard) getSliceKeymap() [][]*Key {
	keymap := make([][]*Key, 6)

	keymap[0] = []*Key{
		k.Keymap[0],
		nil,
		k.Keymap[1],
		k.Keymap[2],
		k.Keymap[3],
		k.Keymap[4],
		k.Keymap[5],
		k.Keymap[6],
		k.Keymap[7],
		k.Keymap[8],
		k.Keymap[9],
		k.Keymap[10],
		k.Keymap[11],
		k.Keymap[12],
		k.Keymap[13],
		k.Keymap[14],
		k.Keymap[15],
	}

	keymap[1] = []*Key{
		k.Keymap[16],
		k.Keymap[17],
		k.Keymap[18],
		k.Keymap[19],
		k.Keymap[20],
		k.Keymap[21],
		k.Keymap[22],
		k.Keymap[23],
		k.Keymap[24],
		k.Keymap[25],
		k.Keymap[26],
		k.Keymap[27],
		k.Keymap[28],
		k.Keymap[29],
		k.Keymap[30],
		k.Keymap[31],
		k.Keymap[32],
	}

	keymap[2] = []*Key{
		k.Keymap[33],
		k.Keymap[34],
		k.Keymap[35],
		k.Keymap[36],
		k.Keymap[37],
		k.Keymap[38],
		k.Keymap[39],
		k.Keymap[40],
		k.Keymap[41],
		k.Keymap[42],
		k.Keymap[43],
		k.Keymap[44],
		k.Keymap[45],
		k.Keymap[46],
		k.Keymap[47],
		k.Keymap[48],
		k.Keymap[49],
	}

	keymap[3] = []*Key{
		k.Keymap[50],
		k.Keymap[51],
		k.Keymap[52],
		k.Keymap[53],
		k.Keymap[54],
		k.Keymap[55],
		k.Keymap[56],
		k.Keymap[57],
		k.Keymap[58],
		k.Keymap[59],
		k.Keymap[60],
		k.Keymap[61],
		nil,
		k.Keymap[62],
		nil,
		nil,
		nil,
	}

	keymap[4] = []*Key{
		k.Keymap[63],
		k.Keymap[64],
		k.Keymap[65],
		k.Keymap[66],
		k.Keymap[67],
		k.Keymap[68],
		k.Keymap[69],
		k.Keymap[70],
		k.Keymap[71],
		k.Keymap[72],
		k.Keymap[73],
		nil,
		nil,
		k.Keymap[74],
		nil,
		k.Keymap[75],
		nil,
	}

	keymap[5] = []*Key{
		k.Keymap[76],
		k.Keymap[77],
		k.Keymap[78],
		nil,
		nil,
		nil,
		k.Keymap[79],
		nil,
		nil,
		nil,
		k.Keymap[80],
		k.Keymap[81],
		k.Keymap[82],
		k.Keymap[83],
		k.Keymap[84],
		k.Keymap[85],
		k.Keymap[86],
	}

	return keymap
}

func (k *Keyboard) getColorBytes() []byte {
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

func (k *Keyboard) Close() {
	if err := k.Device.Close(); err != nil {
		k.ErrorCh <- err
		return
	}

	k.Connected = false
}
