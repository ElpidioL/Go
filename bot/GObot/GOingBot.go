package GObot

import (
	functions "bots/GOing/commandsFunc"
	"bots/GOing/modules"
	"bots/GOing/options"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	Notify = true
	//here i'll be handling the Lol API for now, since i need to update with a different rate than discord handler
	//this will be here.
	for true {
		//sleep so i don't waste much process power for nothing (probably there is a better way to do that...)
		time.Sleep(10 * time.Second)
		//get info from the API
		summonerName, gameMode, champion := modules.GetMatchLol(options.PlayerIdLol)
		//check if the player is in game and set the notify so it don't spamm
		if summonerName != "" && Notify == true {
			Notify = false
			message := "O Crime foi iniciado, preparem seus ouvidos e seus chats porque " + summonerName + " começou a gameplay criminosa jogando de " + champion + " em uma partida " + gameMode + " se preparem para o choro infinito do grande menino rafael"
			//in this case, i need to register all guilds that want this feature, since i don't know who sent the message
			//but i think i'll implement a DB later with player name to track, channel to post, and guild to post and player discord id
			//not sure how i would track the notify for everyone, but this is a joke feature anyway
			functions.PlayHorn(Discord, options.Guild, modules.FindVoiceChannel(Discord, options.Guild, options.Player))
			modules.SendMessage(Discord, options.ChannelText, message, true)

		} else if len(summonerName) == 0 && Notify != true {

			Notify = true
			time.Sleep(300 * time.Second)
		}
	}

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
