package config

import (
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LoadConfig(configPath string, config interface{}) string {
	executablePath, err := os.Executable()
	parentPath := filepath.Dir(executablePath)
	//abs, err := filepath.Abs(path)
	//util.Panic(err)
	loadedPath := filepath.Join(parentPath, configPath)
	b, err := ioutil.ReadFile(filepath.Join(parentPath, configPath))
	if err != nil {
		loadedPath = configPath
		b, err = ioutil.ReadFile(filepath.Join(configPath))
		if err != nil {
			panic(err)
		}
	}
	err = json.Unmarshal(b, config)
	if err != nil {
		panic(err)
	}
	return loadedPath
}

func GetConfig() *Config {
	username := ""
	password := ""
	baseUrl := "https://www.hogwartsrol.com/"
	unicodeOuput := true
	conf := Config{
		Username:      &username,
		Password:      &password,
		BaseUrl:       &baseUrl,
		UnicodeOutput: &unicodeOuput,
		Tasks:         nil,
	}
	return &conf
}

func init() {
	//conf := flag.String("conf", "conf.json", "Config")
	//LoadConfig(*conf, &config)

	config := GetConfig()

	if *config.UnicodeOutput {
		Reset = "\033[0m"
		Red = "\033[31m"
		Green = "\033[32m"
		Yellow = "\033[33m"
		Blue = "\033[34m"
		Purple = "\033[35m"
		Cyan = "\033[36m"
		Gray = "\033[37m"
		White = "\033[97m"
		CheckEmoji = "✔"
		CrossEmoji = "❌"
		RightArrowEmoji = "▶"
	} else {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
		CheckEmoji = "[OK]"
		CrossEmoji = "[X]"
		RightArrowEmoji = "-->"
	}
}
