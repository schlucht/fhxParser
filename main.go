package main

import (
	"log"

	"github.com/schlucht/fhxreader/pkg/database"
	"github.com/schlucht/fhxreader/pkg/fhxModels"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
	p := "./files/UP_Q2800_Up.fhx"
	f := fhxModels.NewFhx(p)
	database.WriteFhx(f, "units")

}
