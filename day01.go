package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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
	var maxCalories []int

	for _, calories := range calorieSets {
		itemsStr := strings.Split(calories, "\n")
		totalCalories := 0

		for _, itemStr := range itemsStr {
			itemCalories, _ := strconv.Atoi(itemStr)

			totalCalories += itemCalories
		}

		maxCalories = append(maxCalories, totalCalories)
	}

	sort.Ints(maxCalories)

	topThree := maxCalories[len(maxCalories)-3:]

	fmt.Println(topThree[0] + topThree[1] + topThree[2])
}
