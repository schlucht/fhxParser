package models

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/schlucht/fhxreader/internal/helpers"
)

type System struct {
}

type DataBase struct {
	Comment string
	SQL     string
}

func (m *DBModel) GetSystemTable() (int, error) {
	tbls, err := m.DB.Query("SELECT COUNT(*) FROM duckdb_tables();")
	if err != nil {
		return 0, err
	}
	var cols int

	for tbls.Next() {
		if err = tbls.Scan(&cols); err != nil {
			return 0, err
		}
	}
	return cols, nil
}

func (m *DBModel) InstallDB() error {
	var data = []DataBase{}
	m.GetSystemTable()
	result, err := helpers.OpenJSON("assets/database/fhxdata.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(result, &data)
	if err != nil {
		return err
	}
	for _, d := range data {
		err = m.createTable(d.SQL)
		if err != nil {
			return fmt.Errorf("%s,%s", d.Comment, err)
		}
	}
	return nil
}

func (m *DBModel) createTable(sql string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := m.DB.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	return nil
}
