package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func SolveDay02() {
	content, err := os.ReadFile("input/day02.txt")
	if err != nil {
		log.Fatal(err)
	}

	contentStr := string(content)

	lines := strings.Split(contentStr, "\n")

	var scores []int

	for _, line := range lines {
		var score int

		switch line {
		case "A X":
			score = 4
		case "A Y":
			score = 8
		case "A Z":
			score = 3
		case "B X":
			score = 1
		case "B Y":
			score = 5
		case "B Z":
			score = 9
		case "C X":
			score = 7
		case "C Y":
			score = 2
		case "C Z":
			score = 6
		}

		scores = append(scores, score)
	}

	totalScore := 0
	for _, score := range scores {
		totalScore += score
	}

	fmt.Println(totalScore)
}
