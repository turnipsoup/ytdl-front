package yt

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//Open the database, create if it does not exist, create history table
func OpenDatabase(file_loc string) {

	log.Printf("Opening the database at %s\n", file_loc)

	db, err := sql.Open("sqlite3", file_loc)
	if err != nil {
		log.Println("Could not create or open database")
		log.Fatal(err)
	}

	defer db.Close()

	sqlStmt := `
	create table ytdl (id text not null primary key, start_date int, end_date int, url text, status text);
	commit;
	`
	_, err = db.Exec(sqlStmt)

	if err != nil {
		log.Printf("%s\n", err)
		return
	}

}
