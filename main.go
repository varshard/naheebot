package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/nlopes/slack"
)

//const BOT_ID = "U5FV0RATF"
//const AT_BOT = "<@" + BOT_ID + ">"
var api *slack.Client

var Config = struct {
	Token string
}{}

func main() {
	configor.Load(&Config, "config.yml")
	api = slack.New(Config.Token)
	fmt.Print("Started\n")
	botID, err := getBotID(api)

	if err != nil {
		panic(fmt.Sprintf("Error: %s\n", err))
	}

	//botName := "<@" + botID + ">"
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch evn := msg.Data.(type) {
		case *slack.MessageEvent:
			//TODO: prevent bot's message triggering the event
			//fmt.Printf("Message: %s, %s, %s\n", evn.Text, evn.Channel, evn.User == BOT_ID)
			if evn.User != botID {
				response := parse(evn.Text)
				fmt.Printf("response: %s\n", response)
				rtm.SendMessage(rtm.NewOutgoingMessage(response, evn.Channel))
			}
		}
	}
}

func getBotID(api *slack.Client) (string, error) {
	users, err := api.GetUsers()
	var userID string
	if err != nil {
		fmt.Printf("errror: %s\n", err)
		return "", err
	}

	for _, user := range users {
		if (user.Name == "naheebot") {
			userID = user.ID
		}
	}

	return userID, nil
}
