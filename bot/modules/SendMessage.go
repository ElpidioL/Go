package modules

import (
	"github.com/bwmarrin/discordgo"
)

func SendMessage(channelSession *discordgo.Session, channelID string, message string, TTS bool) {
	if TTS {
		channelSession.ChannelMessageSendTTS(channelID, message)
	} else {
		channelSession.ChannelMessageSend(channelID, message)
	}
}
