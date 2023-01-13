package fhx

import (
	"log"
)

type Fhx struct {
	Name string	
	Time int32
	Type string
	Parameters []Parameter
}

func New(path string) string {
	rd, err := ReadUTF16(path)
	if err != nil {
		log.Println(err)
	}
	return rd
}
