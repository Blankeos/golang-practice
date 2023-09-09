// This is the package-level database that
// Exports the database connection to be used
// in multiple files.

package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// A regular Database connection.
	Conn *sql.DB

	// An sqlc wrapper around the database connection (For typesafe queries).
	DB *Queries
)

func Init(dataSource string) error {
	conn, err := sql.Open("sqlite3", dataSource)

	if err != nil {
		return err
	}

	Conn = conn
	DB = New(conn)
	return nil
}

func Close() {
	if Conn != nil {
		Conn.Close()
	}
}