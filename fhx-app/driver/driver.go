package driver

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DB holt the connections pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDBCOnn = 10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		return nil, err
	}

	d.SetMaxOpenConns(maxOpenDBCOnn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	dbConn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return nil
	}
	return nil
}

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = testDB(db); err != nil {
		return nil, err
	}
	return db, nil
}
