package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	pathPtr := flag.String("path", "", "The path to the input file from input/")
	flag.Parse()

	if *pathPtr == "" {
		fmt.Println("Please provide the path to the input file from input/ as a flag.")
		return
	}

	dat, err := os.ReadFile(fmt.Sprintf("input/%s", *pathPtr))

	if err != nil {
		fmt.Printf("Failed to read the input file. Did you download it to %s?", *pathPtr)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	fmt.Println(solve2P2(lines))
}
