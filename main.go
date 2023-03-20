package main

import (
	"log"

	"github.com/schlucht/fhxreader/pkg/fhxFactory"
	"github.com/schlucht/fhxreader/pkg/fhxModels"
	_ "github.com/schlucht/fhxreader/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	p := "./files/q2000.fhx"
	f := fhxModels.NewFhx(p)

	fhxFactory.WriteFhx(f, "units")

	// server.Start()

}
