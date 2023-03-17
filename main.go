package main

import (
	_ "github.com/schlucht/fhxreader/pkg/database"
	_ "github.com/schlucht/fhxreader/pkg/fhxModels"
	"github.com/schlucht/fhxreader/server"
)

func main() {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)

	// p := "./files/qall.fhx"
	// f := fhxModels.NewFhx(p)

	// database.WriteFhx(f, "units")

	server.Start()

}
