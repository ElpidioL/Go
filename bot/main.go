package main

import (
	"bots/GOing/GObot"
	"bots/GOing/options"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var Discord *discordgo.Session

func main() {

	// Create a new Discord session using the provided bot token.
	// if it fail to get discord from discordgo, then it will not be nill and will stop
	Discord, errDisc := discordgo.New("Bot " + options.Key)

	if errDisc != nil {
		fmt.Println("error creating Discord session,", errDisc)
		return
	}
	user, err := Discord.User("@me")
	if err != nil {
		fmt.Println("error setting Discord user,", errDisc)
		return
	}
	BotID = user.ID
	//everything is ok, time to get started
	GObot.Start(Discord, BotID)
}
