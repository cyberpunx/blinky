package tool

import (
	"fmt"
	conf "localdev/HrHelper/internal/config"
	"localdev/HrHelper/internal/hrparse"
	"localdev/HrHelper/internal/potion"
	"localdev/HrHelper/internal/util"
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

func (o *Tool) ProcessPotionsSubforum(subforumThreads []*hrparse.Thread, turnLimit int, timeLimit int) {
	fmt.Println("=== Potions Begin ===")
	for threadIndex, thread := range subforumThreads {
		fmt.Println("Processing Thread: " + conf.Purple + strconv.Itoa(threadIndex+1) + "/" + strconv.Itoa(len(subforumThreads)) + conf.Reset)
		fmt.Println("Thread: " + conf.Purple + thread.Title + conf.Reset)
		potion.ClubPotionsProcessor(*thread, turnLimit, timeLimit, o.ForumDateTime)
		fmt.Println("\n")
	}
	fmt.Println("=== Potions End === \n")
}

func (o *Tool) ProcessPotionsThread(thread hrparse.Thread, turnLimit int, timeLimit int) {
	fmt.Println("=== Potion Thread Begin ===")
	fmt.Println("Thread: " + conf.Purple + thread.Title + conf.Reset)
	potion.ClubPotionsProcessor(thread, turnLimit, timeLimit, o.ForumDateTime)
	fmt.Println("\n")
	fmt.Println("=== Potion Thread End === \n")
}
