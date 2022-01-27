package PostgreDB

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

func GetMatchDB(matchId int, playerName string) bool {
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
	var mt string
	////////////Reading single info from Table
	//remember that we created this value on the update
	/* sqlStatement := fmt.Sprintf(`SELECT match_id FROM matchidlist WHERE match_id='%d';`, matchId) */ //2448439398

	sqlStatement := fmt.Sprintf(`SELECT match_id FROM matchidlist WHERE match_id='%d' AND player_name='%s';`, matchId, strings.ToUpper(playerName))
	row := db.QueryRow(sqlStatement)
	err = row.Scan(&mt)

	switch err {
	case sql.ErrNoRows:
		return true
	case nil:
		return false
	default:
		panic(err)
	}
	///End
}
