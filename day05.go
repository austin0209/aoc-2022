package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SolveDay05() {
	content, err := os.ReadFile("input/day05.txt")
	if err != nil {
		log.Fatal(err)
	}

	contentStr := string(content)

	input := strings.Split(contentStr, "\n\n")

	rawStacks := strings.Split(input[0], "\n")
	instructions := strings.Split(input[1], "\n")

	var stacks [][]string

	stacks = append(stacks, make([]string, 0))

	for _, s := range rawStacks {
		var newStack []string
		for _, c := range s {
			newStack = append(newStack, string(c))
		}

		stacks = append(stacks, newStack)
	}

	for _, s := range instructions[:len(instructions)-1] {
		tmp := strings.Split(s, " ")

		move, _ := strconv.Atoi(tmp[1])
		from, _ := strconv.Atoi(tmp[3])
		to, _ := strconv.Atoi(tmp[5])

		fromStack := &stacks[from]
		toStack := &stacks[to]

		removed := (*fromStack)[len(*fromStack)-move:]

		for i, _ := range removed {
			*toStack = append(*toStack, removed[i])
		}

		*fromStack = (*fromStack)[:len(*fromStack)-move]
	}

	fmt.Println(stacks)

	for _, s := range stacks {
		if len(s) > 0 {
			fmt.Print(s[len(s)-1])
		}
	}
}
