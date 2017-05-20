package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/nlopes/slack"
)

const BOT_ID = "U5FV0RATF"
//const AT_BOT = "<@" + BOT_ID + ">"
var api *slack.Client

var Config = struct {
	Token string
}{}

func main() {
	configor.Load(&Config, "config.yml")
	api = slack.New(Config.Token)
	fmt.Print("Started\n")

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch evn := msg.Data.(type) {
		case *slack.MessageEvent:
			//TODO: prevent bot's message triggering the event
			//fmt.Printf("Message: %s, %s, %s\n", evn.Text, evn.Channel, evn.User == BOT_ID)
			if evn.User != BOT_ID {
				response := parse(evn.Text)
				fmt.Printf("response: %s\n", response)
				rtm.SendMessage(rtm.NewOutgoingMessage(response, evn.Channel))
			}
		}
	}
}
