package main

import (
	"localdev/HrHelper/hrHtml"
	"localdev/HrHelper/util"
	"strconv"
	"strings"
)

func ParseDiceRoll(dicerolls []string) []*Dice {
	var dices []*Dice

	for _, diceroll := range dicerolls {
		resultStr := strings.TrimSpace(strings.Split(diceroll, ":")[len(strings.Split(diceroll, ":"))-1])
		//convert result from string to int
		result, err := strconv.Atoi(resultStr)
		util.Panic(err)
		dice := &Dice{
			DiceLine: diceroll,
			Result:   result,
		}
		dices = append(dices, dice)
	}
	return dices
}

func (o *Tool) parseSubforum(subHtml string) []*Thread {
	threadList := hrHtml.GetSubforumThreads(subHtml)

	var threads []*Thread
	for _, thread := range threadList {
		threadUrl := hrHtml.SubGetThreadUrl(thread)
		threadHtml := o.getThread(threadUrl)
		thread := o.parseThread(threadHtml)
		threads = append(threads, thread)
	}
	return threads
}

func (o *Tool) parseThread(threadHtml string) *Thread {
	threadTitle, threadUrl, err := hrHtml.ThreadExtractTitleAndURL(threadHtml)
	util.Panic(err)

	var posts []*Post
	var pagesUrl []string
	pagesUrl = append(pagesUrl, threadUrl)
	for {
		// Extract the post list from the current page
		postList := hrHtml.ThreadListPosts(threadHtml)
		for _, post := range postList {
			post := o.parsePost(post)
			posts = append(posts, post)
		}

		// Check if there is a "next" link in the pagination
		nextPageURL, hasMore := hrHtml.ThreadNextPageURL(threadHtml)

		if !hasMore {
			break // No more pages to fetch
		}

		// Fetch the next page and update the threadHtml
		pagesUrl = append(pagesUrl, nextPageURL)
		nextPageHTML := o.getThread(nextPageURL)
		threadHtml = nextPageHTML
	}

	firstPostId := posts[0].Id
	var filteredPosts []*Post
	filteredPosts = append(filteredPosts, posts[0])
	for _, post := range posts {
		if post.Id != firstPostId {
			filteredPosts = append(filteredPosts, post)
		}
	}

	return &Thread{
		Title:    threadTitle,
		Url:      threadUrl,
		Author:   posts[0].Author,
		Created:  posts[0].Created,
		LastPost: posts[len(posts)-1],
		Pages:    pagesUrl,
		Posts:    filteredPosts,
	}
}

func (o *Tool) parsePost(postHtml string) *Post {
	postUser := hrHtml.PostGetUserName(postHtml)
	postUserUrl := hrHtml.PostGetUserUrl(postHtml)
	postUserHouse := hrHtml.PostGetUserHouse(postHtml)
	postDateTime := hrHtml.PostGetDateAndTime(postHtml)
	postEditedDateTime := hrHtml.PostGetEditedDateAndTime(postHtml)
	postUrl := hrHtml.PostGetUrl(postHtml)
	postContent := hrHtml.PostGetContent(postHtml)
	dices := ParseDiceRoll(hrHtml.PostGetDices(postHtml))

	return &Post{
		Url:     postUrl,
		Author:  &User{Username: postUser, Url: postUserUrl, House: postUserHouse},
		Created: postDateTime,
		Edited:  postEditedDateTime,
		Content: postContent,
		Dices:   dices,
		Id:      postUrl[strings.LastIndex(postUrl, "#")+1:],
	}
}
