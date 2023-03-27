package fhxFactory

import (
	"fmt"

	fhx "github.com/schlucht/fhxreader/fhx/fhxModels"
)

const PATH = "./jsonDb"

type FhxFactory struct {
	File string
	Fhx  []fhx.Fhx
}

func Load(fileName string) *FhxFactory {
	ff := &FhxFactory{
		File: fileName,
		Fhx:  nil,
	}

	fs, err := fhx.NewFhxPath(fileName)
	if err != nil {
		fmt.Println(err)
	}
	ff.Fhx = fs
	return ff
}
