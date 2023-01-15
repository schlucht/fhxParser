package main

import (
	"github.com/schlucht/fhxreader/pkg/database"
	"github.com/schlucht/fhxreader/pkg/fhxModels"
)

func main() {
	p := "./files/UP_Q2800_Up.fhx"
	f := fhxModels.NewFhx(p)
	database.WriteFhx(f, "units")

}
