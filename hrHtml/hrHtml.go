package hrHtml

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"localdev/HrHelper/util"
	"log"
	"regexp"
	"strings"
	"time"
)

func GetSubforumAllThreads(html string) []string {
	var threads []string

	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("li.row").Each(func(index int, element *goquery.Selection) {
		text, _ := element.Html()
		threads = append(threads, text)
	})

	return threads
}

func GetSubforumPinnedThreads(html string) []string {
	var pinnedThreads []string

	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	// Find <li> tags inside <div class="forumbg announcement">
	doc.Find("div.forumbg.announcement li.row").Each(func(index int, element *goquery.Selection) {
		text, _ := element.Html()
		pinnedThreads = append(pinnedThreads, text)
	})

	return pinnedThreads
}

func GetSubforumThreads(html string) []string {
	var threads []string

	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	// Find <li> tags inside <div class="forumbg"> but not within <div class="forumbg announcement">
	doc.Find("div.forumbg:not(.announcement) li.row").Each(func(index int, element *goquery.Selection) {
		text, _ := element.Html()
		threads = append(threads, text)
	})

	return threads
}

func SubGetPostUser(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "<User Not Found>"
	}

	// Find the <dd> element with class "lastpost" inside the <dl> element
	ddLastPost := doc.Find("dl.icon dd.lastpost")

	// Extract the date
	date := ddLastPost.Find("a").Text()

	return date
}

func SubGetPostDateAndTime(html string) (string, string) {
	re := regexp.MustCompile(`(?i)((Ayer|Hoy) a las (\d{2}:\d{2})|(\d{2}/\d{2}/\d{4}, \d{2}:\d{2}))`)
	matches := re.FindStringSubmatch(html)

	if len(matches) < 5 {
		return "<Date not found>", "<Time not found>"
	}

	date, time := "", ""

	if strings.Contains(matches[0], "Hoy a las") || strings.Contains(matches[0], "Ayer a las") {
		date = matches[2]
		time = matches[3]
	} else {
		date, time = util.SplitDateAndTime(matches[0])
	}

	return date, time
}

func SubGetPostTitle(htmlString string) string {
	reader := strings.NewReader(htmlString)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "<Post Title Not Found>"
	}

	// Find the <h2> element with class "topic-title" and get its text
	title := doc.Find("h2.topic-title").Text()

	return title
}

func ThreadExtractTitleAndURL(htmlFragment string) (title, url string, err error) {
	reader := strings.NewReader(htmlFragment)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", "", err
	}

	// Find the <a> element inside the <h1> element
	link := doc.Find("h1.page-title a")

	// Extract the title and URL
	title = link.Text()
	url, _ = link.Attr("href")

	return title, url, nil
}

func ThreadListPosts(html string) []string {
	var posts []string

	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	doc.Find("div.post").Each(func(index int, element *goquery.Selection) {
		text, _ := element.Html()
		posts = append(posts, text)
	})

	return posts
}

func ThreadNextPageURL(html string) (string, bool) {
	// Load the HTML content into a goquery document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", false
	}

	// Find the <p> tag with the "pagination" class
	paginationElement := doc.Find("p.pagination span")
	if paginationElement.Length() == 0 {
		return "", false // No "next" link found
	}

	nextButton := paginationElement.Find("a").Last()
	nextButtonHtml, _ := nextButton.Html()
	if !strings.Contains(nextButtonHtml, "Siguiente") {
		return "", false // No "next" link found
	}
	nextButtonUrl := nextButton.AttrOr("href", "")
	return nextButtonUrl, true
}

func PostGetUserName(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "<Username Not Found>"
	}

	post := doc.Find("div.post1")
	username := post.Find("a[href^='/u']").Text()
	return username
}

func PostGetUserUrl(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "<User URL Not Found>"
	}

	post := doc.Find("div.post1")
	userUrl, _ := post.Find("a[href^='/u']").Attr("href")
	return userUrl
}

func PostGetUrl(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "<Post URL Not Found>"
	}
	var url string
	// Find the <div> element with class "linkfecha"
	linkfecha := doc.Find("div.linkfecha")
	link := linkfecha.Find("a")
	url, _ = link.Attr("href")

	return url
}

func PostGetDateAndTime(html string) *time.Time {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil
	}

	var timeStr, dateStr string

	// Find the <div> element with class "linkfecha"
	dateDiv := doc.Find("div.linkfecha").Nodes[0].LastChild
	datetimeStr := strings.TrimSpace(dateDiv.Data)
	if strings.Contains(datetimeStr, "Hoy a las") || strings.Contains(datetimeStr, "Ayer a las") {
		//extract time from "Hoy a las !2:01" or "Ayer a las 12:01"
		timeStr = strings.Split(datetimeStr, " ")[3]
		dateStr = strings.Split(datetimeStr, " ")[0]
		location, _ := time.LoadLocation("America/New_York")
		currentTime := time.Now().In(location)
		dateStr = util.AdjustDateTimeToStr(currentTime, dateStr)
	} else {
		dateStr = strings.Split(datetimeStr, ",")[0]
		timeStr = strings.TrimSpace(strings.Split(datetimeStr, ",")[1])
	}
	layout := "2/01/2006 15:04"
	dateTime, err := time.Parse(layout, dateStr+" "+timeStr)
	util.Panic(err)
	return &dateTime
}

func PostGetEditedDateAndTime(html string) *time.Time {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil
	}

	// Find the <div> element with class "linkfecha"
	dateDiv := doc.Find("div.post3")
	dateDivStr, _ := dateDiv.Html()

	if strings.Contains(dateDivStr, "Última edición por") {
		numberIndex := -1
		for i, char := range dateDivStr {
			if char >= '0' && char <= '9' {
				numberIndex = i
				break
			}
		}

		// If a number was found, look for the word ", editado"
		if numberIndex >= 0 {
			editedIndex := strings.Index(dateDivStr, ", editado")
			if editedIndex > numberIndex {
				// Extract the desired part of the string
				result := dateDivStr[numberIndex:editedIndex]
				dateStr := strings.Split(result, ",")[0]
				timeStr := strings.TrimSpace(strings.Split(result, ",")[1])
				layout := "2/01/2006 15:04"
				dateTime, err := time.Parse(layout, dateStr+" "+timeStr)
				util.Panic(err)
				return &dateTime
			}
		}
	}
	return nil
}

func PostGetContent(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "<Post Content Not Found>"
	}
	var content string
	// Find the <div> element with class "content"
	content, _ = doc.Find("div.content").Html()

	return content
}
