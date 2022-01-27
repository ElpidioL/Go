package PostgreDB

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

///i Dont like the fact that i need to init the DB in every piece of script, but for now i'll keep that way.

func MatchRegister(matchId int, playerName string, championName string, discord_id string) bool {
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
	sqlStatement := fmt.Sprintf(`SELECT match_id FROM matchidlist WHERE match_id='%d' AND player_name='%s';`, matchId, strings.ToUpper(playerName))
	var matchInfo string

	row := db.QueryRow(sqlStatement)
	err = row.Scan(&matchInfo)
	switch err {
	case sql.ErrNoRows:
		sqlStatement := `
					INSERT INTO matchidlist (match_id, player_name, player_champion,discord_id)
					VALUES ($1, $2, $3, $4)`
		_, err = db.Exec(sqlStatement, matchId, strings.ToUpper(playerName), championName, discord_id)
		if err != nil {
			fmt.Println(err)
			return false
			//panic(err)
		}
		return true
	case nil:
		return false
	default:
		panic(err)
	}
	///End
}
