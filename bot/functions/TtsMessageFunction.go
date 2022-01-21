package functions

import (
	"bots/GOing/modules"
	"bots/GOing/options"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func TtsMessageFunc(channelSession *discordgo.Session, message *discordgo.MessageCreate) {
	modules.SendMessage(channelSession, message.ChannelID, ";"+strings.Replace(message.Content, options.Commands[3], "", -1), true)
}
