package keyboard

import "time"

func (k *Keyboard) WelcomeEffect() {
	n := 3
	sleepTime := 100 * time.Millisecond

	for i := 0; i < n; i++ {
		k.RGBState = k.getResetBytes()
		k.Update()
		time.Sleep(sleepTime)
		k.RGBState = k.getDemoData()
		k.Update()
		time.Sleep(sleepTime)
	}
}

func (k *Keyboard) Fill(r int, g int, b int) error {
	for _, key := range k.Keymap {
		key.Red = r
		key.Green = g
		key.Blue = b
	}

	_, err := k.UpdateWithKeys()
	if err != nil {
		return err
	}

	return nil
}
