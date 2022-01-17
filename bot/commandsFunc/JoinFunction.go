package functions

import (
	"github.com/bwmarrin/discordgo"
	//"bots/GOing/options"
)

func JoinFunction(channelSession *discordgo.Session, message *discordgo.MessageCreate, ch string) {
	channelSession.ChannelVoiceJoin(message.GuildID, ch, false, true)
}
