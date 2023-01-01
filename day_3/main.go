package main

import (
	"common"
	"fmt"
)

func computePrioritySum(lines []string) int {
	sum := 0
	for _, rucksack := range lines {
		same_char := getSameCharInBothCompartments(rucksack)
		sum += computePriority(same_char)
	}
	return sum
}

func computePriority(same_char rune) int {
	if int('a') <= int(same_char) && int(same_char) <= int('z') {
		return int(same_char) - int('a') + 1
	}
	return int(same_char) - int('A') + 27
}

func getSameCharInBothCompartments(rucksack string) rune {
	length := len(rucksack)
	compartment1 := rucksack[0 : length/2]
	compartment2 := rucksack[length/2:]

	m := make(map[rune]bool)

	for _, char := range compartment1 {
		m[char] = true
	}

	for _, char := range compartment2 {
		if m[char] {
			return char
		}
	}

	panic("same char not found")
}

func computePrioritySumForGroup(lines []string) int {
	sum := 0
	for i := 3; i <= len(lines); i += 3 {
		same_char := getSameCharInGroup(lines[i-3 : i])
		sum += computePriority(same_char)
	}
	return sum
}

func getSameCharInGroup(rucksacks []string) rune {

	m := make(map[rune]uint)

	elf_1, elf_2, elf_3 := rucksacks[0], rucksacks[1], rucksacks[2]

	for _, char := range elf_1 {
		m[char] = 1
	}

	for _, char := range elf_2 {
		if m[char] == 1 {
			m[char] = 2
		}
	}

	for _, char := range elf_3 {
		if m[char] == 2 {
			return char
		}
	}

	panic("same char not found")
}

func main() {
	lines := common.ReadFile("input.txt")
	fmt.Println("Part 1: ", computePrioritySum(lines))
	fmt.Println("Part 2: ", computePrioritySumForGroup(lines))
	// for i := 3; i <= len(lines); i += 3 {
	// 	fmt.Println(lines[i-3 : i])
	// }
}
