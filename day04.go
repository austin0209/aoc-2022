package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Contains True if range1 contains range2.
func Contains(range1 []string, range2 []string) bool {
	lower1, _ := strconv.Atoi(range1[0])
	upper1, _ := strconv.Atoi(range1[1])

	lower2, _ := strconv.Atoi(range2[0])
	upper2, _ := strconv.Atoi(range2[1])

	return lower1 >= lower2 && upper1 <= upper2
}

func SolveDay04() {
	content, err := os.ReadFile("input/day04.txt")
	if err != nil {
		log.Fatal(err)
	}

	contentStr := string(content)

	lines := strings.Split(contentStr, "\n")

	total := 0

	for _, line := range lines {
		if line == ""{
			continue
		}

		pair := strings.Split(line, ",")

		range1 := strings.Split(pair[0], "-")
		range2 := strings.Split(pair[1], "-")

		if Contains(range1, range2) || Contains(range2, range1) {
			total++
		}
	}

	fmt.Println(total)
}
