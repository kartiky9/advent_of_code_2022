package main

import (
	"common"
	"fmt"
	"strings"
)

func main() {
	lines := common.ReadFile("input.txt")
	count := 0
	for _, el := range lines {
		elves_range := strings.Split(el, ",")
		range_1, range_2 := strings.Split(elves_range[0], "-"), strings.Split(elves_range[1], "-")
		if (range_1[0] <= range_2[0] && range_1[1] >= range_2[1]) ||
			(range_1[0] >= range_2[0] && range_1[1] <= range_2[1]) {
			count += 1
		}
	}
	fmt.Println("Part 1 : ", count)
}
