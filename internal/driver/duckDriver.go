package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/marcboeker/go-duckdb"
)

func DuckDBOpenDB(dataBaseFile string) (*sql.DB, error) {
	db, err := sql.Open("duckdb", dataBaseFile)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(db.Stats().OpenConnections)
	return db, nil
}
