package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS
var app *App

func main() {
	dir := getAppConfigDir()
	store := initStore(dir)

	go watch()

	// Create an instance of the app structure
	app = NewApp(store)

	// Create application with options
	err := wails.Run(&options.App{
		Title:         appName,
		Width:         960,
		Height:        540,
		Assets:        assets,
		OnStartup:     app.startup,
		OnDomReady:    app.domReady,
		OnBeforeClose: app.beforeClose,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

	defer store.db.Close()
}
