package functions

import (
	"bots/GOing/modules"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func LoopAeternum(channelSession *discordgo.Session, message *discordgo.MessageCreate) {
	modules.SendMessage(channelSession, message.ChannelID, "Verify Go handling multiple requests", false)
	for i := 0; i < 5; i++ {
		str := strconv.Itoa(i)
		modules.SendMessage(channelSession, message.ChannelID, str, false)
		time.Sleep(1 * time.Second)
	}
}
