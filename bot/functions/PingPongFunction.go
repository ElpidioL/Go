package functions

import (
	"bots/GOing/modules"
	"bots/GOing/options"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func PingPongFunc(channelSession *discordgo.Session, message *discordgo.MessageCreate) {
	// If the message is "ping" reply with "Pong!"
	if strings.ToUpper(message.Content) == options.Commands[0] { //bothers me to check the same thing on two palces
		modules.SendMessage(channelSession, message.ChannelID, options.Commands[1], false) //but the less i hard code the better it is to
	} //change this function later

	// If the message is "pong" reply with "Ping!"
	if strings.ToUpper(message.Content) == options.Commands[1] {
		modules.SendMessage(channelSession, message.ChannelID, options.Commands[0], false)
	}

}
