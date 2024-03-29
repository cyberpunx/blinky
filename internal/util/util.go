package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"localdev/HrHelper/internal/config"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func GetInitials(name string) string {
	words := strings.Fields(name)
	var initials string

	if len(words) == 1 {
		initials = string(words[0][0]) + string(words[0][1])
	} else {
		initials = string(words[0][0]) + string(words[1][0])
	}

	return strings.ToUpper(initials)
}

func PrintResponseStatus(status string) {
	statusColor := ""
	statusEmoji := ""
	if status == "200 OK" {
		statusColor = config.Green
		statusEmoji = " " + config.CheckEmoji + " "
	} else {
		statusColor = config.Red
		statusEmoji = " " + config.CrossEmoji + " "
	}
	LongPrintlnPrintln("Response Status: " + statusColor + statusEmoji + " " + status + config.Reset)
}

func LongPrintlnPrintln(a ...any) {
	// Convert all arguments into a string slice
	stringArgs := make([]string, len(a))
	for i, arg := range a {
		stringArgs[i] = fmt.Sprint(arg)
	}

	// Join all arguments into a single string
	fullString := strings.Join(stringArgs, " ")

	// Print to stdout using util.LongPrintlnPrintln
	_, err := fmt.Println(fullString)
	Panic(err)

	// Append to a log file
	file, err := os.OpenFile("log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Panic(err)
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)
	logger.Println(RemoveUnicodeFromStr(fullString))
}

func RemoveUnicodeFromStr(str string) string {
	str = RemoveColorFromStr(str)
	str = strings.Replace(str, config.CheckEmoji, "[OK]", -1)
	str = strings.Replace(str, config.CrossEmoji, "[X]", -1)
	str = strings.Replace(str, config.RightArrowEmoji, "-->", -1)
	return str
}

func RemoveColorFromStr(str string) string {
	str = strings.Replace(str, config.Reset, "", -1)
	str = strings.Replace(str, config.Red, "", -1)
	str = strings.Replace(str, config.Green, "", -1)
	str = strings.Replace(str, config.Yellow, "", -1)
	str = strings.Replace(str, config.Blue, "", -1)
	str = strings.Replace(str, config.Purple, "", -1)
	str = strings.Replace(str, config.Cyan, "", -1)
	str = strings.Replace(str, config.Gray, "", -1)
	str = strings.Replace(str, config.White, "", -1)
	return str
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func PStr(s string) *string {
	return &s
}

func PBool(b bool) *bool {
	return &b
}

type P map[string]interface{}

func Fprint(format string, p P) string {
	args, i := make([]string, len(p)*2), 0
	for k, v := range p {
		args[i] = "{" + k + "}"
		args[i+1] = fmt.Sprint(v)
		i += 2
	}
	return strings.NewReplacer(args...).Replace(format)
}

func UtoA(unicodeStr string) string {
	unicodeStr = strings.Replace(unicodeStr, "\033[0m", "", -1)
	unicodeStr = strings.Replace(unicodeStr, "\033[31m", "", -1)
	unicodeStr = strings.Replace(unicodeStr, "\033[32m", "", -1)
	unicodeStr = strings.Replace(unicodeStr, "\033[33m", "", -1)
	unicodeStr = strings.Replace(unicodeStr, "\033[34m", "", -1)
	unicodeStr = strings.Replace(unicodeStr, "\033[35m", "", -1)
	unicodeStr = strings.Replace(unicodeStr, "\033[36m", "", -1)
	unicodeStr = strings.Replace(unicodeStr, "\033[37m", "", -1)
	unicodeStr = strings.Replace(unicodeStr, "\033[97m", "", -1)
	unicodeStr = strings.Replace(unicodeStr, "✔", "[OK]", -1)
	unicodeStr = strings.Replace(unicodeStr, "❌", "[X]", -1)
	unicodeStr = strings.Replace(unicodeStr, "▶", "-->", -1)
	return unicodeStr
}

func SplitDateAndTime(dateTime string) (string, string) {
	parts := strings.Split(dateTime, ", ")
	if len(parts) != 2 {
		return "", ""
	}
	date := parts[0]
	time := parts[1]
	return date, time
}

func GetTimeFromTimeZone(timezone string) (time.Time, error) {
	url := fmt.Sprintf("http://worldtimeapi.org/api/timezone/%s", timezone)

	resp, err := http.Get(url)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	var response struct {
		Datetime string `json:"datetime"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return time.Time{}, err
	}

	inputLayout := "2006-01-02T15:04:05.999999-07:00"
	outputLayout := "01/02/2006 15:04"
	dateTimeParcial, _ := time.Parse(inputLayout, response.Datetime)
	formattedDateTime := dateTimeParcial.Format(outputLayout)

	dateTime, _ := time.Parse(outputLayout, formattedDateTime)
	if err != nil {
		return time.Time{}, err
	}

	return dateTime, nil
}

func AdjustDateTime(currentDateTime time.Time, dateString string) time.Time {
	if dateString == "Hoy" {
		return currentDateTime
	} else if dateString == "Ayer" {
		return currentDateTime.AddDate(0, 0, -1)
	}
	return currentDateTime
}

func AdjustDateTimeToStr(currentDate time.Time, dateString string) string {
	if dateString == "Hoy" {
		return currentDate.Format("02/01/2006")
	} else if dateString == "Ayer" {
		return currentDate.AddDate(0, 0, -1).Format("02/01/2006")
	}
	return dateString
}

func IsDateInBetween(date time.Time, startDate time.Time, endDate time.Time) bool {
	return date.After(startDate) && date.Before(endDate)
}

func IsDateInCurrentMonth(date time.Time) bool {
	currentDate := time.Now()
	return date.Month() == currentDate.Month() && date.Year() == currentDate.Year()
}

func MarshalJsonPretty(i interface{}) []byte {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "    ")
	Panic(encoder.Encode(i))
	return buffer.Bytes()
}

func TrimAndToLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func IsDateWithinTimeLimit(currentTime, lastPostTime time.Time, timeThreshold time.Duration) bool {
	// Check if the current exceeds the time threshold
	if lastPostTime.Add(timeThreshold).Before(currentTime) {
		return false
	} else {
		return true
	}
}

func Sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
