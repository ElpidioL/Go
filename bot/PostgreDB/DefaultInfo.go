package PostgreDB

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

//This struc holds the Postgree acess format
type DbAcess struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

var DbInfo DbAcess

type MatchUser struct {
	Name  string
	Id    string
	Puuid string
}
type DiscordList struct {
	Name     string
	Discords string
	Discords_text string
}

func init() {
	//getting env values to hide my sensitive info
	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("Error loading .env file")
		return
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		//error if the provided PORT in .env its not an INT  "10"
		fmt.Println("Port error, verify if its a number", err)
		os.Exit(0)
	}
	DbInfo.Host = os.Getenv("HOST")
	DbInfo.Port = port
	DbInfo.User = os.Getenv("USER")
	DbInfo.Password = os.Getenv("PASSWORD")
	DbInfo.Dbname = os.Getenv("DB_NAME")
}
