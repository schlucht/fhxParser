package fhxFactory

import (
	fhx "github.com/schlucht/fhxreader/fhx/fhxModels"
)

const PATH = "./jsonDb"

type FhxFactory struct {
	File string
	Fhx  []fhx.Fhx
}

func Load(fileName string) (*FhxFactory, error) {
	ff := &FhxFactory{
		File: fileName,
		Fhx:  nil,
	}
	fs, err := fhx.NewFhxPath(fileName)
	if err != nil {
		return ff, err
	}

	ff.Fhx = fs

	return ff, nil
}
