package config

type Config struct {
	Username      *string `json:"username" meta-obscure:"default"`
	Password      *string `json:"password" meta-obscure:"default"`
	Remember      *bool   `json:"remember" meta-obscure:"default"`
	BaseUrl       *string `json:"baseUrl" meta-obscure:"default"`
	UnicodeOutput *bool   `json:"unicodeOutput" meta-obscure:"default"`
	Tasks         []*Task `json:"tasks" meta-obscure:"default"`
}

type Task struct {
	Urls      *[]string `json:"urls" meta-obscure:"default"`
	Method    *string   `json:"method" meta-obscure:"default"`
	TimeLimit *int      `json:"timeLimit" meta-obscure:"default"`
	TurnLimit *int      `json:"turnLimit" meta-obscure:"default"`
}

type PotionSubforumConfig struct {
	Url       *string `json:"url" meta-obscure:"default"`
	TimeLimit *int    `json:"timeLimit" meta-obscure:"default"`
	TurnLimit *int    `json:"turnLimit" meta-obscure:"default"`
}

type PotionThreadConfig struct {
	Url       *string `json:"url" meta-obscure:"default"`
	TimeLimit *int    `json:"timeLimit" meta-obscure:"default"`
	TurnLimit *int    `json:"turnLimit" meta-obscure:"default"`
}

const (
	LogTagInfo    = "Blinky!"
	LogTagPotions = "potionsClub"
	passphrase    = "yourEncryptionKey"
)

var config *Config
var loadedConfigPath string
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
