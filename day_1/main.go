package main

import (
	"common"
	"fmt"
	"strconv"
)

func max(a int, b int) int {
	if b > a {
		return b
	}
	return a
}

func maxCaloriesBeingCarried(lines []string) int {

	maxSum := 0
	currentSum := 0
	for _, calorieStr := range lines {
		if calorieStr == "" {
			maxSum = max(maxSum, currentSum)
			currentSum = 0
			continue
		}

		calorie, _ := strconv.Atoi(calorieStr)
		currentSum += calorie
	}

	return maxSum
}

func top3Sum(lines []string) int {
	top3 := make([]int, 3)
	currentSum := 0
	for _, calorieStr := range lines {
		if calorieStr == "" {
			if currentSum > top3[0] {
				top3[2], top3[1], top3[0] = top3[1], top3[0], currentSum
			} else if currentSum > top3[1] {
				top3[2], top3[1] = top3[1], currentSum
			} else if currentSum > top3[2] {
				top3[2] = currentSum
			}
			currentSum = 0
			continue
		}

		calorie, _ := strconv.Atoi(calorieStr)
		currentSum += calorie
	}
	sum := 0
	for _, e := range top3 {
		sum += e
	}
	return sum
}

func main() {
	lines := common.ReadFile("input.txt")
	fmt.Printf("Max Calories Carried : %d\n", maxCaloriesBeingCarried(lines))
	fmt.Printf("Top3 Calories : %d\n", top3Sum(lines))
}
