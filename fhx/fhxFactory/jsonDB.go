package fhxFactory

import (
	"fmt"
	"path"

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
	fmt.Println(path.Join(PATH, fileName+".json"))
	return ff
}
