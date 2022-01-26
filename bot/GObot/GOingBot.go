package GObot

import (
	"bots/GOing/API"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	//encoding/json"

	"github.com/bwmarrin/discordgo"
)

var Notify bool

func Start(Discord *discordgo.Session, BotID string) {
	///////////////////////////////////////////////////setting Discord//////////////////////////
	//start discord or return an error if it fails
	err := Discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	Discord.AddHandler(ReceiveMessage)
	//Discord.AddHandler(HandleReactions)
	//Discord.Identify.Intents = discordgo.IntentsAll
	Discord.Identify.Intents = discordgo.IntentsGuildMessages // not sure what it does actually, i think it set the "Intent"

	///////////////////////////////////////////////////End////////////////////////////////////////

	//
	go API.NotifyLol(Discord)
	//
	////////////////////////////////////////////////Keeping the application alive////////////////////////
	// to only get messages
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill) // it will keep track of the console to stop the application
	<-sc                                                                      // until someone press ctr-c to stop it.
	// no line of code will be executed after it, only when it stops.
	// Cleanly close down the Discord session.
	Discord.Close()
	///////////////////////////////////////////////////End////////////////////////////////////////
}
