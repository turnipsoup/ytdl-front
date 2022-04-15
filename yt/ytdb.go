package yt

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type ytRow struct {
	Id        string
	Status    string
	StartTime int
	EndTime   int
	Url       string
	Genre     string
}

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
	defer db.Close()

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

// Marks a record as 'Done' in the DB given an ID.
func MarkDownloadDone(dbLoc string, ytId string) {
	db := OpenDatabase(dbLoc)
	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		log.Println(err)
	}

	stmt, err := tx.Prepare(fmt.Sprintf("UPDATE ytdl SET status = 'Done' WHERE id = '%s'", ytId))

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		log.Println(err)
	}

	tx.Commit()

}

// Returns all rows from theytdl tabls as an array of ytRows
func GetAllDownloads(dbLoc string) []ytRow {

	var ytRows []ytRow

	db := OpenDatabase(dbLoc)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM ytdl;")

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var id string
		var status string
		var startTime int
		var endTime int
		var url string
		var genre string
		err = rows.Scan(&id, &startTime, &endTime, &url, &genre, &status)
		if err != nil {
			log.Println(err)
		}
		newRow := ytRow{id, status, startTime, endTime, url, genre}
		ytRows = append(ytRows, newRow)
	}

	return ytRows

}

// Convers passed row into a JSON string
func RowToJSON(row ytRow) string {
	data, err := json.Marshal(row)

	if err != nil {
		log.Println("Error unpacking row")
		log.Println(row)
	}

	return string(data)
}
