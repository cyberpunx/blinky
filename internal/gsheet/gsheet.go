package gsheet

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"localdev/HrHelper/internal/util"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

const (
	SheetRangeDaysOff     = "Permisos Pociones!A269:B"
	SheetRangePlayerBonus = "Logros Pociones!A2:B"
)

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	// Create a random state token
	state := "st" + string(rand.New(rand.NewSource(time.Now().UnixNano())).Int63())

	// Create a channel to receive the authorization code
	codeChan := make(chan string)

	// Start a web server to listen on the callback URL
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check the state token
		if r.URL.Query().Get("state") != state {
			http.Error(w, "State token does not match", http.StatusBadRequest)
			return
		}
		// Send the code to the channel
		codeChan <- r.URL.Query().Get("code")
		fmt.Fprintf(w, "Authorization complete, you can close this window.")
		server.Shutdown(context.Background())
	})

	// Open the authorization URL in the user's browser
	authURL := config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	fmt.Printf("Opening browser for authorization: %s\n", authURL)
	openbrowser(authURL)

	// Start the server and wait for the auth code
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			util.Panic(err)
		}
	}()
	code := <-codeChan

	// Exchange the code for a token
	tok, err := config.Exchange(context.Background(), code)
	util.Panic(err)
	return tok
}

func openbrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	util.Panic(err)
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	util.Panic(err)
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func GetClient(tokFile string, ctx context.Context, config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first time.
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(ctx, tok)
}

func ReadCredentials(credPath string) ([]byte, error) {
	b, err := ioutil.ReadFile(credPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file: %w", err)
	}
	return b, nil
}

func ReadSheetData(srv *sheets.Service, spreadsheetId, readRange string) ([][]interface{}, error) {
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve data from sheet: %w", err)
	}
	return resp.Values, nil
}

func DisplayData(data [][]interface{}) {
	if len(data) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Data:")
		for _, row := range data {
			fmt.Printf("%s\n", row)
		}
	}
}

func GetSheetService(tokFile, credPath string) *sheets.Service {
	ctx := context.Background()
	// Read Credentials
	credentials, err := ReadCredentials(credPath)
	util.Panic(err)

	// Configure OAuth2 Client
	gconfig, err := google.ConfigFromJSON(credentials, sheets.SpreadsheetsReadonlyScope)
	util.Panic(err)
	client := GetClient(tokFile, ctx, gconfig)

	// Create Sheets Service
	service, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	util.Panic(err)
	return service
}

func ParseDayOff(rows [][]interface{}) []DayOff {
	var daysOff []DayOff
	for _, row := range rows {
		//skip first row because it's the header
		//if i == 0 {
		//	continue
		//}
		if len(row) == 2 {
			date, err := time.Parse("2/01/2006", row[1].(string)) //format DD/MM/YYYY
			util.Panic(err)
			dayOff := DayOff{
				Username: row[0].(string),
				Date:     date,
			}
			daysOff = append(daysOff, dayOff)
		}
	}
	return daysOff
}

func ParsePlayerBonus(rows [][]interface{}) []PlayerBonus {
	var playerBonus []PlayerBonus
	for _, row := range rows {
		//skip first row because it's the header
		//if i == 0 {
		//	continue
		//}
		if len(row) == 2 {
			bonusValue, err := strconv.Atoi(row[1].(string))
			util.Panic(err)
			pBonus := PlayerBonus{
				Username: row[0].(string),
				Bonus:    bonusValue,
			}
			playerBonus = append(playerBonus, pBonus)
		}
	}
	return playerBonus
}

func FindDayOffForUser(daysOff *[]DayOff, username string) *DayOff {
	for _, dayOff := range *daysOff {
		dayOffUsername := util.TrimAndToLower(dayOff.Username)
		username = util.TrimAndToLower(username)
		if dayOffUsername == username {
			return &dayOff
		}
	}
	return nil
}

func FindDayOffForPLayerBetweenDates(daysOff *[]DayOff, username string, startDate, endDate time.Time) *DayOff {
	for _, dayOff := range *daysOff {
		dayOffUsername := util.TrimAndToLower(dayOff.Username)
		username = util.TrimAndToLower(username)
		if dayOffUsername == username && dayOff.Date.After(startDate) && dayOff.Date.Before(endDate) {
			return &dayOff
		}
	}
	return nil
}

func GetPlayerBonusForUsername(playerBonus *[]PlayerBonus, username string) int {
	for _, pb := range *playerBonus {
		pbUser := util.TrimAndToLower(pb.Username)
		username = util.TrimAndToLower(username)
		if pbUser == username {
			return pb.Bonus
		}
	}
	return 0
}
