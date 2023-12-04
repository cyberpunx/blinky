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
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	config := conf.GetConfig()
	//util.ConfigLoggers("reporte.log", 2000000, 10, false, []string{LogTagInfo, LogTagPotions}...)
	fmt.Println(" === 💫 ¡BLINKY A SU SERVICIO! 💫 ===")

	user := *config.Username
	pass := *config.Password
	client, isLoggedIn := tool.LoginAndGetCookies(user, pass)
	if !isLoggedIn {
		fmt.Println("Not logged in. Exiting...")
		return
	}
	hrTool := tool.NewTool(config, client)
	forumDateTime, err := util.GetTimeFromTimeZone("America/Mexico_City")
	util.Panic(err)
	fmt.Println("Forum Datetime: " + conf.Purple + forumDateTime.Format("01/02/2006 15:04") + conf.Reset + "\n")

	endpoints := endpoint.NewEndpoints(hrTool)
	endpoints.ConfigureAndServeEndpoints()

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err2 := wails.Run(&options.App{
		Title:  "HrHelper",
		Width:  1024,
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
