package main

import (
	"common"
	"fmt"
	"strings"
)

// ROCK - A
// PAPER - B
// SCISSORS - C

var picks = []rune{'A', 'B', 'C'}

var pick_scores = map[string]int{"A": 1, "B": 2, "C": 3}

var match_scores = map[string]int{"W": 6, "D": 3, "L": 0}

func getCrypticPick(opponent string, strategy string) string {
	switch strategy {
	case "X":
		return "A"
	case "Y":
		return "B"
	case "Z":
		return "C"
	default:
		panic("Invalid Strategy")
	}
}

func getDecryptedPick(opponent string, strategy string) string {
	opponentChar := []rune(opponent)[0]
	switch strategy {
	case "X":
		ascii_diff := (int(opponentChar) - int('A') - 1) % 3
		if ascii_diff < 0 {
			ascii_diff = 2
		}
		return fmt.Sprintf("%c", (ascii_diff + int('A')))
	case "Y":
		return opponent
	case "Z":
		ascii_diff := (int(opponentChar) - int('A') + 1) % 3
		return fmt.Sprintf("%c", (ascii_diff + int('A')))
	default:
		panic("Invalid Strategy")
	}
}

func getRoundScore(opponent string, you string, pickFunc func(string, string) string) int {
	yourPick := pickFunc(opponent, you)
	if (opponent == "A" && yourPick == "B") || (opponent == "B" && yourPick == "C") || (opponent == "C" && yourPick == "A") {

		return pick_scores[yourPick] + match_scores["W"]
	}
	if opponent == yourPick {

		return pick_scores[yourPick] + match_scores["D"]
	}

	return pick_scores[yourPick]
}

func getTotalScore(lines []string, pickFunc func(string, string) string) int {
	sum := 0
	for _, e := range lines {
		res := strings.Split(e, " ")
		opponent, you := res[0], res[1]
		sum += getRoundScore(opponent, you, pickFunc)
	}
	return sum
}

func main() {
	lines := common.ReadFile("input.txt")

	fmt.Println("Part 1: ", getTotalScore(lines, getCrypticPick))

	fmt.Println("Part 2: ", getTotalScore(lines, getDecryptedPick))
}

// func main() {
// 	fmt.Println(getDecryptedPick("A", "X")) // lose -> C
// 	fmt.Println(getDecryptedPick("A", "Y")) // draw -> A
// 	fmt.Println(getDecryptedPick("A", "Z")) // win -> B
// 	fmt.Println(getDecryptedPick("B", "X")) // lose -> A
// 	fmt.Println(getDecryptedPick("B", "Y")) // draw -> B
// 	fmt.Println(getDecryptedPick("B", "Z")) // win -> C
// 	fmt.Println(getDecryptedPick("C", "X")) // lose -> B
// 	fmt.Println(getDecryptedPick("C", "Y")) // draw -> C
// 	fmt.Println(getDecryptedPick("C", "Z")) // win -> A
// }
