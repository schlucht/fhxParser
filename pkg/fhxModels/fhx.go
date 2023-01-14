package fhxModels

import (
	"fmt"
	"log"

	"github.com/schlucht/fhxreader/pkg/fhxReader"
)

var regFhx = map[string]string{
	"user":     `user="(?P<s>.*)" t`,
	"time":     `time=(?P<i>\d{10})/*`,
	"VERSION":  `VERSION="(?P<s>.*)"`,
	"Recipe":   `BATCH_RECIPE NAME="(?P<s>.*)" T`,
	"TYPE":     `TYPE=(?P<s>.*) C`,
	"CATEGORY": `CATEGORY="(?P<s>.*)"`,
	"Desc":     `DESCRIPTION="(?P<s>.*)"`,
	"Unit":     `EQUIPMENT_UNIT="(?P<s>.*)"`,
}

type Fhx struct {
	Unit Unit `json:"unit"`
}

func New(path string) Fhx {

	fileText, err := fhxReader.ReadFhxFile16(path)
	var fhx Fhx
	var unit Unit
	var ups []Up = []Up{}
	if err != nil {
		log.Println(err)
	}
	params := readParam(fileText, "Unit")
	upsText := readParam(fileText, "Recipe")
	for _, u := range upsText {
		ups = append(ups, Up{Name: u})
	}
	unit = Unit{Name: params[0], Ups: ups}

	fhx = Fhx{Unit: unit}
	return fhx
}

func readParam(lines []string, regexName string) []string {
	var results []string
	regex := regFhx[regexName]
	fmt.Println(regex)
	for _, l := range lines {
		u, _ := fhxReader.ReadRegex(regex, l)
		// fmt.Println(u)
		if u != "" {
			results = append(results, u)
		}
	}
	return results
}
