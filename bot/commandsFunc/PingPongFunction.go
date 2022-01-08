package commands

import ("github.com/bwmarrin/discordgo"
		"bots/GOing/options"
		"strings")

func PingPongFunc(channelSession *discordgo.Session, message *discordgo.MessageCreate){
	// If the message is "ping" reply with "Pong!"
	 if strings.ToUpper(message.Content) == options.Commands[0] {		//bothers me to check the same thing on two palces
		channelSession.ChannelMessageSend(message.ChannelID,"Pong!")    //but the less i hard code the better it is to
	}																	//change this function later 

	// If the message is "pong" reply with "Ping!"
	if strings.ToUpper(message.Content) == options.Commands[1] {
		channelSession.ChannelMessageSend(message.ChannelID, "Ping!")
	}
}