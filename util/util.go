package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
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
