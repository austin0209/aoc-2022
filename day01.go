package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input/day01.txt")
	if err != nil {
		log.Fatal(err)
	}

	contentStr := string(content)

	calorieSets := strings.Split(contentStr, "\n\n")
	var maxCalorie int

	for _, calories := range calorieSets {
		itemsStr := strings.Split(calories, "\n")
		totalCalories := 0

		for _, itemStr := range itemsStr {
			itemCalories, _ := strconv.Atoi(itemStr)

			totalCalories += itemCalories
		}

		if totalCalories > maxCalorie {
			maxCalorie = totalCalories
		}
	}

	fmt.Println(maxCalorie)
}
