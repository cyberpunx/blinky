package tool

import (
	"fmt"
	"localdev/HrHelper/internal/chronology"
	conf "localdev/HrHelper/internal/config"
	"localdev/HrHelper/internal/hrparse"
	"localdev/HrHelper/internal/potion"
	"localdev/HrHelper/internal/util"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func (o *Tool) ParseSubforum(subHtml string) []*hrparse.Thread {
	threadList := hrparse.GetSubforumThreads(subHtml)

	var threads []*hrparse.Thread
	for _, thread := range threadList {
		threadUrl := hrparse.SubGetThreadUrl(thread)
		threadHtml := o.GetThread(threadUrl)
		thread := o.ParseThread(threadHtml)
		threads = append(threads, thread)
	}
	return threads
}

func (o *Tool) ParseThread(threadHtml string) *hrparse.Thread {
	threadTitle, threadUrl, err := hrparse.ThreadExtractTitleAndURL(threadHtml)
	util.Panic(err)

	var posts []*hrparse.Post
	var pagesUrl []string
	pagesUrl = append(pagesUrl, threadUrl)
	for {
		// Extract the post list from the current page
		postList := hrparse.ThreadListPosts(threadHtml)
		for _, post := range postList {
			post := o.ParsePost(post)
			posts = append(posts, post)
		}

		// Check if there is a "next" link in the pagination
		nextPageURL, hasMore := hrparse.ThreadNextPageURL(threadHtml)

		if !hasMore {
			break // No more pages to fetch
		}

		// Fetch the next page and update the threadHtml
		pagesUrl = append(pagesUrl, nextPageURL)
		nextPageHTML := o.GetThread(nextPageURL)
		threadHtml = nextPageHTML
	}

	if posts == nil || len(posts) == 0 {
		return nil
	}

	firstPostId := posts[0].Id
	var filteredPosts []*hrparse.Post
	filteredPosts = append(filteredPosts, posts[0])
	for _, post := range posts {
		if post.Id != firstPostId {
			filteredPosts = append(filteredPosts, post)
		}
	}

	return &hrparse.Thread{
		Title:    threadTitle,
		Url:      threadUrl,
		Author:   posts[0].Author,
		Created:  posts[0].Created,
		LastPost: posts[len(posts)-1],
		Pages:    pagesUrl,
		Posts:    filteredPosts,
	}
}

func (o *Tool) ParsePost(postHtml string) *hrparse.Post {
	postUser := hrparse.PostGetUserName(postHtml)
	postUserUrl := hrparse.PostGetUserUrl(postHtml)
	postUserHouse := hrparse.PostGetUserHouse(postHtml)
	postDateTime := hrparse.PostGetDateAndTime(postHtml, o.ForumDateTime)
	postEditedDateTime := hrparse.PostGetEditedDateAndTime(postHtml)
	postUrl := hrparse.PostGetUrl(postHtml)
	postContent := hrparse.PostGetContent(postHtml)
	dices := hrparse.ParseDiceRoll(hrparse.PostGetDices(postHtml))

	return &hrparse.Post{
		Url:     postUrl,
		Author:  &hrparse.User{Username: postUser, Url: postUserUrl, House: postUserHouse},
		Created: postDateTime,
		Edited:  postEditedDateTime,
		Content: postContent,
		Dices:   dices,
		Id:      postUrl[strings.LastIndex(postUrl, "#")+1:],
	}
}

func (o *Tool) ProcessPotionsSubforum(subforumThreads []*hrparse.Thread, turnLimit int, timeLimit int) []potion.PotionClubReport {
	fmt.Println("=== Potions Begin ===")
	var reportList []potion.PotionClubReport
	for threadIndex, thread := range subforumThreads {
		fmt.Println("Processing Thread: " + conf.Purple + strconv.Itoa(threadIndex+1) + "/" + strconv.Itoa(len(subforumThreads)) + conf.Reset)
		fmt.Println("Thread: " + conf.Purple + thread.Title + conf.Reset)
		report := potion.ClubPotionsProcessor(*thread, turnLimit, timeLimit, o.ForumDateTime)
		reportList = append(reportList, report)
		fmt.Println("\n")
	}
	fmt.Println("=== Potions End === \n")
	return reportList
}

func (o *Tool) ProcessPotionsThread(thread hrparse.Thread, turnLimit int, timeLimit int) potion.PotionClubReport {
	fmt.Println("=== Potion Thread Begin ===")
	fmt.Println("Thread: " + conf.Purple + thread.Title + conf.Reset)
	var report potion.PotionClubReport
	report = potion.ClubPotionsProcessor(thread, turnLimit, timeLimit, o.ForumDateTime)
	fmt.Println("\n")
	fmt.Println("=== Potion Thread End === \n")
	return report
}

func (o *Tool) ProcessChronoMainThread(chronoMainThread hrparse.Thread, hrTool *Tool) {
	fmt.Println("=== Chronology Thread Begin ===")
	fmt.Println("Thread: " + conf.Purple + chronoMainThread.Title + conf.Reset)

	var chronoLinks []string
	for _, post := range chronoMainThread.Posts {
		chronoLink := hrparse.PostGetLinks(post.Content)
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
		chronoThreadtHtml := hrTool.GetThread(link)
		if hrparse.IsThreadVisible(chronoThreadtHtml) {
			threadListHtml = append(threadListHtml, chronoThreadtHtml)
		}
	}

	var chronoThreads []*chronology.ChronoThread
	for _, threadHtml := range threadListHtml {
		thread := hrTool.ParseThread(threadHtml)
		chronoThread := chronology.ChronoThreadProcessor(*thread)
		chronoThreads = append(chronoThreads, chronoThread)
	}

	chronoReport := chronology.ChronoReport{
		ChronoThreads: chronoThreads,
	}
	//fmt.Printf("%s\n", util.MarshalJsonPretty(chronoReport))

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
