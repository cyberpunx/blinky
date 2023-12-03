package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LoadConfig(configPath string, config interface{}) {
	executablePath, err := os.Executable()
	parentPath := filepath.Dir(executablePath)
	//abs, err := filepath.Abs(path)
	//util.Panic(err)
	b, err := ioutil.ReadFile(filepath.Join(parentPath, configPath))
	if err != nil {
		b, err = ioutil.ReadFile(filepath.Join(configPath))
		if err != nil {
			panic(err)
		}
	}
	err = json.Unmarshal(b, config)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return config
}

func init() {
	conf := flag.String("conf", "conf.json", "Config")
	LoadConfig(*conf, &config)

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
