package modules

import (
		"github.com/bwmarrin/discordgo"
		"bots/GOing/options"
		"strings"
		//"fmt"
		"bots/GOing/commandsFunc"
	)

func ReceiveMessage(channelSession *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if message.Author.ID == channelSession.State.User.ID{
		return
	}

	if strings.ToUpper(message.Content) == options.Commands[0] || strings.ToUpper(message.Content) == options.Commands[1]{ // 0 == !ping, 1 == !pong
		commands.PingPongFunc(channelSession,message)
	}

	if strings.ToUpper(message.Content) == options.Commands[2] { // 2 == !help
		channelSession.ChannelMessageSend(message.ChannelID, "Sim, ele é")
	}  
}