package main

func solve2(lines []string) int {
	totalTwos := 0
	totalThrees := 0

	for _, line := range lines {
		characters := make(map[rune]int)
		for _, char := range line {
			characters[char]++
		}

		hasTwo := 0
		hasThree := 0
		for _, count := range characters {
			if hasTwo != 1 && count == 2 {
				hasTwo = 1
			}

			if hasThree != 1 && count == 3 {
				hasThree = 1
			}
		}

		totalTwos += hasTwo
		totalThrees += hasThree
	}

	return totalTwos * totalThrees
}

func solve2P2(lines []string) string {
	IDLength := len(lines[0])

	// n^3 algorithm but it's probably still fast enough
	for i := 0; i < len(lines); i++ {
		for j := 0; j < i; j++ {
			differences := 0
			dPos := 0

			for k := 0; k < IDLength; k++ {
				if lines[i][k] != lines[j][k] {
					differences++
					dPos = k
				}
			}

			if differences == 1 {
				bytes := []byte(lines[i])
				bytes = append(bytes[:dPos], bytes[dPos+1:]...)
				return string(bytes)
			}
		}
	}
	// Turns out it's easily fast enough, there's only ~250 lines

	return ""
}
