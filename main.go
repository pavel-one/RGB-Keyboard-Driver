package main

import (
	"KeyboardDriver/keyboard"
	"KeyboardDriver/ui"
	"log"
	"os"
)

func main() {
	fatalErr := make(chan error, 1)

	// construct keyboard
	k, err := keyboard.NewKeyboard(fatalErr)
	if err != nil {
		log.Fatalln(err)
	}

	go k.Run()           //keyboard worker
	go k.WelcomeEffect() //run start effect

	// construct UI
	u, err := ui.NewUI(k, "Dark Project KD87a")
	if err != nil {
		log.Fatalln(err)
	}

	// Run UI
	go u.Application.Run(os.Args)

	log.Println("Ready")
	err = <-fatalErr
	log.Fatalln(err)
}
