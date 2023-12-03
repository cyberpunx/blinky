package potion

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"localdev/HrHelper/internal/config"
	"localdev/HrHelper/internal/hogwartsforum/parser"
	"localdev/HrHelper/internal/util"
	"strconv"
	"strings"
	"time"
)

const (
	Player1   = "Player1"
	Player2   = "Player2"
	Moderator = "Moderator"
	Other     = "Moderator"

	StatusSuccess        Status = "Success"
	StatusFail           Status = "Fail"
	StatusWaitingPlayer1 Status = "WaitingPlayer1"
	StatusWaitingPlayer2 Status = "WaitingPlayer2"
)

type Status string

type Potion struct {
	Name        string
	Ingredients []string
	TargetScore int
	TurnLimit   int
}

type PotionsUser struct {
	Name        string
	House       string
	Role        string
	PlayerBonus int
	Posts       []*parser.Post
}
type PotionClubReport struct {
	Player1   PotionsUser
	Player2   PotionsUser
	Moderator PotionsUser
	Other     []PotionsUser
	Thread    parser.Thread
	Potion    Potion
	Status    Status
	Score     PotionClubScoreBoard
	TurnLimit int
	TimeLimit int
	Turns     []PotionClubTurn
}

type PotionClubTurn struct {
	Player      PotionsUser
	Number      int
	DiceValue   int
	OnTime      bool
	TimeElapsed time.Time
}

type PotionClubScoreBoard struct {
	DiceScoreSum   int
	ModeratorBonus int
	ModeratorMalus int
	Player1Bonus   int
	Player2Bonus   int
	TargetScore    int
	TotalScore     int
	Success        bool
}

func getPotionFromThread(thread parser.Thread) *Potion {
	potionPostHtml := thread.Posts[0].Content

	reader := strings.NewReader(potionPostHtml)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil
	}

	var name string
	var turnLimit string
	var targetScore string
	var ingredients []string
	var targetScoreInt int
	var turnLimitInt int

	potionInfo := doc.Find("div.xxEDV").Last()
	potionInfo.Find("li").Each(func(i int, liSelection *goquery.Selection) {
		if i == 0 {
			//Potion Name
			name = liSelection.Text()
			name = strings.Split(name, ":")[1]
		} else if i == 1 {
			//Potion TurnLimit
			turnLimit = liSelection.Text()
			turnLimit := strings.Split(strings.Split(turnLimit, ":")[1], " ")[1]
			turnLimitInt, _ = strconv.Atoi(turnLimit)
		} else if i == 2 {
			//Potion TargetScore
			targetScore = liSelection.Text()
			targetScore := strings.Split(strings.Split(targetScore, ":")[1], " ")[1]
			targetScoreInt, _ = strconv.Atoi(targetScore)
		} else {
			ingredients = append(ingredients, liSelection.Text())
		}
	})

	return &Potion{
		Name:        name,
		Ingredients: ingredients,
		TargetScore: targetScoreInt,
		TurnLimit:   turnLimitInt,
	}
}
func identifyRolesOnThread(thread parser.Thread) (player1 PotionsUser, player2 PotionsUser, moderator PotionsUser, other []PotionsUser) {
	moderator.Name = thread.Author.Username
	moderator.Role = Moderator

	for _, post := range thread.Posts {
		if post.Author.Username != moderator.Name {
			if player1.Name == "" {
				player1.Name = post.Author.Username
				player1.Role = Player1
				player1.House = post.Author.House
			} else if player2.Name == "" && post.Author.Username != player1.Name {
				player2.Name = post.Author.Username
				player2.Role = Player2
				player2.House = post.Author.House
			}
		}
	}

	for _, post := range thread.Posts {
		if post.Author.Username != moderator.Name && post.Author.Username != player1.Name && post.Author.Username != player2.Name {
			other = append(other, PotionsUser{Name: post.Author.Username, Role: Other})
		}
	}

	for _, post := range thread.Posts {
		if post.Author.Username == player1.Name {
			player1.Posts = append(player1.Posts, post)
		} else if post.Author.Username == player2.Name {
			player2.Posts = append(player2.Posts, post)
		} else if post.Author.Username == moderator.Name {
			moderator.Posts = append(moderator.Posts, post)
		} else {
			for _, otherUser := range other {
				if post.Author.Username == otherUser.Name {
					otherUser.Posts = append(otherUser.Posts, post)
				}
			}
		}
	}

	return
}
func isPlayer(post parser.Post, player1, player2 PotionsUser) bool {
	return post.Author.Username == player1.Name || post.Author.Username == player2.Name
}
func findPlayerAndRole(username string, player1, player2, moderator PotionsUser, others []PotionsUser) (*PotionsUser, string) {
	if username == player1.Name {
		return &player1, player1.Role
	} else if username == player2.Name {
		return &player2, player2.Role
	} else if username == moderator.Name {
		return &moderator, moderator.Role
	} else {
		for _, otherUser := range others {
			if username == otherUser.Name {
				return &otherUser, otherUser.Role
			}
		}
	}
	return nil, ""
}
func removeOtherPostsFromThread(player1 PotionsUser, player2 PotionsUser, moderator PotionsUser, other []PotionsUser, thread parser.Thread) parser.Thread {
	var threadWithoutOthers parser.Thread
	threadWithoutOthers = thread
	threadWithoutOthers.Posts = nil

	for _, post := range thread.Posts {
		if post.Author.Username == player1.Name || post.Author.Username == player2.Name || post.Author.Username == moderator.Name {
			threadWithoutOthers.Posts = append(threadWithoutOthers.Posts, post)
		}
	}

	return threadWithoutOthers
}
func isPostWithinTimeLimit(currentPostTime, lastPostTime time.Time, timeThreshold time.Duration) bool {
	// Check if the current post exceeds the time threshold
	if lastPostTime.Add(timeThreshold).Before(currentPostTime) {
		return false
	} else {
		return true
	}

}
func printPostReport(isPlayer bool, postCount int, postUser string, role string, turnCount int, onTime bool, turnDice string, diceTotal int) string {
	strReport := ""
	timeStatus := ""
	if !isPlayer {
		strReport = "{i}) " + config.Purple + "{postUser} " + config.Reset + " ({role})" + config.Reset
	} else {
		if onTime {
			timeStatus = config.Green + "OK" + config.Reset
		} else {
			timeStatus = config.Red + "FAIL" + config.Reset
		}
		strReport = "{i}) Turn {turnCount} " + config.Purple + "{postUser} " + config.Reset + " ({role}) | Time: {timeStatus} | Dice: {turnDice} | Total: {diceTotal}" + config.Reset
	}

	s := util.Fprint(strReport,
		util.P{"i": strconv.Itoa(postCount),
			"postUser":   postUser,
			"role":       role,
			"turnCount":  strconv.Itoa(turnCount),
			"timeStatus": timeStatus,
			"turnDice":   turnDice,
			"diceTotal":  config.Purple + strconv.Itoa(diceTotal) + config.Reset,
		})
	return s
}
func PotionGetReportFromThread(rawThread parser.Thread, turnLimit int, timeLimitHours int, forumDateTime time.Time) PotionClubReport {
	timeThreshold := time.Duration(timeLimitHours) * time.Hour
	potion := getPotionFromThread(rawThread)
	player1, player2, moderator, others := identifyRolesOnThread(rawThread)
	threadWithoutOthers := removeOtherPostsFromThread(player1, player2, moderator, others, rawThread)
	playerPostCount := make(map[string]int)
	lastPostTime := *threadWithoutOthers.Created
	turnCount := 1
	postCount := 1
	postDice := ""
	diceTotal := 0
	postOnTime := false
	threadLastPost := *threadWithoutOthers.Posts[len(threadWithoutOthers.Posts)-1]

	result := PotionClubReport{
		Player1:   player1,
		Player2:   player2,
		Moderator: moderator,
		Other:     others,
		Thread:    rawThread,
		Potion:    *potion,
		Status:    StatusWaitingPlayer1,
		Score: PotionClubScoreBoard{
			DiceScoreSum:   0,
			ModeratorBonus: 0,
			ModeratorMalus: 0,
			Player1Bonus:   0,
			Player2Bonus:   0,
			TargetScore:    potion.TargetScore,
			TotalScore:     0,
			Success:        false,
		},
		TurnLimit: turnLimit,
		TimeLimit: timeLimitHours,
		Turns:     nil,
	}

	for _, post := range threadWithoutOthers.Posts {
		postUser := post.Author.Username
		postPlayer, postRole := findPlayerAndRole(postUser, player1, player2, moderator, others)
		isPlayerFlag := postPlayer.Name == player1.Name || postPlayer.Name == player2.Name

		if isPlayerFlag {
			playerPostCount[postUser]++
			if postRole == Player1 {
				result.Status = StatusWaitingPlayer2
			} else if postRole == Player2 {
				result.Status = StatusWaitingPlayer1
			}
			postOnTime = isPostWithinTimeLimit(*post.Created, lastPostTime, timeThreshold)

			postDiceValue := 0
			if len(post.Dices) != 1 {
				postDice = "N/A"
			} else {
				postDiceValue = post.Dices[0].Result
				postDice = config.Yellow + strconv.Itoa(post.Dices[0].Result) + config.Reset
				diceTotal += postDiceValue
				result.Score.DiceScoreSum += postDiceValue
			}
			turn := PotionClubTurn{
				Player:      *postPlayer,
				Number:      turnCount,
				DiceValue:   postDiceValue,
				OnTime:      postOnTime,
				TimeElapsed: lastPostTime,
			}
			result.Turns = append(result.Turns, turn)
			lastPostTime = *post.Created
		}

		s := printPostReport(isPlayerFlag, postCount, postUser, postRole, turnCount, postOnTime, postDice, diceTotal)
		fmt.Println(s)

		if threadLastPost.Id == post.Id && isPlayerFlag {
			elapsedTime := forumDateTime.Sub(*post.Created)
			if elapsedTime > timeThreshold {
				fmt.Println(config.Red+"Time Passed: "+config.Reset, elapsedTime)
			} else {
				fmt.Println(config.Green+"Time Passed: "+config.Reset, elapsedTime)
			}
		}

		if turnCount == turnLimit {
			if diceTotal > potion.TargetScore {
				result.Status = StatusSuccess
				result.Score.Success = true
			} else {
				result.Status = StatusFail
				result.Score.Success = false
			}
			result.Score = PotionClubScoreBoard{
				DiceScoreSum: diceTotal,
				TotalScore:   diceTotal + result.Score.ModeratorBonus + result.Score.ModeratorMalus + result.Score.Player1Bonus + result.Score.Player2Bonus,
			}

		}

		postCount++
		if playerPostCount[player1.Name] > 0 && playerPostCount[player2.Name] > 0 {
			playerPostCount[player1.Name] = 0
			playerPostCount[player2.Name] = 0
			turnCount++
		}
	}

	return result
}

/*
func ClubPotionsProcessorOld(thread Thread, hoursLimit int) {
	timeThreshold := time.Duration(hoursLimit) * time.Hour
	turnCount := 1
	player1, player2 := "", ""
	player1PostCount, player2PostCount := 0, 0
	moderator := thread.Author.Username
	authorRole := ""
	timeLimitStatus := ""
	postCount := 1
	diceSum := 0
	turnDice := ""

	// Initialize maps to count each player's posts and store the time of the last post by each player
	playerPostCount := make(map[string]int)
	lastPostTime := *thread.Created

	// Iterate through the posts to identify players and count turns
	for _, post := range thread.Posts {
		authorUsername := post.Author.Username

		if authorUsername == moderator {
			authorRole = Cyan + "Moderator" + Reset
		}

		// Identify the players
		if player1 == "" && authorUsername != moderator {
			player1 = authorUsername
		} else if player2 == "" && authorUsername != player1 && authorUsername != moderator {
			player2 = authorUsername
		}

		// Count the post for the current player and update the last post time
		if authorUsername != moderator {
			playerPostCount[authorUsername]++
		}

		if authorUsername == player1 {
			authorRole = "Player 1"
			player1PostCount++
		} else if authorUsername == player2 {
			authorRole = "Player 2"
			player2PostCount++
		} else {
			authorRole = Cyan + "Other" + Reset
		}

		if authorUsername != moderator {
			// Check if the current post exceeds the time threshold
			if lastPostTime.Add(timeThreshold).Before(*post.Created) {
				timeLimitStatus = Red + "FAIL" + Reset
			} else {
				timeLimitStatus = Green + "OK" + Reset
			}
			lastPostTime = *post.Created

			if len(post.Dices) != 1 {
				turnDice = "N/A"
			} else {
				turnDice = Yellow + strconv.Itoa(post.Dices[0].Result) + Reset
				diceSum += post.Dices[0].Result
			}
		}

		strReport := ""
		if authorUsername == moderator {
			strReport = "{i}) " + Purple + "{authorUsername} " + Reset + " ({role})" + Reset
		} else {
			strReport = "{i}) Turn {turnCount} " + Purple + "{authorUsername} " + Reset + " ({role}) | Time: {timeLimitStatus} | Dice: {turnDice} | Total: {diceSum}" + Reset
		}

		s := util.Fprint(strReport,
			util.P{"i": strconv.Itoa(postCount),
				"authorUsername":  authorUsername,
				"role":            authorRole,
				"turnCount":       strconv.Itoa(turnCount),
				"timeLimitStatus": timeLimitStatus,
				"turnDice":        turnDice,
				"diceSum":         Purple + strconv.Itoa(diceSum) + Reset,
			})
		fmt.Println(s)
		postCount++

		// Check if both players have posted, indicating a turn
		if playerPostCount[player1] > 0 && playerPostCount[player2] > 0 {
			turnCount++
			playerPostCount[player1] = 0
			playerPostCount[player2] = 0
		}
	}
}

*/
