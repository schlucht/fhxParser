package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/schlucht/fhxreader/fhx-parser/models"
)

type DB struct {
	connect *sql.DB
}

var conn *sql.DB

const dbTimeout = time.Second * 3

func New() DB {
	conn, err := dBConnect()
	if err != nil {
		log.Fatal(err)
	}
	return DB{
		connect: conn,
	}
}

func dBConnect() (*sql.DB, error) {
	name := "schmidschluch4"
	host := "db8.hostpark.net"
	pw := "Schlucht6"
	connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		name, pw, host, name,
	)

	_db, err := sql.Open("mysql", connectString)

	if err != nil {
		return nil, err
	}
	return _db, nil
}

// Speichert die Daten in der Tabelle Unit
func (db *DB) InsertUnit(unit models.Unit) error {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	defer db.connect.Close()
	// Daten in die Tabelle einfügen
	stmt, err := db.connect.PrepareContext(ctx, `INSERT INTO units 
		(unit_name, unit_position, time, author, description) 
		VALUES 
		(?, ?, ?, ?, ?);
		`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		&unit.UnitName,
		&unit.UnitPosition,
		&unit.Time, &unit.Author,
		&unit.Description,
	)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) InsertUnits(units []models.Unit) error {

	return nil
}

// Id der UNIT aus der Tabelle holen
func (db *DB) IdFromUnitname(unitname string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	defer db.connect.Close()

	stmt, err := db.connect.PrepareContext(ctx, `
		SELECT unit_id FROM units
		WHERE unit_name=?
	`)
	if err != nil {
		return 0, err
	}
	var id int
	err = stmt.QueryRowContext(ctx, unitname).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		} else {
			return 0, err
		}
	}
	return id, nil
}

// Unitdatensatz in der Tabelle aktualisieren
func (db *DB) UpdateUnit(units []models.Unit) error {

	// Datensatz suchen

	// Felder mit Parameter vergleichen

	// UPDATE Statment an DB absenden

	return nil
}

func (db *DB) DeleteUnit(id int) error {

	// Datensatz suchen

	// DELETE Statmemt absetzen

	return nil
}
