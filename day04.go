package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Contains True if range1 contains range2 OR vice versa.
func Contains(range1 []string, range2 []string) bool {
	lower1, _ := strconv.Atoi(range1[0])
	upper1, _ := strconv.Atoi(range1[1])

	lower2, _ := strconv.Atoi(range2[0])
	upper2, _ := strconv.Atoi(range2[1])

	return lower1 >= lower2 && upper1 <= upper2 || lower2 >= lower1 && upper2 <= upper1
}

func Overlaps(range1 []string, range2 []string) bool {
	lower1, _ := strconv.Atoi(range1[0])
	upper1, _ := strconv.Atoi(range1[1])

	lower2, _ := strconv.Atoi(range2[0])
	upper2, _ := strconv.Atoi(range2[1])

	return Contains(range1, range2) || (lower1 >= lower2 && lower1 <= upper2) || (upper1 <= upper2 && upper1 >= lower2)
}

func SolveDay04() {
	content, err := os.ReadFile("input/day04.txt")
	if err != nil {
		log.Fatal(err)
	}

	contentStr := string(content)

	lines := strings.Split(contentStr, "\n")

	part1Total := 0
	part2Total := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		pair := strings.Split(line, ",")

		range1 := strings.Split(pair[0], "-")
		range2 := strings.Split(pair[1], "-")

		if Contains(range1, range2) {
			part1Total++
		}

		if Overlaps(range1, range2) {
			part2Total++
		}
	}

	fmt.Printf("Part 1: %d\nPart 2: %d", part1Total, part2Total)
}
