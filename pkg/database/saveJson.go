package database

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func WriteFhx(obj interface{}, pathName string) {
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
