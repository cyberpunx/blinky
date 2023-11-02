package main

import (
	"fmt"
	"localdev/HrHelper/util"
	"strconv"
	"time"
)

func ClubPotionsProcessor(thread Thread, hoursLimit int) {
	turnCount := 1
	player1 := ""
	player2 := ""
	moderatorUsername := thread.Author.Username
	authorRole := ""
	timeLimitStatus := ""
	postCount := 1
	diceSum := 0
	turnDice := ""

	// Initialize maps to count each player's posts and store the time of the last post by each player
	playerPostCount := make(map[string]int)
	lastPostTime := *thread.Created

	player1PostCount := 0
	player2PostCount := 0

	timeThreshold := time.Duration(hoursLimit) * time.Hour

	// Iterate through the posts to identify players and count turns
	for _, post := range thread.Posts {
		authorUsername := post.Author.Username

		if authorUsername == moderatorUsername {
			authorRole = Cyan + "Moderator" + Reset
		}

		// Identify the players
		if player1 == "" && authorUsername != moderatorUsername {
			player1 = authorUsername
		} else if player2 == "" && authorUsername != player1 && authorUsername != moderatorUsername {
			player2 = authorUsername
		}

		// Count the post for the current player and update the last post time
		if authorUsername != moderatorUsername {
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

		if authorUsername != moderatorUsername {
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
		if authorUsername == moderatorUsername {
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
