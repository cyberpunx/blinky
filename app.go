package main

import (
	"context"
	"fmt"
	conf "localdev/HrHelper/internal/config"
	"localdev/HrHelper/internal/hogwartsforum/tool"
)

// App struct
type App struct {
	tool *tool.Tool
	ctx  context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Login(user, pass string) bool {
	config := conf.GetConfig()
	client, isLoggedIn := tool.LoginAndGetCookies(user, pass)
	if !isLoggedIn {
		fmt.Println("Not logged in. Exiting...")
		return false
	}
	hrTool := tool.NewTool(config, client)
	a.tool = hrTool
	return true
}
