package main

import (
	"KeyboardDriver/keyboard"
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"time"
)

// App struct
type App struct {
	ctx      context.Context
	Keyboard *keyboard.Keyboard
	FatalCh  chan<- error
}

// NewApp creates a new App application struct
func NewApp(ch chan<- error, ctx context.Context) *App {
	k := keyboard.NewKeyboard(ch)

	return &App{
		ctx:      ctx,
		Keyboard: k,
		FatalCh:  ch,
	}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	//a.ctx = ctx

	log.Println("startup")

	if err := a.Keyboard.OpenDevice(); err != nil {
		return
	}

	_, err := a.Keyboard.SetDriverMode()
	if err != nil {
		return
	}

	go a.Keyboard.Run()
}

// domReady is called after front-end resources have been loaded
func (a *App) domReady(ctx context.Context) {
	go a.Keyboard.WelcomeEffect()

	go func() {
		for true {
			time.Sleep(time.Second * 5)
			a.Keyboard.WelcomeEffect()
		}
	}()

	runtime.LogInfo(a.ctx, "Ready")
}

func (a *App) shutdown(ctx context.Context) {
	runtime.LogInfo(a.ctx, "Shutdown")
	if a.Keyboard.Connected {
		a.Keyboard.Close()
	}
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

func (a *App) GetConnectedStatus() bool {
	if !a.Keyboard.Connected {
		if err := a.Keyboard.OpenDevice(); err != nil {
			runtime.LogErrorf(a.ctx, "Error open device: %s", err)
			return false
		}
	}

	return a.Keyboard.Connected
}
