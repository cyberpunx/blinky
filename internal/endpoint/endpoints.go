package endpoint

import (
	"encoding/json"
	"fmt"
	"localdev/HrHelper/internal/potion"
	"localdev/HrHelper/internal/tool"
	"net/http"
)

const (
	ServerPort = ":8080"
)

func (o *Endpoints) SubforumPotionsClub(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request SubforumPotionsClubRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("\n\n ========= SUBFORUM CLUB DE POCIONES =========\n\n")
	if len(*request.SubForumUrls) == 0 {
		fmt.Println("No subforums URLs to process")
	}
	var reportMainList []potion.PotionClubReport
	for _, url := range *request.SubForumUrls {
		fmt.Println("=== Fetching Subforum === \n")
		potionSubHtml := o.Tool.GetSubforum(url)
		subforumThreads := o.Tool.ParseSubforum(potionSubHtml)
		fmt.Println("=== Fetch Ended === \n")
		reportList := o.Tool.ProcessPotionsSubforum(subforumThreads, *request.TurnLimit, *request.TimeLimit)
		reportMainList = append(reportMainList, reportList...)
	}

	response := SubforumPotionsClubResponse{ThreadReports: reportMainList}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (o *Endpoints) ThreadsPotionsClub(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request ThreadsPotionsClubRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("\n\n ========= SUBFORUM CLUB DE POCIONES =========\n\n")
	if len(*request.ThreadsUrls) == 0 {
		fmt.Println("No Threads URLs to process")
	}
	var reportMainList []potion.PotionClubReport
	for _, url := range *request.ThreadsUrls {
		potionThreadHtml := o.Tool.GetThread(url)
		potionThread := o.Tool.ParseThread(potionThreadHtml)
		report := o.Tool.ProcessPotionsThread(*potionThread, *request.TurnLimit, *request.TimeLimit)
		reportMainList = append(reportMainList, report)
	}

	response := ThreadsPotionsClubResponse{ThreadReports: reportMainList}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (o *Endpoints) ConfigureAndServeEndpoints() {
	http.HandleFunc("/subforumPotionsClub", o.SubforumPotionsClub)
	http.HandleFunc("/threadsPotionsClub", o.ThreadsPotionsClub)

	go func() {
		if err := http.ListenAndServe(ServerPort, nil); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Server HTTP listening on port " + ServerPort + "")
}

func NewEndpoints(tool *tool.Tool) *Endpoints {
	return &Endpoints{
		Tool: tool,
	}
}
