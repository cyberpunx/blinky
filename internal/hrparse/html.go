package hrparse

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	htmlpkg "html"
	"localdev/HrHelper/internal/util"
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

func SubGetThreadUrl(htmlString string) string {
	reader := strings.NewReader(htmlString)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "<Post Title Not Found>"
	}

	// Find the <h2> element with class "topic-title" and get its text
	threadURL := doc.Find("h2.topic-title a").AttrOr("href", "")

	return threadURL
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

func PostGetUserHouse(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "<User URL Not Found>"
	}

	information := doc.Find("div.informacion img").Last()
	house := information.AttrOr("alt", "")
	return house
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

func parseDateTime(datetimeStr string) (string, string) {
	// Split the datetime string using '-' or ','
	parts := strings.FieldsFunc(datetimeStr, func(r rune) bool { return r == '-' || r == ',' })

	// Initialize date and time strings
	var dateStr, timeStr string

	if len(parts) == 2 {
		dateStr = parts[0]
		timeStr = strings.TrimSpace(parts[1])
	} else if len(parts) == 3 {
		dateStr = parts[0]
		timeStr = parts[1] + " " + strings.TrimSpace(parts[2])
	}

	return dateStr, timeStr
}

func PostGetDateAndTime(html string, forumDateTime time.Time) *time.Time {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil
	}

	var timeStr, dateStr string

	// Find the <div> element with class "linkfecha" format:  29/11/2023, 08:31
	dateDiv := doc.Find("div.linkfecha").Nodes[0].LastChild
	datetimeStr := strings.TrimSpace(dateDiv.Data)
	if strings.Contains(datetimeStr, "Hoy a las") || strings.Contains(datetimeStr, "Ayer a las") {
		//extract time from "Hoy a las !2:01" or "Ayer a las 12:01"
		timeStr = strings.Split(datetimeStr, " ")[3]
		dateStr = strings.Split(datetimeStr, " ")[0]
		dateStr = util.AdjustDateTimeToStr(forumDateTime, dateStr)
	} else {
		dateStr = strings.Split(datetimeStr, ",")[0]
		timeStr = strings.TrimSpace(strings.Split(datetimeStr, ",")[1])
	}
	layout := "2/1/2006 15:04"
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
	encodedString, _ := dateDiv.Html()

	dateDivStr := htmlpkg.UnescapeString(encodedString)

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
				layout := "2/1/2006 15:04"
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

func PostGetDices(html string) []string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println("Error:", err)
	}
	pattern1 := `^Número aleatorio \(\d+,\d+\) : \d+$`
	pattern2 := `^Número aleatorio \(\d+,\d+\) : \(\+\d+\) : \d+$`

	var diceRolls []string
	// Find all <dl> elements with class "codebox"
	doc.Find("dl.codebox").Each(func(i int, dlSelection *goquery.Selection) {
		// Find <dd> tags inside the <dl> element
		ddSelection := dlSelection.Find("dd")
		ddSelection.Each(func(j int, ddSelection *goquery.Selection) {
			// Print the text content of <dd>
			diceLine := ddSelection.Text()
			//match both patterns
			match1, _ := regexp.MatchString(pattern1, diceLine)
			match2, _ := regexp.MatchString(pattern2, diceLine)
			if match1 || match2 {
				diceRolls = append(diceRolls, diceLine)
			}

		})
	})
	return diceRolls
}

func PostGetLinks(contentHtml string) []string {
	reader := strings.NewReader(contentHtml)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var allLinks []string
	// Find all <a> elements inside the <div class="content">
	doc.Find("a").Each(func(i int, aSelection *goquery.Selection) {
		// Get the href attribute
		link, _ := aSelection.Attr("href")
		allLinks = append(allLinks, link)
	})

	var filteredLinks []string
	for _, link := range allLinks {
		parts := strings.Split(link, "/")
		if len(parts) > 3 {
			firstSegment := parts[3]
			if strings.HasPrefix(firstSegment, "t") { //only threads links
				filteredLinks = append(filteredLinks, link)
			}
		}
	}

	return filteredLinks
}

func IsThreadVisible(html string) bool {
	searchString := `<h1 class="page-title">Informaciones</h1><p>Lo sentimos pero solamente los <strong>usuarios que tengan permisos</strong> pueden leer temas en este foro</p>`

	if strings.Contains(html, searchString) {
		return false
	} else {
		return true
	}
}
