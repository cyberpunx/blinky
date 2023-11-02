package main

import (
	"localdev/HrHelper/util"
	"net/http"
	"time"
)

const (
	LogTagInfo    = "Blinky!"
	LogTagPotions = "potionsClub"
)

type Config struct {
	Username      *string `json:"username" meta-obscure:"default"`
	Password      *string `json:"password" meta-obscure:"default"`
	BaseUrl       *string `json:"baseUrl" meta-obscure:"default"`
	UnicodeOutput *bool   `json:"unicodeOutput" meta-obscure:"default"`
	Tasks         []*Task `json:"tasks" meta-obscure:"default"`
}

type Task struct {
	Urls      *[]string `json:"urls" meta-obscure:"default"`
	Method    *string   `json:"method" meta-obscure:"default"`
	TimeLimit *int      `json:"timeLimit" meta-obscure:"default"`
}

type Tool struct {
	Config        *Config
	Client        *http.Client
	ForumDateTime time.Time
}

func NewTool(config *Config, client *http.Client) *Tool {
	forumDateTime, err := util.GetTimeFromTimeZone("America/Mexico_City")
	util.Panic(err)
	return &Tool{
		Config:        config,
		Client:        client,
		ForumDateTime: forumDateTime,
	}
}

type Thread struct {
	Title    string
	Url      string
	Author   *User
	Created  *time.Time
	LastPost *Post
	Pages    []string
	Posts    []*Post
}

type Post struct {
	Url     string
	Author  *User
	Created *time.Time
	Edited  *time.Time
	Content string
	Id      string
	Dices   []*Dice
}

type Dice struct {
	DiceLine string
	Result   int
}

type User struct {
	Username string
	Url      string
	House    string
}

type Potion struct {
	Name         string
	Ingredients  []string
	ScoreTarget  int
	ScoreCurrent int
	TurnTarget   int
	TurnCurrent  int
}

var config *Config
var Reset = ""
var Red = ""
var Green = ""
var Yellow = ""
var Blue = ""
var Purple = ""
var Cyan = ""
var Gray = ""
var White = ""
var CheckEmoji = ""
var CrossEmoji = ""
var RightArrowEmoji = ""
