package main

import (
	"fmt"
	conf "localdev/HrHelper/internal/config"
	"localdev/HrHelper/internal/endpoint"
	"localdev/HrHelper/internal/tool"
	"localdev/HrHelper/internal/util"
)

func main() {
	config := conf.GetConfig()
	//util.ConfigLoggers("reporte.log", 2000000, 10, false, []string{LogTagInfo, LogTagPotions}...)
	fmt.Println(" === 💫 ¡BLINKY A SU SERVICIO! 💫 ===")

	user := *config.Username
	pass := *config.Password
	client := tool.LoginAndGetCookies(user, pass)
	hrTool := tool.NewTool(config, client)
	forumDateTime, err := util.GetTimeFromTimeZone("America/Mexico_City")
	util.Panic(err)
	fmt.Println("Forum Datetime: " + conf.Purple + forumDateTime.Format("01/02/2006 15:04") + conf.Reset + "\n")

	endpoints := endpoint.NewEndpoints(hrTool)
	endpoints.ConfigureAndServeEndpoints()

	select {}

}

/*
tasks := config.Tasks
	for _, task := range tasks {
		taskUrls := *task.Urls
		taskMethod := *task.Method
		timeLimit := task.TimeLimit
		turnLimit := task.TurnLimit

		switch taskMethod {
		case "subforumPotionsClub":
			fmt.Println("\n\n ========= SUBFORUM CLUB DE POCIONES =========\n\n")
			if len(taskUrls) == 0 {
				fmt.Println("No subforums URLs to process")
			}
			for _, taskUrl := range taskUrls {
				fmt.Println("=== Fetching Subforum === \n")
				potionSubHtml := hrTool.GetSubforum(taskUrl)
				subforumThreads := hrTool.ParseSubforum(potionSubHtml)
				fmt.Println("=== Fetch Ended === \n")
				hrTool.ProcessPotionsSubforum(subforumThreads, *turnLimit, *timeLimit)
			}
		case "threadPotionsClub":
			fmt.Println("\n\n ========= THREADS CLUB DE POCIONES =========\n\n")
			if len(taskUrls) == 0 {
				fmt.Println("No Threads URLs to process")
			}
			for _, taskUrl := range taskUrls {
				potionThreadHtml := hrTool.GetThread(taskUrl)
				potionThread := hrTool.ParseThread(potionThreadHtml)
				hrTool.ProcessPotionsThread(*potionThread, *turnLimit, *timeLimit)
			}
		case "mainThreadChronology":
			fmt.Println("\n\n ========= MAIN THREAD CHRONOLOGY =========\n\n")
			if len(taskUrls) == 0 {
				fmt.Println("No Posts URLs to process")
			}
			for _, taskUrl := range taskUrls {
				chronoMainThreadHtml := hrTool.GetThread(taskUrl)
				chronoMainThread := hrTool.ParseThread(chronoMainThreadHtml)
				hrTool.ProcessChronoMainThread(*chronoMainThread, hrTool)
			}
		case "threadChronology":
			fmt.Println("\n\n ========= THREAD CHRONOLOGY =========\n\n")
			if len(taskUrls) == 0 {
				fmt.Println("No Posts URLs to process")
			}
			for _, taskUrl := range taskUrls {
				threadHtml := hrTool.GetThread(taskUrl)
				thread := hrTool.ParseThread(threadHtml)
				chronoThread := chronology.ChronoThreadProcessor(*thread)
				fmt.Printf("%s\n", util.MarshalJsonPretty(chronoThread))
			}
		}

	}
*/
