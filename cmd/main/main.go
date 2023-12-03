package main

import (
	"fmt"
	conf "localdev/HrHelper/internal/config"
	"localdev/HrHelper/internal/endpoint"
	tool2 "localdev/HrHelper/internal/hogwartsforum/tool"
	"localdev/HrHelper/internal/util"
)

func main() {
	config := conf.GetConfig()
	//util.ConfigLoggers("reporte.log", 2000000, 10, false, []string{LogTagInfo, LogTagPotions}...)
	fmt.Println(" === 💫 ¡BLINKY A SU SERVICIO! 💫 ===")

	user := *config.Username
	pass := *config.Password
	client := tool2.LoginAndGetCookies(user, pass)
	hrTool := tool2.NewTool(config, client)
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
				potionSubHtml := hrTool.getSubforum(taskUrl)
				subforumThreads := hrTool.parseSubforum(potionSubHtml)
				fmt.Println("=== Fetch Ended === \n")
				hrTool.processPotionsSubforum(subforumThreads, *turnLimit, *timeLimit)
			}
		case "threadPotionsClub":
			fmt.Println("\n\n ========= THREADS CLUB DE POCIONES =========\n\n")
			if len(taskUrls) == 0 {
				fmt.Println("No Threads URLs to process")
			}
			for _, taskUrl := range taskUrls {
				potionThreadHtml := hrTool.getThread(taskUrl)
				potionThread := hrTool.parseThread(potionThreadHtml)
				hrTool.processPotionsThread(*potionThread, *turnLimit, *timeLimit)
			}
		case "mainThreadChronology":
			fmt.Println("\n\n ========= MAIN THREAD CHRONOLOGY =========\n\n")
			if len(taskUrls) == 0 {
				fmt.Println("No Posts URLs to process")
			}
			for _, taskUrl := range taskUrls {
				chronoMainThreadHtml := hrTool.getThread(taskUrl)
				chronoMainThread := hrTool.parseThread(chronoMainThreadHtml)
				hrTool.processChronoMainThread(*chronoMainThread, hrTool)
			}
		case "threadChronology":
			fmt.Println("\n\n ========= THREAD CHRONOLOGY =========\n\n")
			if len(taskUrls) == 0 {
				fmt.Println("No Posts URLs to process")
			}
			for _, taskUrl := range taskUrls {
				threadHtml := hrTool.getThread(taskUrl)
				thread := hrTool.parseThread(threadHtml)
				chronoThread := chronology.ChronoThreadProcessor(*thread)
				fmt.Printf("%s\n", util.MarshalJsonPretty(chronoThread))
			}
		}

	}
*/
