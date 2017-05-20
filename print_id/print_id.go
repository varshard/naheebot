package main

import (
	"fmt"
	"github.com/nlopes/slack"
)

func main() {
	fmt.Print("main\n")
	api := slack.New("xoxb-185986860933-8Heehy7rurZMRi0smbK7XaVl")
	users, err := api.GetUsers()
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("users %d\n", len(users))
	for _, user := range users {
		fmt.Printf("ID: %s, Name %s\n", user.ID, user.Name)
	}
}
