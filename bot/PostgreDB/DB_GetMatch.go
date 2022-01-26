package PostgreDB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetMatchDB(matchId int, alertChan chan bool) {
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
	sqlStatement := fmt.Sprintf(`SELECT match_id FROM matchidlist WHERE match_id='%d';`, matchId)

	row := db.QueryRow(sqlStatement)
	err = row.Scan(&mt)
	switch err {
	case sql.ErrNoRows:
		alertChan <- true
	case nil:
		alertChan <- false
	default:
		panic(err)
	}
	///End
}
