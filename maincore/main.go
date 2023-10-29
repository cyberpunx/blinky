package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"localdev/HrHelper/hrHtml"
	"localdev/HrHelper/util"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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
	fmt.Println("\n === HR HELPER ===")
	config := getConfig()
	user := *config.Username
	pass := *config.Password
	client := loginAndGetCookies(user, pass)
	hrTool := NewTool(config, client)
	forumDateTime, error := util.GetTimeFromTimeZone("America/Mexico_City")
	util.Panic(error)
	fmt.Println("Forum Datetime: " + Purple + forumDateTime.Format("01/02/2006 15:04") + Reset + "\n")
	//threadHtml := hrTool.getThread("t83491-happy-birthday")
	threadHtml := hrTool.getThread("t83679p100-happy-halloween")
	thread := hrTool.parseThread(threadHtml)
	fmt.Println("Thread Title: " + Purple + thread.Title + Reset)
}

func loginAndGetCookies(user, pass string) *http.Client {
	fmt.Println("Logging in with User: " + Purple + user + " " + Reset)
	params := url.Values{}
	params.Add("username", user)
	params.Add("password", pass)
	params.Add("autologin", `on`)
	params.Add("redirect", ``)
	params.Add("query", ``)
	params.Add("login", `Conectarse`)
	bodyRequest := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://www.hogwartsrol.com/login", bodyRequest)
	util.Panic(err)
	req.Header.Set("Authority", "www.hogwartsrol.com")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "es")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "_gid=GA1.2.203551841.1698500337; toolbar_state=fa_show; _pbjs_userid_consent_data=6683316680106290; _pubcid=19fb297d-d17d-4729-8dba-f5dfc67ec7bb; trc_cookie_storage=taboola%2520global%253Auser-id%3Ddc9b0bdc-e3d7-4fb3-962b-c8b849fd38b6-tuctc369474; cto_bidid=2LV_hl9JMTY2SERCbUlDUFRSWHd4QnFNYnU2eFFRaEZ1bzMzcVJwbW9Nb1hDOTNURFFjdThycThLMXYzbVBDa0N4YmRza0p0cTNMVm81a2J6eWp1em5EWSUyRlBnJTNEJTNE; cto_bundle=X7xxSl9tM0VvSUFRV1d4alJqOW5NNDNCSmtaJTJGS1d5WW1jbUJwTVVOSDlOcTI5Nk1wNkI4aWJSRnR2NGpueWJaRWNFUnJua0ZwYkElMkJmdUx1bkwybmJ2Ynl4OFJTaXlmbjZMZWxsTkRScGVCTzBkZzJMT2ZJS3NiVXdyNTk0aGRSN1JVbnI; _fa-screen=%7B%22w%22%3A1681%2C%22h%22%3A1058%7D; _gat_gtag_UA_144386270_1=1; _ga_TTF1KWE3G4=GS1.1.1698500337.1.1.1698500422.59.0.0; _ga=GA1.1.1824435064.1698500337")
	req.Header.Set("Origin", "https://www.hogwartsrol.com")
	req.Header.Set("Referer", "https://www.hogwartsrol.com/login?")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"118\", \"Google Chrome\";v=\"118\", \"Not=A?Brand\";v=\"99\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36")

	jar, err := cookiejar.New(nil)
	if err != nil {
		// error handling
	}

	client := &http.Client{
		Jar: jar,
	}
	resp, err := client.Do(req)
	util.Panic(err)

	defer resp.Body.Close()
	printResponseStatus(resp.Status)
	return client
}

func (o *Tool) getPotionsSub() string {
	fmt.Println("Getting Sub: " + Purple + "Pociones" + Reset)

	req, err := http.NewRequest("GET", "https://www.hogwartsrol.com/"+*o.Config.PotionsClubUrl, nil)
	util.Panic(err)

	resp, err := o.Client.Do(req)
	util.Panic(err)
	defer resp.Body.Close()
	printResponseStatus(resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	util.Panic(err)

	return string(body)
}

func (o *Tool) parsePotionsSub(subHtml string) {
	potionsSubforumHtml := o.getPotionsSub()
	potionThreads := hrHtml.GetSubforumThreads(potionsSubforumHtml)

	for _, thread := range potionThreads {
		postDate, postTime := hrHtml.SubGetPostDateAndTime(thread)
		postUser := hrHtml.SubGetPostUser(thread)
		postTitle := hrHtml.SubGetPostTitle(thread)
		fmt.Println("Title: " + Purple + postTitle + Reset)
		fmt.Println("Date: " + Purple + postDate + Reset)
		fmt.Println("Time: " + Purple + postTime + Reset)
		fmt.Println("User: " + Purple + postUser + Reset)
		fmt.Println("")
	}
}

func (o *Tool) getForumHome() string {
	fmt.Println("Getting Home (Get Forum Datetime): ")

	req, err := http.NewRequest("GET", "https://www.hogwartsrol.com/", nil)
	util.Panic(err)

	resp, err := o.Client.Do(req)
	util.Panic(err)
	defer resp.Body.Close()
	printResponseStatus(resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	util.Panic(err)

	return string(body)
}

func (o *Tool) getThread(threadUrl string) string {
	fmt.Println("Getting Thread: " + Purple + threadUrl + Reset)

	req, err := http.NewRequest("GET", "https://www.hogwartsrol.com/"+threadUrl, nil)
	util.Panic(err)

	resp, err := o.Client.Do(req)
	util.Panic(err)
	defer resp.Body.Close()
	printResponseStatus(resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	util.Panic(err)

	return string(body)
}

func (o *Tool) parseThread(threadHtml string) *Thread {
	threadTitle, threadUrl, err := hrHtml.ThreadExtractTitleAndURL(threadHtml)
	util.Panic(err)

	var posts []*Post
	var pagesUrl []string
	currentPage := 1
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
		currentPage++
	}

	return &Thread{
		Title: threadTitle,
		Url:   threadUrl,
		Posts: posts,
	}
}

func (o *Tool) parsePost(postHtml string) *Post {
	postUser := hrHtml.PostGetUserName(postHtml)
	postUserUrl := hrHtml.PostGetUserUrl(postHtml)
	postDateTime := hrHtml.PostGetDateAndTime(postHtml)
	postEditedDateTime := hrHtml.PostGetEditedDateAndTime(postHtml)
	postUrl := hrHtml.PostGetUrl(postHtml)
	postContent := hrHtml.PostGetContent(postHtml)

	return &Post{
		Url:        postUrl,
		Author:     &User{Username: postUser, Url: postUserUrl, House: ""},
		DatePosted: postDateTime,
		DateEdited: postEditedDateTime,
		Content:    postContent,
	}
}
