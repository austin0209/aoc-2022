package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func SolveTemplate() {
	content, err := os.ReadFile("input/REPLACEME.txt")
	if err != nil {
		log.Fatal(err)
	}

	contentStr := string(content)

	lines := strings.Split(contentStr, "\n")

	for _, line := range lines {
		fmt.Println(line)
	}
}
