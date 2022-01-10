package functions

import (
	"bots/GOing/options"

	"github.com/bwmarrin/discordgo"

	//"fmt"
	"strings"
)

func PingPongFunc(channelSession *discordgo.Session, message *discordgo.MessageCreate) {
	// If the message is "ping" reply with "Pong!"
	if strings.ToUpper(message.Content) == options.Commands[0] { //bothers me to check the same thing on two palces
		channelSession.ChannelMessageSend(message.ChannelID, options.Commands[1]) //but the less i hard code the better it is to
	} //change this function later

	// If the message is "pong" reply with "Ping!"
	if strings.ToUpper(message.Content) == options.Commands[1] {
		channelSession.ChannelMessageSend(message.ChannelID, options.Commands[0])
	}
	if message.Content == "Ping!" {
		channelSession.ChannelMessageDelete(message.ChannelID, message.ID)
	}
}

/* func (s *Session) ChannelMessageSendReply(channelID string, content string, reference *MessageReference) (*Message, error) {
	if reference == nil {
		return nil, fmt.Errorf("reply attempted with nil message reference")
	}
	return s.ChannelMessageSendComplex(channelID, &MessageSend{
		Content:   content,
		Reference: reference,
	})
} */

/* func (s *Session) ChannelMessageEdit(channelID, messageID, content string) (*Message, error) {
	return s.ChannelMessageEditComplex(NewMessageEdit(channelID, messageID).SetContent(content))
} */

// ChannelMessageDelete deletes a message from the Channel.
