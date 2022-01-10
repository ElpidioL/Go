package functions

import (
	"github.com/bwmarrin/discordgo"
	//"bots/GOing/options"
)

func VoiceFunc(channelSession *discordgo.Session, message *discordgo.MessageCreate, ch string) {
	channelSession.ChannelVoiceJoin(message.GuildID, ch, false, true)
}
