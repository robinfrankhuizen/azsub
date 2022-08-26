package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Subscription struct {
	Name      string `json:"name"`
	Id        string `json:"id"`
	IsDefault bool   `json:"isDefault"`
}

func main() {
	subs := getSubs()

	newSub := selectSub(subs)
	if newSub == 0 {
		os.Exit(0)
	}

	setSub(subs, newSub)
	fmt.Printf("Changed subscription to: %s\n", subs[newSub-1].Name)
}

func getSubs() []Subscription {
	cmd := exec.Command("az", "account", "list")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	if cleanInput(string(out)) == "[]" {
		fmt.Println("Please run \"az login\" to login")
		os.Exit(0)
	}

	var subs []Subscription
	err = json.Unmarshal(out, &subs)
	if err != nil {
		log.Fatal(err)
	}
	return subs
}

func selectSub(subs []Subscription) int {
	for index, sub := range subs {
		fmt.Println(displayName(sub, index))
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select subscription to use or press 0 to exit: ")
	newSubRaw, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	newSub, err := parseIndex(newSubRaw)
	if err != nil || newSub < 0 || newSub > len(subs) {
		fmt.Printf("Please enter a number between 0 and %d\n", len(subs))
		selectSub(subs)
	}
	return newSub
}

func setSub(subs []Subscription, newSub int) {
	cmd := exec.Command("az", "account", "set", "--subscription", subs[newSub-1].Id)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func displayName(sub Subscription, index int) string {
	var prefix string
	if sub.IsDefault {
		prefix = "*"
	} else {
		prefix = " "
	}
	return fmt.Sprintf("%s%d. %s", prefix, index+1, sub.Name)
}

func parseIndex(rawInput string) (int, error) {
	cleanedInput := cleanInput(rawInput)
	return strconv.Atoi(cleanedInput)
}

func cleanInput(rawInput string) string {
	input := rawInput
	input = strings.TrimRight(input, "\r\n")
	input = strings.ReplaceAll(input, " ", "")
	return input
}
