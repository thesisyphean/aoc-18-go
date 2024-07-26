package main

import (
	"math"
	"strings"
)

// Simple doubly-linked list
type Polymer struct {
	unit byte
	prev *Polymer
	next *Polymer
}

// This will error on a string of length less than 2
func newPolymer(polymer string) *Polymer {
	head := &Polymer{polymer[0], nil, nil}
	current := &Polymer{polymer[1], head, nil}
	head.next = current

	for i := 2; i < len(polymer); i++ {
		newNode := &Polymer{polymer[i], current, nil}
		current.next = newNode
		current = newNode
	}

	return head
}

func (p *Polymer) String() string {
	result := make([]byte, 0)

	current := p
	for current.prev != nil {
		current = current.prev
	}

	for ; current != nil; current = current.next {
		result = append(result, current.unit)
	}

	return string(result)
}

func collides(a, b byte) bool {
	return math.Abs(float64(a)-float64(b)) == 32
}

func unitsAfterReaction(current *Polymer, units int) int {
	for current.next != nil {
		if collides(current.unit, current.next.unit) {
			units -= 2

			if current.next.next == nil {
				return units
			}

			if current.prev == nil {
				current = current.next.next
				current.prev = nil
			} else {
				current.prev.next = current.next.next
				current.next.next.prev = current.prev
				current = current.prev
			}
		} else {
			current = current.next
		}
	}

	return units
}

func solve5(lines []string) int {
	current := newPolymer(lines[0])
	return unitsAfterReaction(current, len(lines[0]))
}

// TODO: This is the most naive approach possible
func solve5P2(lines []string) int {
	polymerString := lines[0]
	bestUnits := len(polymerString)

	for char := 'a'; char <= 'z'; char++ {
		firstFilter := strings.ReplaceAll(polymerString, string(char), "")
		secondFilter := strings.ReplaceAll(firstFilter, strings.ToUpper(string(char)), "")
		units := unitsAfterReaction(newPolymer(secondFilter), len(secondFilter))

		if units < bestUnits {
			bestUnits = units
		}
	}

	return bestUnits
}
