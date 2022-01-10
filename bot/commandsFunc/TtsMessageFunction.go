package functions

import ("github.com/bwmarrin/discordgo"
		"bots/GOing/options"
		"strings"
	)

func TtsMessageFunc(channelSession *discordgo.Session, message *discordgo.MessageCreate){
	channelSession.ChannelMessageSendTTS(message.ChannelID, ";" + strings.Replace(message.Content, options.Commands[3],"",-1))
}	
