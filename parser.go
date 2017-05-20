package main

import (
	"strings"
)

func parse(msg string) string {
	var response string
	if strings.HasPrefix(msg, "Hello") {
		response = "World"
	} else if strings.HasPrefix(strings.ToLower(msg), REMIND_ME) {
		response = remindMe(msg)
	}

	return response
}
