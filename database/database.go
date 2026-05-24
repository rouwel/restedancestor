// Package database takes care of read / write process.
package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Database(filepath string) error {
	var err error
	DB, err = sql.Open("sqlite", filepath)
	if err != nil {
		return err
	}
	return DB.Ping()
}
