package main

import (
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"localdev/HrHelper/internal/config"
	"localdev/HrHelper/internal/endpoint"
	"localdev/HrHelper/internal/gsheet"
	"localdev/HrHelper/internal/hogwartsforum/tool"
	"localdev/HrHelper/internal/storage"
	"localdev/HrHelper/internal/util"
)

//go:embed all:frontend/dist
var assets embed.FS

const DEBUG = true

func main() {
	//logpanic.InitPanicFile()

	// Create an instance of the app structure

	db := storage.InitDB()
	defer db.Close()
	conf := storage.GetConfig(db)
	config.InitUnicodeConfig(conf)

	if DEBUG {
		client, loginResponse := tool.LoginAndGetCookies(*conf.Username, *conf.Password)
		if !*loginResponse.Success {
			fmt.Println("Not logged in. Exiting...")
			return
		}
		sheetService := gsheet.GetSheetService(*conf.GSheetTokenFile, *conf.GSheetCredFile)
		hrTool := tool.NewTool(conf, client, sheetService)
		forumDateTime, err := util.GetTimeFromTimeZone("America/Mexico_City")
		util.Panic(err)
		fmt.Println("Forum Datetime: " + config.Purple + forumDateTime.Format("01/02/2006 15:04") + config.Reset + "\n")

		endpoints := endpoint.NewEndpoints(hrTool)
		endpoints.ConfigureAndServeEndpoints()

		select {}
	} else {
		app := NewApp(db)

		// Create application with options
		err2 := wails.Run(&options.App{
			Title:  "HrHelper",
			Width:  1366,
			Height: 768,
			AssetServer: &assetserver.Options{
				Assets: assets,
			},
			BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
			OnStartup:        app.startup,
			Bind: []interface{}{
				app,
			},
			Windows: &windows.Options{
				WindowIsTranslucent: true,
				BackdropType:        windows.Acrylic,
				Theme:               windows.Dark,
			},
		})

		if err2 != nil {
			println("Error:", err2.Error())
		}
	}

}
