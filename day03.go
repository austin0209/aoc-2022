package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GetPriority(c int32) (int32, error) {
	if c >= 'a' && c <= 'z' {
		return c - 'a' + 1, nil
	} else if c >= 'A' && c <= 'Z' {
		return c - 'A' + 27, nil
	} else {
		return -1, fmt.Errorf("cannot get priority from %c", c)
	}
}

func SolveDay03() {
	content, err := os.ReadFile("input/day03.txt")
	if err != nil {
		log.Fatal(err)
	}

	contentStr := string(content)

	lines := strings.Split(contentStr, "\n")

	var total int32

	for i := 0; i < len(lines) - 2; i += 3 {
		line1, line2, line3 := lines[i], lines[i + 1], lines[i + 2]

		for _, c := range line1 {
			if strings.Contains(line2, string(c)) && strings.Contains(line3, string(c)) {
				priority, err := GetPriority(c)
				if err != nil {
					log.Fatal(err)
				}

				total += priority
				break
			}
		}
	}

	fmt.Println(total)
}
