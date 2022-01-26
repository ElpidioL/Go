package PostgreDB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetAllUsers() []MatchUser {
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

	////////////Reading mutiple values from Table
	// here we are using the struct defined on the start.
	var userList []MatchUser // going to use it to see all information

	rows, err := db.Query("SELECT name, id, discord_register FROM playerslol;")

	if err != nil {
		panic(err)
	}
	defer rows.Close() //Remember to close for Query

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var us MatchUser
		if err := rows.Scan(&us.Name, &us.Id, &us.Discord_register); err != nil {
			panic(err)
		}
		userList = append(userList, us)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return userList
}
