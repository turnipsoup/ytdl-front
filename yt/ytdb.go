package yt

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//Open the database, create if it does not exist, create history table
func OpenDatabaseInit(file_loc string) {

	log.Printf("Opening the database at %s\n", file_loc)

	db, err := sql.Open("sqlite3", file_loc)
	if err != nil {
		log.Println("Could not create or open database")
		log.Fatal(err)
	}

	defer db.Close()

	sqlStmt := `
	create table ytdl (id text not null primary key, start_date int, end_date int, url text, genre text, status text);
	`

	log.Printf("Creating table: %s\n", sqlStmt)

	_, err = db.Exec(sqlStmt)

	if err != nil {
		log.Printf("%s\n", err)
		return
	}

}

// Returns the DB object to be used later. Do *not* forget to defer db.Close()!!
func OpenDatabase(file_loc string) *sql.DB {
	log.Printf("Opening the database at %s\n", file_loc)

	db, err := sql.Open("sqlite3", file_loc)
	if err != nil {
		log.Println("Could not create or open database")
		log.Fatal(err)
	}

	return db
}

// Inserts a record into the ytdl table
func InsertYTDLRecord(dbLoc string, ytId string, startTime int, endTime int, url string, genre string, status string) {
	db := OpenDatabase(dbLoc)

	tx, err := db.Begin()

	if err != nil {
		log.Println(err)
	}

	stmt, err := tx.Prepare("insert into ytdl(id, start_date, end_date, url, genre, status) values(?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(ytId, 0, 0, CreateYTUrl(ytId), genre, "Active")
	if err != nil {
		log.Println(err)
	}

	tx.Commit()
}
