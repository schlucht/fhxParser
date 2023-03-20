package fhxfactory

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/schlucht/fhxreader/pkg/fhxModels"
)

func WriteFhx(obj []fhxModels.Fhx, pathName string) {
	b, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	p := filepath.Join("./", "database", pathName+".json")
	f, err := os.Create(p)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write(b)
}

func LoadAllStandardFilename() (map[string][]string, error) {
	return make(map[string][]string), nil
}
