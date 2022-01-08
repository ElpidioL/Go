package options

import (
	"github.com/joho/godotenv"
	"os"
	"fmt"
)

var BotShortcut string  /// below on init you can change the prefix of the bot
var Key string
var Commands [3] string // you need yo increase here if you put more commands

func init(){
	//getting env Token to be able to use the bot
	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("Error loading .env file")  // if it return nil, then its ok, if not, break plz.
		return
	}
	//
	// init the shortcut to be used  "!" atm
	BotShortcut = os.Getenv("BOT_SHORTCUT")  // setting the bot shortcut using env, cuz, why not?
	Key = os.Getenv("KEY_TOKEN")

	////If you want to change a command, its here.
	////command list, i wanted to asign all at same time, but couldn't find how, and if i do it out of the init, it doesnt get the prefix
 	Commands[0]	=	BotShortcut+"PING" 
	Commands[1]	=	BotShortcut+"PONG"
	Commands[2] =	BotShortcut+"HELP"   /// all upper case plz
	
//////////////////////////////// the positions are important if you want to change the Function of this command
//                               otherwise, you just need to change 
//                               i'll keep just numbers so it can be changed without affecting the Bot with different names
}



