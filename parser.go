package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"strings"
)

func parse(msg string, channel string) {
	var response string
	if strings.HasPrefix(msg, "Hello") {
		response = "World"
	} else if strings.HasPrefix(strings.ToLower(msg), REMIND_ME) {
		response = remindMe(msg)
	}

	channelId, _, err := api.PostMessage(channel, response, slack.PostMessageParameters{})

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message was sent to %s\n", channelId)
}
