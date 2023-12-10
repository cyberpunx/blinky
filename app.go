package main

import (
	"context"
	"database/sql"
	"fmt"
	"localdev/HrHelper/internal/config"
	"localdev/HrHelper/internal/endpoint"
	"localdev/HrHelper/internal/hogwartsforum/tool"
	"localdev/HrHelper/internal/storage"
)

// App struct
type App struct {
	db   *sql.DB
	tool *tool.Tool
	ctx  context.Context
}

// NewApp creates a new App application struct
func NewApp(db *sql.DB) *App {
	config := storage.GetConfig(db)
	tool := tool.NewTool(config, nil)
	return &App{
		tool: tool,
		db:   db,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Login(user, pass string, remeber bool) *tool.LoginResponse {
	config := storage.GetConfig(a.db)

	client, loginResponse := tool.LoginAndGetCookies(user, pass)
	a.tool.Client = client

	if !*loginResponse.Success {
		fmt.Println("Not logged in. Exiting...")
		return loginResponse
	}

	if remeber {
		config.Username = &user
		config.Password = &pass
	} else {
		config.Username = nil
		config.Password = nil
	}
	config.Remember = &remeber
	storage.UpdateConfig(a.db, config)

	hrTool := tool.NewTool(config, client)
	a.tool = hrTool
	return loginResponse
}

func (a *App) SubforumPotionsClub(subforumUrls []string, timeLimit, turnLimit int) *endpoint.SubforumPotionsClubResponse {
	report := a.tool.ProcessPotionsSubforumList(&subforumUrls, &timeLimit, &turnLimit)
	response := endpoint.SubforumPotionsClubResponse{ThreadReports: report}
	return &response
}

func (a *App) GetConfig() *config.Config {
	return a.tool.Config
}

func (a *App) GetTool() *tool.Tool {
	return a.tool
}

func (a *App) GetPotionSubforum() *[]config.PotionSubforumConfig {
	return storage.GetPotionSubforum(a.db)
}

func (a *App) UpdatePotionSubforum(potionSubConfig *[]config.PotionSubforumConfig) {
	storage.UpdatePotionSubforum(a.db, potionSubConfig)
}

func (a *App) GetPotionThread() *[]config.PotionThreadConfig {
	return storage.GetPotionThread(a.db)
}

func (a *App) UpdatePotionThread(potionThrConfig *[]config.PotionThreadConfig) {
	storage.UpdatePotionThread(a.db, potionThrConfig)
}
