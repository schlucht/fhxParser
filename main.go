package main

import (
	"fmt"

	"github.com/schlucht/fhxreader/pkg/fhx"
)

func main() {
	p := "./files/deltaV.fhx"
	fhx.New(p)
	fmt.Println("Hallo Lothar", p)
}
