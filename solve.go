package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the location of the input file as the first argument.")
		return
	}

	dat, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Printf("Failed to read the input file. Did you download it to %s?", os.Args[1])
		return
	}

	lines := strings.Split(string(dat), "\n")
	fmt.Println(solve1(lines))
}
