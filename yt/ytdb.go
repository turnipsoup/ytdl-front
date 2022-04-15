package yt

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//Open the database, create if it does not exist, create history table
func openDatabase(file_loc string) {

	db, err := sql.Open("sqlite3", "./ytdl-front.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

}
