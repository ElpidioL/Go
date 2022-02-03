package PostgreDB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetAllDiscords(puuid string) []DiscordList {
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
	var discordList []DiscordList // going to use it to see all information

	rows, err := db.Query("SELECT name, discord_register, discord_text FROM playerslol WHERE puuid='%s';", puuid)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close() //Remember to close for Query

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var disc DiscordList
		if err := rows.Scan(&disc.Name, &disc.Discords); err != nil {
			panic(err)
		}
		discordList = append(discordList, disc)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return discordList
}
