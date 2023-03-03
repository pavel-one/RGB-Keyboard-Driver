package main

import (
	"KeyboardDriver/keyboard"
	"context"
	"log"
)

// App struct
type App struct {
	ctx      context.Context
	Keyboard *keyboard.Keyboard
	FatalCh  chan<- error
}

// NewApp creates a new App application struct
func NewApp(k *keyboard.Keyboard, ch chan<- error) *App {
	return &App{
		Keyboard: k,
		FatalCh:  ch,
	}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a *App) domReady(ctx context.Context) {
	go a.Keyboard.WelcomeEffect()
	log.Println("Ready")
}

func (a *App) shutdown(ctx context.Context) {
	log.Println("shutdown!!!")
	a.Keyboard.Close()
}

func (a *App) GetKeyboardKeys() []*keyboard.Key {
	return a.Keyboard.Keymap
}

func (a *App) GetKeyboardMatrix() [][]*keyboard.Key {
	return a.Keyboard.KeymapMatrix
}

func (a *App) Reload() {
	go a.Keyboard.WelcomeEffect()

}
