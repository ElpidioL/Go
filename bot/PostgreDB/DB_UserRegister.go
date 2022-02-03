package PostgreDB

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

///i Dont like the fact that i need to init the DB in every piece of script, but for now i'll keep that way.

func UserRegister(name string, id string, puuid string, accountid string, guildId, discordChannel string) string {
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

	sqlStatement := fmt.Sprintf("SELECT nome_base FROM player_base WHERE puuid_base='%s';", puuid)
	var playerName string

	row := db.QueryRow(sqlStatement)
	err = row.Scan(&playerName)
	switch err {
	case sql.ErrNoRows:
		sqlStatement := `
					INSERT INTO player_base (nome_base, id_base, puuid_base, account_id_base)
					VALUES ($1, $2, $3, $4)`
		_, err = db.Exec(sqlStatement, strings.ToUpper(name), id, puuid, accountid)

		if err != nil {
			{
			}
		}
	case nil:
		{
		}
	default:
		panic(err)
	}

	sqlStatement = fmt.Sprintf("SELECT name FROM playerslol WHERE puuid='%s' AND discord_register='%s';", puuid, guildId)

	row = db.QueryRow(sqlStatement)
	err = row.Scan(&playerName)
	switch err {
	case sql.ErrNoRows:
		sqlStatement := `
					INSERT INTO playerslol (name, id, puuid, accountid, discord_register, discord_text)
					VALUES ($1, $2, $3, $4, $5, $6)`
		_, err = db.Exec(sqlStatement, strings.ToUpper(name), id, puuid, accountid, guildId, discordChannel)

		if err != nil {
			return "Something went wrong, try to contact the bot creator"
			//panic(err)
		}
		return "Successfully registered " + name + " in our DB for this server"
	case nil:
		return playerName + " is already registered in our DB for this server"
	default:
		panic(err)
	}
}
