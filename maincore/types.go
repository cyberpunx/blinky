package main

import (
	"localdev/HrHelper/util"
	"net/http"
	"time"
)

type Config struct {
	Username       *string `json:"username" meta-obscure:"default"`
	Password       *string `json:"password" meta-obscure:"default"`
	BaseUrl        *string `json:"baseUrl" meta-obscure:"default"`
	PotionsClubUrl *string `json:"potionsClubUrl" meta-obscure:"default"`
	UnicodeOutput  *bool   `json:"unicodeOutput" meta-obscure:"default"`
}

type Tool struct {
	Config        *Config
	Client        *http.Client
	ForumDateTime time.Time
}

func NewTool(config *Config, client *http.Client) *Tool {
	forumDateTime, error := util.GetTimeFromTimeZone("America/Mexico_City")
	util.Panic(error)
	return &Tool{
		Config:        config,
		Client:        client,
		ForumDateTime: forumDateTime,
	}
}

type Thread struct {
	Title          string
	Url            string
	Author         *User
	Created        *time.Time
	LastPost       *Post
	SecondLastPost *Post
	Pages          []string
	Posts          []*Post
}

type Post struct {
	Url     string
	Author  *User
	Created *time.Time
	Edited  *time.Time
	Content string
	Id      string
}

type User struct {
	Username string
	Url      string
	House    string
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
