package endpoint

import (
	"encoding/json"
	"fmt"
	"localdev/HrHelper/internal/hogwartsforum/tool"
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

	threadReports := o.Tool.ProcessPotionsSubforumList(request.SubForumUrls, request.TimeLimit, request.TurnLimit)
	response := SubforumPotionsClubResponse{ThreadReports: threadReports}

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

	threadReports := o.Tool.ProcessPotionsThreadList(request.ThreadUrls, request.TimeLimit, request.TurnLimit)
	response := SubforumPotionsClubResponse{ThreadReports: threadReports}

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
