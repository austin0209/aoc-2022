package main

import (
	"fmt"
	"log"
	"os"
)

func SolveDay06() {
	content, err := os.ReadFile("input/day06.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	for i := 0; i < len(input)-4; i++ {
		c1, c2, c3, c4 := input[i], input[i+1], input[i+2], input[i+3]

		if c1 == c2 || c1 == c3 || c1 == c4 || c2 == c3 || c2 == c4 || c3 == c4 {
			continue
		}

		fmt.Println(i + 4)
		break
	}
}
