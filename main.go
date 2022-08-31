package main

import (
	"fmt"
	"os"
)

func main() {
	flags := parseFlags()
	subs := getSubs()
	if len(subs) == 0 {
		fmt.Println("Please run \"az login\" to login")
		os.Exit(0)
	}

	config := createConfig(subs, flags)

	fmt.Println("Subscriptions: ")
	for index, sub := range subs {
		fmt.Println(displayName(sub, index, config))
	}
	fmt.Print("Select subscription to use or press 0 to exit: ")

	newSubRaw := readInput()
	newSub := validateInput(newSubRaw, subs)
	if newSub == 0 {
		os.Exit(0)
	}

	setSub(subs, newSub)
	fmt.Printf("Changed subscription to: %s\n", subs[newSub-1].Name)
}
