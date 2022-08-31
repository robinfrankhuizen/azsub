package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	newSubRaw, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return newSubRaw
}

func validateInput(newSubRaw string, subs []Subscription) int {
	newSub, err := parseInput(newSubRaw)
	if err != nil || newSub < 0 || newSub > len(subs) {
		fmt.Printf("Please enter a number between 0 and %d: ", len(subs))
		return validateInput(readInput(), subs)
	}
	return newSub
}

func parseInput(rawInput string) (int, error) {
	cleanedInput := cleanInput(rawInput)
	return strconv.Atoi(cleanedInput)
}

func cleanInput(rawInput string) string {
	input := rawInput
	input = strings.TrimRight(input, "\r\n")
	input = strings.ReplaceAll(input, " ", "")
	return input
}
