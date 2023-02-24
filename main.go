package main

import (
	"KeyboardDriver/keyboard"
	"log"
	"time"
)

func main() {
	k, err := keyboard.NewKeyboard()
	if err != nil {
		log.Fatalln(err)
	}

	go k.Run()
	time.Sleep(5 * time.Second)
	k.Keymap[5].Red = 255
	k.Keymap[5].Green = 255

	time.Sleep(15 * time.Second)
}
