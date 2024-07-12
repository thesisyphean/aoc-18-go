package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Solving!")

	dat, _ := os.ReadFile("/input/input1.txt")
	lines := strings.Split(string(dat), "\n")
	fmt.Println(solve1(lines))
}
