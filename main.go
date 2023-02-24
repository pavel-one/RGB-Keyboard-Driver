package main

import (
	"KeyboardDriver/keyboard"
	"log"
	"time"
)

func main() {
	_, err := keyboard.NewKeyboard()
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(15 * time.Second)
}
