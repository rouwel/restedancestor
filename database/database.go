// Package database takes care of read / write process.
package database

import (
	"database/sql"
	_ "modernc.org/sqlite"
	

)

var DB *sql.DB

func Database(filepath string) error {
	var err error
	DB , err = sql.Open("sqlite", filepath)
	if err != nil {
		return err
	}
	return DB.Ping()
}


// Database describes accesses to a storage
//type Database interface {
//	Prepare(string, ...interface{}) (Stmt, error)
//	LastInsertRowID() int64
//}

// Db represents a connection to the storage
//type Db struct {
//	sqlite *sqlite3.Conn
//}

// Prepare encapsulates the inner connection for testability
//func (d Db) Prepare(sql string, args ...interface{}) (Stmt, error) {
//	return d.sqlite.Prepare(sql, args...)
//}

// LastInsertRowID encapsulates the inner return
//func (d Db) LastInsertRowID() int64 {
//	return d.sqlite.LastInsertRowID()
//}

// Stmt represents a query statement
// Cf. sqlite3.Stmt
//type Stmt interface {
//	Close() error
//	Step() (bool, error)
//	Exec(...interface{}) error
//	Scan(dst ...interface{}) error
//}

// NewDb initialise a new connection
//func NewDb() Database {
//	p := "/src/github.com/bruno-chavez/restedancestor/database/database.db"
//	s, err := sqlite3.Open(os.Getenv("GOPATH") + p)
//	if err != nil {
//		log.Fatal(err)
//	}
//	s.BusyTimeout(5 * time.Second)

//	return Db{
//		sqlite: s,
//	}
//}
