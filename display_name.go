package main

import (
	"fmt"
	"strings"
)

func displayName(sub Subscription, index int, config Config) string {
	prefix := subPrefix(sub)
	base := fmt.Sprintf("%s%d. %s", prefix, index+1, sub.Name)
	postfix := postfix(sub, config)
	return base + postfix
}

func postfix(sub Subscription, config Config) string {
	if !config.Verbose {
		return ""
	}
	padding := config.NameWidth - len(sub.Name)
	return strings.Repeat(" ", padding) + sub.Id
}

func subPrefix(sub Subscription) string {
	if sub.IsDefault {
		return "*"
	}
	return " "
}
