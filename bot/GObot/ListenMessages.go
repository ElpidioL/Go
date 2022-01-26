package GObot

import (
	API "bots/GOing/API"
	db "bots/GOing/PostgreDB"
	modules "bots/GOing/modules"
	"bots/GOing/options"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	functions "bots/GOing/functions"
	"time"
)

//strings.HasPrefix  dont see the reason to use, but its nice to keep here for later maybe

func ReceiveMessage(channelSession *discordgo.Session, message *discordgo.MessageCreate) {
	messageToUpper := strings.ToUpper(message.Content)
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if message.Author.ID == channelSession.State.User.ID {
		if strings.HasPrefix(messageToUpper, ";") { // ; will be the prefix to messages that i want to delete from the bot
			time.Sleep(3 * time.Second)
			channelSession.ChannelMessageDelete(message.ChannelID, message.ID)
		}
		return

	} else if messageToUpper == options.Commands[0] || messageToUpper == options.Commands[1] { // 0 == !ping, 1 == !pong
		functions.PingPongFunc(channelSession, message)

	} else if messageToUpper == options.Commands[2] { // 2 == !help
		channelSession.ChannelMessageSend(message.ChannelID, strings.Join(options.Commands, ", "))

	} else if strings.Contains(messageToUpper, options.Commands[3]) { //!TTS
		functions.TtsMessageFunc(channelSession, message)

	} else if messageToUpper == options.Commands[4] { //!JOIN
		functions.JoinFunction(channelSession, message, modules.FindVoiceChannel(channelSession, message.GuildID, message.Author.ID))

	} else if messageToUpper == options.Commands[5] { //!PLAY
		functions.PlayHorn(channelSession, message.GuildID, modules.FindVoiceChannel(channelSession, message.GuildID, message.Author.ID))

	} else if messageToUpper == options.Commands[6] { //!LEAVE
		channelSession.ChannelVoiceJoin("", "", false, true)

	} else if messageToUpper == options.Commands[7] { //!LOOP
		functions.LoopAeternum(channelSession, message)

	} else if strings.Contains(messageToUpper, options.Commands[8]) { //!LOL  (register)
		userN := strings.Replace(messageToUpper, "!LOL ", "", -1)
		msg := API.GetUserLol(userN, message.ChannelID)
		modules.SendMessage(channelSession, message.ChannelID, msg, false)

	} else if strings.Contains(messageToUpper, options.Commands[9]) { //!RELOL  (remove)
		fmt.Println("listen")
		userN := strings.Replace(messageToUpper, "!RELOL ", "", -1)
		msg := db.UserRemove(userN)
		modules.SendMessage(channelSession, message.ChannelID, msg, false)
	}
}
