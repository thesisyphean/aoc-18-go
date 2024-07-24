package main

import (
	"strconv"
	"strings"
)

const fabricSize = 1_000

type Rectangle struct {
	ID, left, top, width, height int
}

func parseRectangle(line string) Rectangle {
	parts := strings.Split(line, " ")

	ID, _ := strconv.Atoi(parts[0][1:])

	corner := strings.Split(parts[2], ",")
	dimensions := strings.Split(parts[3], "x")

	left, _ := strconv.Atoi(corner[0])
	top, _ := strconv.Atoi(corner[1][:len(corner[1])-1])
	width, _ := strconv.Atoi(dimensions[0])
	height, _ := strconv.Atoi(dimensions[1])

	return Rectangle{ID, left, top, width, height}
}

func solve3(lines []string) int {
	array := make([]int, fabricSize*fabricSize)

	for _, line := range lines {
		rect := parseRectangle(line)

		for x := rect.left; x < rect.left+rect.width; x++ {
			for y := rect.top; y < rect.top+rect.height; y++ {
				array[y*fabricSize+x]++
			}
		}
	}

	count := 0
	for i := 0; i < len(array); i++ {
		if array[i] > 1 {
			count++
		}
	}
	return count
}

func solve3P2(lines []string) int {
	array := make([]int, fabricSize*fabricSize)
	rects := make([]Rectangle, len(lines))

	for i, line := range lines {
		rect := parseRectangle(line)
		rects[i] = rect

		for x := rect.left; x < rect.left+rect.width; x++ {
			for y := rect.top; y < rect.top+rect.height; y++ {
				array[y*fabricSize+x]++
			}
		}
	}

outer:
	for _, rect := range rects {
		for x := rect.left; x < rect.left+rect.width; x++ {
			for y := rect.top; y < rect.top+rect.height; y++ {
				if array[y*fabricSize+x] > 1 {
					continue outer
				}
			}
		}
		return rect.ID
	}

	return -1
}
