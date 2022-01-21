package options

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var BotShortcut string
var Key string
var LolKey string
var Commands = []string{} // you need yo increase here if you put more commands
var Guild string
var Player string
var PlayerIdLol string
var ChannelText string

func init() {
	//getting env Token to be able to use the bot
	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// init the shortcut to be used  "!" atm
	// setting the bot shortcut using env, cuz, why not?
	BotShortcut = os.Getenv("BOT_SHORTCUT")
	Key = os.Getenv("KEY_TOKEN")
	LolKey = os.Getenv("LOL_TOKEN")

	Guild = os.Getenv("GUILD_DISCORD")
	Player = os.Getenv("PLAYER_DISCORD")
	PlayerIdLol = os.Getenv("PLAYER_ID_LOL")
	ChannelText = os.Getenv("CHANNEL_TEXT")

	////If you want to change a command, its here.
	////command list, i wanted to asign all at same time, but couldn't find how, and if i do it out of the init, it doesnt get the prefix
	Commands = append(Commands, BotShortcut+"PING")
	Commands = append(Commands, BotShortcut+"PONG")
	Commands = append(Commands, BotShortcut+"HELP") /// all upper case plz
	Commands = append(Commands, BotShortcut+"TTS")
	Commands = append(Commands, BotShortcut+"JOIN")
	Commands = append(Commands, BotShortcut+"PLAY")
	Commands = append(Commands, BotShortcut+"LEAVE")
	Commands = append(Commands, BotShortcut+"LOOP")
	fmt.Println(Commands)
	/// the positions are important if you want to change the Function of this command
	//  otherwise, you just need to change
	//  i'll keep just numbers so it can be changed without affecting the Bot with different names
}
