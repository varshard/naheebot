package main

import (
	"fmt"
	"strings"
	"github.com/nlopes/slack"
	"github.com/jinzhu/configor"
)

//const BOT_ID = "U5FV0RATF"
//const AT_BOT = "<@" + BOT_ID + ">"
var api *slack.Client

var Config = struct {
	Token string
}{}

func main()  {
	configor.Load(&Config, "config.yml")
	api = slack.New(Config.Token)
	fmt.Print("Started\n")

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch evn := msg.Data.(type) {
		case *slack.MessageEvent:
			fmt.Printf("Message: %s, at %s\n", evn.Text, evn.Channel)
			handleCommand(evn.Text, evn.Channel)
		}
	}
}

func handleCommand(cmd string, channel string)  {
	var response string
	if strings.HasPrefix(cmd, "Hello") {
		response = "World"
	}

	channelId, _, err := api.PostMessage(channel, response, slack.PostMessageParameters{})

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message was sent to %s", channelId)
}
