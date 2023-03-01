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

	u, err := ui.NewUi(k, "Dark Project KD87a")
	if err != nil {
		log.Fatalln(err)
	}

	if err := u.Run(); err != nil {
		log.Fatalln(err)
	}

	log.Println("Ready")
	err = <-fatalErr
	log.Fatalln(err)
}
