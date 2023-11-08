package tool

import (
	"localdev/HrHelper/internal/config"
	"localdev/HrHelper/internal/util"
	"net/http"
	"time"
)

type Tool struct {
	Config        *config.Config
	Client        *http.Client
	ForumDateTime time.Time
}

func NewTool(config *config.Config, client *http.Client) *Tool {
	forumDateTime, err := util.GetTimeFromTimeZone("America/Mexico_City")
	util.Panic(err)
	return &Tool{
		Config:        config,
		Client:        client,
		ForumDateTime: forumDateTime,
	}
}
