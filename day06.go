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

	for i := 0; i < len(input)-14; i++ {
		seen := make(map[uint8]struct{})

		mark := true

		for j := i; j < i+14; j++ {
			_, exists := seen[input[j]]

			if exists {
				mark = false
				break
			}

			seen[input[j]] = struct{}{}
		}

		if !mark {
			continue
		}

		fmt.Println(i + 14)
		break
	}
}
