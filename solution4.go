package main

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

type Guard struct {
	id           int
	minutes      map[int]int
	totalMinutes int
}

const (
	startsShift = iota
	fallsAsleep = iota
	wakesUp     = iota
)

type Message struct {
	timestamp time.Time
	kind      int
	guard     *Guard
}

func parseGuards(lines []string) []Guard {
	messages := make([]Message, len(lines))
	// This is to ensure that appending to the slice does not reallocate
	// the underlying array and therefore ruin all the pointers
	// Absolutely hellish bug, my God
	guards := make([]Guard, len(lines))
	guardMap := make(map[int]*Guard)

	for i, line := range lines {
		referenceTime := "[2006-01-02 15:04]"
		timestamp, _ := time.Parse(referenceTime, line[:18])

		kind := fallsAsleep
		var guardRef *Guard
		if line[19] == 'G' {
			kind = startsShift
			ID, _ := strconv.Atoi(strings.Split(line[19:], " ")[1][1:])

			guard, prs := guardMap[ID]
			if prs {
				guardRef = guard
			} else {
				guards = append(guards, Guard{ID, make(map[int]int), 0})
				guardMap[ID] = &guards[len(guards)-1]
				guardRef = &guards[len(guards)-1]
			}
		} else if line[19] == 'w' {
			kind = wakesUp
		}
		messages[i] = Message{timestamp, kind, guardRef}
	}

	sort.Slice(messages, func(i, j int) bool {
		if messages[i].timestamp == messages[j].timestamp {
			return messages[i].kind < messages[j].kind
		}
		return messages[i].timestamp.Before(messages[j].timestamp)
	})

	var currentGuard *Guard
	var currentSnooze time.Time
	for _, message := range messages {
		if message.kind == startsShift {
			currentGuard = message.guard
		} else if message.kind == fallsAsleep {
			currentSnooze = message.timestamp
		} else {
			start := currentSnooze.Minute()
			end := message.timestamp.Minute()
			currentGuard.totalMinutes += end - start
			for t := start; t < end; t++ {
				currentGuard.minutes[t]++
			}
		}
	}

	return guards
}

func solve4(lines []string) int {
	guards := parseGuards(lines)

	sleepiestGuard := &guards[0]
	for _, guard := range guards[1:] {
		if guard.totalMinutes > sleepiestGuard.totalMinutes {
			sleepiestGuard = &guard
		}
	}

	mostCommonMinute := 0
	maxTimes := 0
	for minute, times := range sleepiestGuard.minutes {
		if times > maxTimes {
			mostCommonMinute = minute
			maxTimes = times
		}
	}
	return sleepiestGuard.id * mostCommonMinute
}

func solve4P2(lines []string) int {
	guards := parseGuards(lines)

	ID := 0
	bestMinute := 0
	maxFrequency := 0
	for _, guard := range guards {
		for minute, frequency := range guard.minutes {
			if frequency > maxFrequency {
				ID = guard.id
				bestMinute = minute
				maxFrequency = frequency
			}
		}
	}

	return ID * bestMinute
}
