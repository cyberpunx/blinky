package tool

import (
	"fmt"
	"io/ioutil"
	"localdev/HrHelper/internal/config"
	"localdev/HrHelper/internal/hogwartsforum/parser"
	"localdev/HrHelper/internal/util"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

func LoginAndGetCookies(user, pass string) (*http.Client, *LoginResponse) {
	fmt.Println("Logging in with User: " + config.Purple + user + " " + config.Reset)
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
	util.PrintResponseStatus(resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// error handling
	}

	var loginResponse LoginResponse
	isLoginCorrect, msg := parser.IsLoginCorrect(string(body))
	if !isLoginCorrect {
		fmt.Println(config.Red + "ERROR: " + config.CrossEmoji + config.Reset + "  Usuario y/o Contraseña Incorrectos ")
		loginResponse = LoginResponse{
			Success:  util.PBool(false),
			Messaage: &msg,
			Username: nil,
			Initials: nil,
		}
		return nil, &loginResponse
	}
	username := parser.GetUsername(string(body))
	fmt.Println("Bienvenido: " + config.Green + username + config.Reset)
	initials := util.GetInitials(username)
	loginResponse = LoginResponse{
		Success:  util.PBool(true),
		Messaage: &msg,
		Username: &username,
		Initials: &initials,
	}

	return client, &loginResponse
}

func (o *Tool) getSubforum(subUrl string) string {
	fmt.Println("Getting Sub: " + config.Purple + subUrl + config.Reset)

	req, err := http.NewRequest("GET", "https://www.hogwartsrol.com/"+subUrl, nil)
	util.Panic(err)

	resp, err := o.Client.Do(req)
	util.Panic(err)
	defer resp.Body.Close()
	util.PrintResponseStatus(resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	util.Panic(err)

	return string(body)
}

func (o *Tool) getForumHome() string {
	fmt.Println("Getting Home (Get Forum Datetime): ")

	req, err := http.NewRequest("GET", "https://www.hogwartsrol.com/", nil)
	util.Panic(err)

	resp, err := o.Client.Do(req)
	util.Panic(err)
	defer resp.Body.Close()
	util.PrintResponseStatus(resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	util.Panic(err)

	return string(body)
}

func (o *Tool) getThread(threadUrl string) string {
	fmt.Println("Getting Thread: " + config.Purple + threadUrl + config.Reset)

	baseDomain := *o.Config.BaseUrl

	_, err := url.ParseRequestURI(threadUrl)
	if err != nil || !strings.HasPrefix(threadUrl, baseDomain) {
		threadUrl = baseDomain + threadUrl
	}

	req, err := http.NewRequest("GET", threadUrl, nil)
	util.Panic(err)

	resp, err := o.Client.Do(req)
	util.Panic(err)
	defer resp.Body.Close()
	util.PrintResponseStatus(resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	util.Panic(err)

	return string(body)
}

func (o *Tool) GetPostSecrets() (string, string) {
	fmt.Println("Getting Post Secrets: ")
	baseDomain := *o.Config.BaseUrl
	postUrl := baseDomain + "/post?f=44&mode=newtopic"

	req, err := http.NewRequest("GET", postUrl, nil)
	util.Panic(err)

	resp, err := o.Client.Do(req)
	util.Panic(err)
	defer resp.Body.Close()
	util.PrintResponseStatus(resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	util.Panic(err)

	secret1, secret2 := parser.GetPostSecrets(string(body))
	if secret1 == "" || secret2 == "" {
		fmt.Println(config.Red + "ERROR: " + config.CrossEmoji + config.Reset + "  Could not get post secrets")
		os.Exit(1)
	}

	return secret1, secret2
}
