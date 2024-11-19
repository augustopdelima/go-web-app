package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

const DB_PATH = "./db/resume.db"
const SQL_DRIVER = "sqlite"

func InitDatabase() *sql.DB {
	databaseConnection, err := sql.Open(SQL_DRIVER, DB_PATH)
	if err != nil {
		log.Fatal(err)
	}

	_, err = databaseConnection.Exec(`CREATE TABLE IF NOT EXISTS resume(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		cellphone TEXT,
		email TEXT NOT NULL,
		webAddress TEXT,
		experience TEXT NOT NULL
	);`)

	if err != nil {
		log.Fatal(err)
	}

	return databaseConnection

}
