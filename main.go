package main

import (
	"fmt"

	"github.com/schlucht/fhxreader/fhx-app/server"
)

func main() {
	srv := server.Server{}
	fmt.Println("Hallo Lothar")
	srv.Start()

}
