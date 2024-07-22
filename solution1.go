package main

import "strconv"

func solve1(lines []string) (sum int) {
	// Simply summing all the numbers in the list
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		sum += num
	}
	return
}

func solve1P2(lines []string) int {
	frequencies := make([]int, len(lines))
	for i, line := range lines {
		num, _ := strconv.Atoi(line)
		frequencies[i] = num
	}

	// Simply finding the first duplicate for current_frequency
	current_frequency := 0
	found_frequencies := make(map[int]bool) // a simple set
	found_frequencies[current_frequency] = true
	for {
		for _, frequency := range frequencies {
			current_frequency += frequency

			if found_frequencies[current_frequency] {
				return current_frequency
			}

			found_frequencies[current_frequency] = true
		}
	}
}
