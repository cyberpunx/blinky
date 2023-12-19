package tool

import (
	"google.golang.org/api/sheets/v4"
	"localdev/HrHelper/internal/config"
	"localdev/HrHelper/internal/util"
	"net/http"
	"time"
)

type LoginRequest struct {
	User *string `json:"user"`
	Pass *string `json:"pass"`
}

type LoginResponse struct {
	Success  *bool   `json:"success"`
	Messaage *string `json:"message"`
	Username *string `json:"username"`
	Initials *string `json:"initials"`
}

type Tool struct {
	Config        *config.Config
	Client        *http.Client
	SheetService  *sheets.Service
	ForumDateTime time.Time
}

func NewTool(config *config.Config, client *http.Client, gSheetService *sheets.Service) *Tool {
	forumDateTime, err := util.GetTimeFromTimeZone("America/Mexico_City")
	util.Panic(err)
	return &Tool{
		Config:        config,
		Client:        client,
		ForumDateTime: forumDateTime,
		SheetService:  gSheetService,
	}
}
