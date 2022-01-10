package modules

import (
	"github.com/bwmarrin/discordgo"
)

func FindVoiceChannel(channelSession *discordgo.Session, message *discordgo.MessageCreate) string {
	var x string
	//fmt.Println(channelSession.State.Guild(message.GuildID))  i'll try it later
	for _, guild := range channelSession.State.Guilds {
		if guild.ID == message.GuildID {
			for _, vs := range guild.VoiceStates {
				if vs.UserID == message.Author.ID {
					x = vs.ChannelID
				}
			}
		}
	}
	return x
}
