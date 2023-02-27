package main

import (
	"KeyboardDriver/keyboard"
	"KeyboardDriver/ui"
	"log"
)

func main() {
	fatalErr := make(chan error, 1)

	k, err := keyboard.NewKeyboard(fatalErr)
	if err != nil {
		log.Fatalln(err)
	}

	go k.Run() //worker
	go k.WelcomeEffect()

	ui.NewUi(k)

	log.Println("Ready")
	err = <-fatalErr
	log.Fatalln(err)
}
