package main

import (
	"encoding/json"
	"log"
	"os/exec"
)

type Subscription struct {
	Name      string `json:"name"`
	Id        string `json:"id"`
	IsDefault bool   `json:"isDefault"`
}

func getSubs() []Subscription {
	cmd := exec.Command("az", "account", "list")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	var subs []Subscription
	err = json.Unmarshal(out, &subs)
	if err != nil {
		log.Fatal(err)
	}
	return subs
}

func setSub(subs []Subscription, newSub int) {
	cmd := exec.Command("az", "account", "set", "--subscription", subs[newSub-1].Id)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
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
