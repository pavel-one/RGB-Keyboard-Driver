package main

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"log"
	"os"
	"os/exec"
	"os/user"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

const ProjectName = "Dark Project KD87a"

var isRoot bool = false

func checkUdev() {
	defer func() {
		if !isRoot {
			return
		}

		exec.Command("udevadm", "control", "--reload-rules")
		exec.Command("udevadm", "trigger")
		log.Println("Turn off and on keyboard")
		os.Exit(1)
	}()

	u, err := user.Current()
	if err != nil {
		log.Fatalln("Not get current user")
	}

	if u.Uid != "0" {
		isRoot = false
		return
	}

	isRoot = true

	if _, err := os.ReadFile("/etc/udev/rules.d/21-kd87-keyboard.rules"); err == nil { //if file exists, exit
		log.Println("File existing, run from user")
		return
	}

	str := []byte("SUBSYSTEM==\"usb\", ATTR{idVendor}==\"0416\", ATTR{idProduct}==\"c345\", TAG+=\"uaccess\"")
	if err := os.WriteFile("/etc/udev/rules.d/21-kd87-keyboard.rules", str, 0644); err != nil {
		return
	}

	return
}

func main() {
	checkUdev()

	fatalErr := make(chan error, 1)
	go func() {
		log.Fatalln(<-fatalErr)
	}()

	app := NewApp(fatalErr, context.Background())

	err := wails.Run(&options.App{
		Title:  ProjectName,
		Width:  1366,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		//MaxWidth:         1920,
		//MaxHeight:        1024,
		//MinWidth:         1024,
		//MinHeight:        768,
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
