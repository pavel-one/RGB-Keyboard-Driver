package main

import (
	"KeyboardDriver/keyboard"
	"log"
)

func main() {
	fatalErr := make(chan error, 1)

	k, err := keyboard.NewKeyboard(fatalErr)
	if err != nil {
		log.Fatalln(err)
	}

	go k.Run() //worker
	k.WelcomeEffect()

	log.Println("Ready")
	err = <-fatalErr
	log.Fatalln(err)
}
