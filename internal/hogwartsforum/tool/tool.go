package tool

import (
	"fmt"
	conf "localdev/HrHelper/internal/config"
	"localdev/HrHelper/internal/gsheet"
	"localdev/HrHelper/internal/hogwartsforum/dynamics/chronology"
	"localdev/HrHelper/internal/hogwartsforum/dynamics/potion"
	parser "localdev/HrHelper/internal/hogwartsforum/parser"
	"localdev/HrHelper/internal/util"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func (o *Tool) parseSubforum(subHtml string) []*parser.Thread {
	threadList := parser.GetSubforumThreads(subHtml)

	var threads []*parser.Thread
	for _, thread := range threadList {
		threadUrl := parser.SubGetThreadUrl(thread)
		threadHtml := o.getThread(threadUrl)
		thread := o.parseThread(threadHtml)
		threads = append(threads, thread)
	}
	return threads
}

func (o *Tool) parseThread(threadHtml string) *parser.Thread {
	threadTitle, threadUrl, err := parser.ThreadExtractTitleAndURL(threadHtml)
	util.Panic(err)

	var posts []*parser.Post
	var pagesUrl []string
	pagesUrl = append(pagesUrl, threadUrl)
	for {
		// Extract the post list from the current page
		postList := parser.ThreadListPosts(threadHtml)
		for _, post := range postList {
			post := o.parsePost(post)
			posts = append(posts, post)
		}

		// Check if there is a "next" link in the pagination
		nextPageURL, hasMore := parser.ThreadNextPageURL(threadHtml)

		if !hasMore {
			break // No more pages to fetch
		}

		// Fetch the next page and update the threadHtml
		pagesUrl = append(pagesUrl, nextPageURL)
		nextPageHTML := o.getThread(nextPageURL)
		threadHtml = nextPageHTML
	}

	if posts == nil || len(posts) == 0 {
		return nil
	}

	firstPostId := posts[0].Id
	var filteredPosts []*parser.Post
	filteredPosts = append(filteredPosts, posts[0])
	for _, post := range posts {
		if post.Id != firstPostId {
			filteredPosts = append(filteredPosts, post)
		}
	}

	return &parser.Thread{
		Title:    threadTitle,
		Url:      threadUrl,
		Author:   posts[0].Author,
		Created:  posts[0].Created,
		LastPost: posts[len(posts)-1],
		Pages:    pagesUrl,
		Posts:    filteredPosts,
	}
}

func (o *Tool) parsePost(postHtml string) *parser.Post {
	postUser := parser.PostGetUserName(postHtml)
	postUserUrl := parser.PostGetUserUrl(postHtml)
	postUserHouse := parser.PostGetUserHouse(postHtml)
	postDateTime := parser.PostGetDateAndTime(postHtml, o.ForumDateTime)
	postEditedDateTime := parser.PostGetEditedDateAndTime(postHtml)
	postUrl := parser.PostGetUrl(postHtml)
	postContent := parser.PostGetContent(postHtml)
	dices := parser.ParseDiceRoll(parser.PostGetDices(postHtml))

	return &parser.Post{
		Url:     postUrl,
		Author:  &parser.User{Username: postUser, Url: postUserUrl, House: postUserHouse},
		Created: postDateTime,
		Edited:  postEditedDateTime,
		Content: postContent,
		Dices:   dices,
		Id:      postUrl[strings.LastIndex(postUrl, "#")+1:],
	}
}

func (o *Tool) processPotionsSubforum(subforumThreads []*parser.Thread, timeLimit, turnLimit int) []potion.PotionClubReport {
	fmt.Println("=== Potions Begin ===")
	var reportList []potion.PotionClubReport
	for threadIndex, thread := range subforumThreads {
		fmt.Println("Processing Thread: " + conf.Purple + strconv.Itoa(threadIndex+1) + "/" + strconv.Itoa(len(subforumThreads)) + conf.Reset)
		report := o.processPotionsThread(*thread, timeLimit, turnLimit)
		reportList = append(reportList, report)
		fmt.Println("\n")
	}
	fmt.Println("=== Potions End === \n")
	return reportList
}

func (o *Tool) processPotionsThread(thread parser.Thread, timeLimit, turnLimit int) potion.PotionClubReport {
	fmt.Println("=== Potion Thread Begin ===")
	fmt.Println("Thread: " + conf.Purple + thread.Title + conf.Reset + " (Time: " + strconv.Itoa(timeLimit) + "| Turn: " + strconv.Itoa(turnLimit) + ")" + "\n")
	var report potion.PotionClubReport
	daysOff := o.getGoogleSheetPotionsDayOff()
	report = potion.PotionGetReportFromThread(thread, timeLimit, turnLimit, o.ForumDateTime, daysOff)
	fmt.Println("\n")
	fmt.Println("=== Potion Thread End === \n")
	return report
}

func (o *Tool) processChronoMainThread(chronoMainThread parser.Thread, hrTool *Tool) {
	fmt.Println("=== Chronology Thread Begin ===")
	fmt.Println("Thread: " + conf.Purple + chronoMainThread.Title + conf.Reset)

	var chronoLinks []string
	for _, post := range chronoMainThread.Posts {
		chronoLink := parser.PostGetLinks(post.Content)
		chronoLinks = append(chronoLinks, chronoLink...)
	}

	re, err := regexp.Compile(`p\d+`)
	if err != nil {
		panic(err)
	}
	var cleanedURLs []string
	for _, link := range chronoLinks {
		parsedURL, err := url.Parse(link)
		util.Panic(err)
		parsedURL.Fragment = ""
		urlWithoutFragment := parsedURL.String()
		cleanedUrl := re.ReplaceAllString(urlWithoutFragment, "")
		cleanedURLs = append(cleanedURLs, cleanedUrl)
	}

	var threadListHtml []string
	for _, link := range cleanedURLs {
		chronoThreadtHtml := hrTool.getThread(link)
		if parser.IsThreadVisible(chronoThreadtHtml) {
			threadListHtml = append(threadListHtml, chronoThreadtHtml)
		}
	}

	var chronoThreads []*chronology.ChronoThread
	for _, threadHtml := range threadListHtml {
		thread := hrTool.parseThread(threadHtml)
		chronoThread := chronology.ChronoThreadProcessor(*thread)
		chronoThreads = append(chronoThreads, chronoThread)
	}

	chronoReport := chronology.ChronoReport{
		ChronoThreads: chronoThreads,
	}

	stringContents := fmt.Sprintf("%s\n", util.MarshalJsonPretty(chronoReport))
	filename := "output.json"

	// Create the file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write the content to file
	_, err = file.WriteString(stringContents)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n")
	fmt.Println("=== Chronology Thread End === \n")
}

func (o *Tool) getGoogleSheetPotionsDayOff() *[]gsheet.DayOff {
	rows, err := gsheet.ReadSheetData(o.SheetService, *o.Config.GSheetId, "Permisos Pociones!A269:B")
	util.Panic(err)
	daysOff := gsheet.ParseDayOff(rows)
	return &daysOff
}

func (o *Tool) ProcessPotionsSubforumList(subForumUrls *[]string, timeLimit, turnLimit *int) []potion.PotionClubReport {
	fmt.Println("\n\n ========= SUBFORUM CLUB DE POCIONES =========\n\n")
	fmt.Println("Time Limit: " + conf.Purple + strconv.Itoa(*timeLimit) + conf.Reset)
	fmt.Println("Turn Limit: " + conf.Purple + strconv.Itoa(*turnLimit) + conf.Reset)
	if len(*subForumUrls) == 0 {
		fmt.Println("No subforums URLs to process")
	}
	var reportMainList []potion.PotionClubReport
	for _, subforumUrl := range *subForumUrls {
		fmt.Println("=== Fetching Subforum === \n")
		potionSubHtml := o.getSubforum(subforumUrl)
		subforumThreads := o.parseSubforum(potionSubHtml)
		fmt.Println("=== Fetch Ended === \n")
		reportList := o.processPotionsSubforum(subforumThreads, *timeLimit, *turnLimit)
		reportMainList = append(reportMainList, reportList...)
	}

	return reportMainList
}

func (o *Tool) ProcessPotionsThreadList(threadsUrls *[]string, timeLimit, turnLimit *int) []potion.PotionClubReport {
	fmt.Println("\n\n ========= THREADS DE POCIONES =========\n\n")
	if len(*threadsUrls) == 0 {
		fmt.Println("No Threads URLs to process")
	}
	var reportMainList []potion.PotionClubReport
	for _, threadUrl := range *threadsUrls {
		potionThreadHtml := o.getThread(threadUrl)
		potionThread := o.parseThread(potionThreadHtml)
		report := o.processPotionsThread(*potionThread, *timeLimit, *turnLimit)
		reportMainList = append(reportMainList, report)
	}

	return reportMainList
}
