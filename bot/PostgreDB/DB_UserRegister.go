package PostgreDB

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

///i Dont like the fact that i need to init the DB in every piece of script, but for now i'll keep that way.

func UserRegister(name string, id string, puuid string, accountid string, discordChannel string) string {
	//storing the info to access the DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DbInfo.Host, DbInfo.Port, DbInfo.User, DbInfo.Password, DbInfo.Dbname)

	//starting db
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//don't forget to close it
	defer db.Close()
	//ping to check if its working
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	////////////Reading single info from Table
	//remember that we created this value on the update
	sqlStatement := fmt.Sprintf("SELECT name FROM playerslol WHERE puuid='%s';", puuid)
	var playerName string

	row := db.QueryRow(sqlStatement)
	err = row.Scan(&playerName)
	switch err {
	case sql.ErrNoRows:
		sqlStatement := `
					INSERT INTO playerslol (name, id, puuid, accountid, discord_register)
					VALUES ($1, $2, $3, $4, $5)`
		_, err = db.Exec(sqlStatement, strings.ToUpper(name), id, puuid, accountid, discordChannel)

		//simplified version db.Exec(sqlStatement, "10", "Cake", "black#9999")
		if err != nil {
			return "Something went wrong, try to contact the bot creator"
			//panic(err)
		}
		return "Successfully registered " + name + " in our DB"
	case nil:
		return playerName + " is already registered in our DB"
	default:
		panic(err)
	}
	///End
}
