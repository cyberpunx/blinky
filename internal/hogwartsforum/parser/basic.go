package parser

import (
	"localdev/HrHelper/internal/util"
	"strconv"
	"strings"
)

func ParseDiceRoll(dicerolls []string) []*Dice {
	var dices []*Dice

	for _, diceroll := range dicerolls {
		resultStr := strings.TrimSpace(strings.Split(diceroll, ":")[len(strings.Split(diceroll, ":"))-1])
		//convert result from string to int
		result, err := strconv.Atoi(resultStr)
		util.Panic(err)
		dice := &Dice{
			DiceLine: diceroll,
			Result:   result,
		}
		dices = append(dices, dice)
	}
	return dices
}
