package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"localdev/HrHelper/util"
	"os"
	"path/filepath"
	"strconv"
)

func LoadConfig(configPath string, config interface{}) {
	executablePath, err := os.Executable()
	parentPath := filepath.Dir(executablePath)
	//abs, err := filepath.Abs(path)
	//util.Panic(err)
	b, err := ioutil.ReadFile(filepath.Join(parentPath, configPath))
	util.Panic(err)
	err = json.Unmarshal(b, config)
	util.Panic(err)
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

func printResponseStatus(status string) {
	statusColor := ""
	statusEmoji := ""
	if status == "200 OK" {
		statusColor = Green
		statusEmoji = " " + CheckEmoji + " "
	} else {
		statusColor = Red
		statusEmoji = " " + CrossEmoji + " "
	}
	fmt.Println("Response Status: " + statusColor + statusEmoji + " " + status + Reset)
}

func getConfig() *Config {
	return config
}

func main() {
	config := getConfig()
	//util.ConfigLoggers("reporte.log", 2000000, 10, false, []string{LogTagInfo, LogTagPotions}...)
	fmt.Println(" === 💫 ¡BLINKY A SU SERVICIO! 💫 ===")

	user := *config.Username
	pass := *config.Password
	client := loginAndGetCookies(user, pass)
	hrTool := NewTool(config, client)
	forumDateTime, err := util.GetTimeFromTimeZone("America/Mexico_City")
	util.Panic(err)
	fmt.Println("Forum Datetime: " + Purple + forumDateTime.Format("01/02/2006 15:04") + Reset + "\n")

	tasks := config.Tasks
	for _, task := range tasks {
		taskUrls := *task.Urls
		taskMethod := *task.Method
		timeLimit := *task.TimeLimit

		switch taskMethod {
		case "subforumPotionsClub":
			fmt.Println("\n\n ========= SUBFORUM CLUB DE POCIONES =========\n\n")
			for _, taskUrl := range taskUrls {
				fmt.Println("=== Fetching Subforum === \n")
				potionSubHtml := hrTool.getSubforum(taskUrl)
				subforumThreads := hrTool.parseSubforum(potionSubHtml)
				fmt.Println("=== Fetch Ended === \n")
				hrTool.ProcessPotionsSubforum(subforumThreads, timeLimit)
			}
		case "threadPotionsClub":
			fmt.Println("\n\n ========= THREADS CLUB DE POCIONES =========\n\n")
			for _, taskUrl := range taskUrls {
				potionThreadHtml := hrTool.getThread(taskUrl)
				potionThread := hrTool.parseThread(potionThreadHtml)
				hrTool.ProcessPotionsThread(*potionThread, timeLimit)
			}
		}

	}
}

func (o *Tool) ProcessPotionsSubforum(subforumThreads []*Thread, timeLimit int) {
	fmt.Println("=== Potions Begin ===")
	for threadIndex, thread := range subforumThreads {
		fmt.Println("Processing Thread: " + Purple + strconv.Itoa(threadIndex+1) + "/" + strconv.Itoa(len(subforumThreads)) + Reset)
		fmt.Println("Thread: " + Purple + thread.Title + Reset)
		ClubPotionsProcessor(*thread, timeLimit)
		fmt.Println("\n")
	}
	fmt.Println("=== Potions End === \n")
}

func (o *Tool) ProcessPotionsThread(thread Thread, timeLimit int) {
	fmt.Println("=== Potion Thread Begin ===")
	fmt.Println("Thread: " + Purple + thread.Title + Reset)
	ClubPotionsProcessor(thread, timeLimit)
	fmt.Println("\n")
	fmt.Println("=== Potion Thread End === \n")
}
