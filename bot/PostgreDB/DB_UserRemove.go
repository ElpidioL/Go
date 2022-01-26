package PostgreDB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

///i Dont like the fact that i need to init the DB in every piece of script, but for now i'll keep that way.

func UserRemove(userName string) string {
	fmt.Println("db in")
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

	////////////Deleting info to Table
	sqlStatement := `
					DELETE FROM playerslol 
					WHERE name = $1;`
	_, err = db.Exec(sqlStatement, userName)
	fmt.Println(userName)
	if err != nil {
		fmt.Println("ta de boa")
		return fmt.Sprintf("Sorry but %s does not exist in our DB", userName)
	}
	return fmt.Sprintf("%s successfully deleted from our DB", userName)
}