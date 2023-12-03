package main

import (
	"embed"
	"fmt"
	conf "localdev/HrHelper/internal/config"
	"localdev/HrHelper/internal/endpoint"
	"localdev/HrHelper/internal/hogwartsforum/tool"
	"localdev/HrHelper/internal/util"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	config := conf.GetConfig()
	//util.ConfigLoggers("reporte.log", 2000000, 10, false, []string{LogTagInfo, LogTagPotions}...)
	fmt.Println(" === 💫 ¡BLINKY A SU SERVICIO! 💫 ===")

	user := *config.Username
	pass := *config.Password
	hrTool := tool.NewTool(config, tool.LoginAndGetCookies(user, pass))
	forumDateTime, err := util.GetTimeFromTimeZone("America/Mexico_City")
	util.Panic(err)
	fmt.Println("Forum Datetime: " + conf.Purple + forumDateTime.Format("01/02/2006 15:04") + conf.Reset + "\n")

	endpoints := endpoint.NewEndpoints(hrTool)
	endpoints.ConfigureAndServeEndpoints()

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "HrHelper",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
