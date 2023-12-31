package endpoint

import (
	"localdev/HrHelper/internal/hogwartsforum/dynamics/chronology"
	"localdev/HrHelper/internal/hogwartsforum/dynamics/potion"
	"localdev/HrHelper/internal/hogwartsforum/tool"
)

type Endpoints struct {
	Tool *tool.Tool
}

type SubforumPotionsClubRequest struct {
	SubForumUrls *[]string `json:"subforumUrls"`
	TimeLimit    *int      `json:"timeLimit"`
	TurnLimit    *int      `json:"turnLimit"`
}

type SubforumPotionsClubResponse struct {
	ThreadReports []potion.PotionClubReport `json:"threadReports"`
}

type ThreadsPotionsClubRequest struct {
	ThreadUrls *[]string `json:"threadUrls"`
	TimeLimit  *int      `json:"timeLimit"`
	TurnLimit  *int      `json:"turnLimit"`
}

type ThreadsPotionsClubResponse struct {
	ThreadReports []potion.PotionClubReport `json:"threadReports"`
}

type MainThreadChronologyRequest struct {
	MainThreadChronologyUrls *[]string `json:"mainThreadChronologyUrls"`
}

type MainThreadChronologyResponse struct {
	ChronologyReport []chronology.ChronoReport `json:"chronologyReport"`
}

type ThreadChronologyRequest struct {
	ThreadChronologyUrls *[]string `json:"threadChronologyUrls"`
}

type ThreadChronologyResponse struct {
	ChronologyReport []chronology.ChronoReport `json:"chronologyReport"`
}
