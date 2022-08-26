package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Options struct {
	Verbose bool
}

type Subscription struct {
	Name      string `json:"name"`
	Id        string `json:"id"`
	IsDefault bool   `json:"isDefault"`
}

func main() {
	options := parseOptions()
	subs := getSubs()

	displaySubs(subs, options)

	newSub := selectSub(subs)
	if newSub == 0 {
		os.Exit(0)
	}

	setSub(subs, newSub)
	fmt.Printf("Changed subscription to: %s\n", subs[newSub-1].Name)
}

func parseOptions() Options {
	verbosePtr := flag.Bool("verbose", false, "Show more information about subscriptions")
	shortVerbosePtr := flag.Bool("v", false, "Show more information about subscriptions")
	flag.Parse()
	return Options{
		Verbose: *verbosePtr || *shortVerbosePtr,
	}
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

func displaySubs(subs []Subscription, options Options) {
	if options.Verbose {
		displayLongSubs(subs)
	} else {
		displayShortSubs(subs)
	}
}

func displayLongSubs(subs []Subscription) {
	maxNameLength := maxSubName(subs)

	for index, sub := range subs {
		fmt.Println(longDisplayName(sub, index, maxNameLength-len(sub.Name)))
	}
}

func maxSubName(subs []Subscription) int {
	maxNameLength := 0
	for _, sub := range subs {
		if len(sub.Name) > maxNameLength {
			maxNameLength = len(sub.Name)
		}
	}
	return maxNameLength
}

func displayShortSubs(subs []Subscription) {
	for index, sub := range subs {
		fmt.Println(baseDisplayName(sub, index))
	}
}

func longDisplayName(sub Subscription, index int, padding int) string {
	baseDisplayName := baseDisplayName(sub, index)
	return baseDisplayName + strings.Repeat(" ", padding+1) + sub.Id
}

func baseDisplayName(sub Subscription, index int) string {
	prefix := subPrefix(sub)
	return fmt.Sprintf("%s%d. %s", prefix, index+1, sub.Name)
}

func subPrefix(sub Subscription) string {
	if sub.IsDefault {
		return "*"
	}
	return " "
}

func selectSub(subs []Subscription) int {
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
