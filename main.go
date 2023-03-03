package main

import (
	"KeyboardDriver/keyboard"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"log"
	"time"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

const ProjectName = "Dark Project KD87a"

func main() {
	fatalErr := make(chan error, 1)
	go func() {
		log.Fatalln(<-fatalErr)
	}()

	// construct keyboard
	k := keyboard.NewKeyboard(fatalErr)
	time.Sleep(time.Millisecond * 500)
	go k.Run() //keyboard worker

	app := NewApp(k, fatalErr)

	err := wails.Run(&options.App{
		Title:  ProjectName,
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		OnStartup:        app.startup,
		OnDomReady:       app.domReady,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatalln(err)
	}
}
