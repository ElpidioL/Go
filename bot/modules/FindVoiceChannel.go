package modules

import (
	"github.com/bwmarrin/discordgo"
)

func FindVoiceChannel(channelSession *discordgo.Session, guildID string, authorID string) string {
	var x string
	//fmt.Println(channelSession.State.Guild(message.GuildID))  i'll try it later
	for _, guild := range channelSession.State.Guilds {
		if guild.ID == guildID {
			for _, vs := range guild.VoiceStates {
				if vs.UserID == authorID {
					x = vs.ChannelID
				}
			}
		}
	}
	return x
}
