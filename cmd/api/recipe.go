package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/schlucht/fhxreader/internal/models"
	_ "github.com/schlucht/fhxreader/internal/parser"
)

type Recipe struct {
	Name string
	Author string
	Description string
	
}