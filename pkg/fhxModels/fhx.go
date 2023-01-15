package fhxModels

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/schlucht/fhxreader/pkg/fhxReader"
)

var regFhx = map[string]string{
	"Autor":       `user="(?P<s>.*)" t`,
	"Time":        `time=(?P<i>\d{10})/*`,
	"VERSION":     `VERSION="(?P<s>.*)"`,
	"Recipe":      `BATCH_RECIPE NAME="(?P<s>.*)" T`,
	"Params":      `FORMULA_PARAMETER NAME="(?P<s>.*)" T`,
	"TYPE":        `TYPE=(?P<s>.*) C`,
	"CATEGORY":    `CATEGORY="(?P<s>.*)"`,
	"Desc":        `DESCRIPTION="(?P<s>.*)"`,
	"Unit":        `EQUIPMENT_UNIT="(?P<s>.*)"`,
	"ValueSet":    `SET="(?P<s>.*)"`,
	"ValueString": `STRING_VALUE="(?P<s>.*)"`,
	"Value":       `.* HIGH=(?P<s1>.*).LOW=(?P<s2>.*).SC.*CV=(?P<s3>.*).UNITS="(?P<s4>.*)"`,
	"ValueDesc":   `CV="(?P<s>.*)"`,
}

type Fhx struct {
	Description string      `json:"desc"`
	Unitname    string      `json:"unitname"`
	Procedures  []Procedure `json:"unitprocedure"`
}

/*
Einstiegspunkt zum laden der Daten aus einer Fhx Datei
*/
func NewFhx(path string) []Fhx {
	ext := fhxReader.IsFhxFile(path)
	var f = Fhx{}
	var fs = []Fhx{}
	if ext == ".FHX" {
		fs = f.readFhx(path)
	}
	return fs
}

/*
Wird ausgeführt wenn ein FHX Datei eingelesen wird diese Eingelesen und neu gespeichert
*/
func (m *Fhx) readFhx(path string) []Fhx {
	var fhx = Fhx{
		Description: "Alle Standard Units vom der Unit",
		Unitname:    "",
	}
	var fhxs = []Fhx{}

	fileText, err := fhxReader.ReadFhxFile16(path)
	if err != nil {
		log.Fatal("FHX: Line 48, readFhx()", err)
	}

	block, err := fhxReader.ReadBlock("BATCH_RECIPE", fileText)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Repeat("%", 100))

	for _, b := range block {
		// Names of Unit
		unitName := fhxReader.ReadParam(b, regFhx["Unit"])
		if len(unitName) > 0 {
			m.Unitname = unitName[0]
		}
		fhx.ReadUps(b)
		fhxs = append(fhxs, fhx)
	}
	return fhxs
}
func (m *Fhx) readJson(path string) {
	// ToDo
}

/*
Liest die Unit Proceduren aus der FHX Datei
*/
func (m *Fhx) ReadUps(lines []string) {
	var ups = []Procedure{}
	up := Procedure{}
	// Names of UP
	upName := fhxReader.ReadParam(lines, regFhx["Recipe"])
	if len(upName) > 0 {
		up.Name = upName[0]
	}
	time := fhxReader.ReadParam(lines, regFhx["Time"])
	if len(time) > 0 {
		t, err := strconv.ParseInt(time[0], 2, 12)
		if err != nil {
			log.Fatal(err)
		}
		up.Time = t
	}
	author := fhxReader.ReadParam(lines, regFhx["Author"])
	if len(author) > 0 {
		up.Author = author[0]
	}
	paramBlocks, err := fhxReader.ReadBlock("FORMULA_PARAMETER", lines)
	if err != nil {
		log.Fatal((err))
	}
	attrBlocks, err := fhxReader.ReadBlock("ATTRIBUTE_INSTANCE", lines)
	if err != nil {
		log.Fatal((err))
	}
	up.Parameters = m.ReadParameters(paramBlocks, attrBlocks)

	ups = append(ups, up)
	m.Procedures = ups
}

/*
Auslesen der Parameter und hinzufügen ihrer Werte
*/
func (m *Fhx) ReadParameters(paramBlock [][]string, attrBlock [][]string) []Parameter {
	var params = []Parameter{}
	var param = Parameter{}
	for _, b := range paramBlock {
		for _, l := range b {
			u, _ := fhxReader.ReadRegex(regFhx["Params"], l)
			if u != "" {
				param.Name = u
				param.Value = m.ReadAttribute(attrBlock, u)
			}
			u, _ = fhxReader.ReadRegex(regFhx["Desc"], l)
			if u != "" {
				param.Description = u
			}
		}

		params = append(params, param)
	}
	return params
}

/*
Auslesen der Wert aus der ATTRIBUTE_INSTANCE die Werte werden den Parameter hinzugefügt.
*/
func (m *Fhx) ReadAttribute(block [][]string, paramName string) Value {
	var val = Value{}
	for _, b := range block {
		for i, l := range b {
			// Wenn der Name vom Parameter gefunden wir
			if strings.Contains(l, paramName) {
				// Auslesen ob ein Set vorhanden ist
				u, _ := fhxReader.ReadRegex(regFhx["ValueSet"], l)
				if u != "" {
					val.Set = u
					strVal, _ := fhxReader.ReadRegex(regFhx["ValueString"], b[i+1])
					val.StringValue = strVal
					return val // Beenden da Wert gefunden
				} else {
					// Nur bei Texten wie Melden
					d, _ := fhxReader.ReadRegex(regFhx["ValueDesc"], l)
					if d != "" {
						val.Cv = d
						return val // Beenden da Wert gefunden
					}
					// Werte für Zahlen
					v, _ := fhxReader.ReadRegexSubexp(regFhx["Value"], l)
					val.High = v["s1"]
					val.Low = v["s2"]
					val.Cv = v["s3"]
					val.Unit = v["s4"]
					return val // Beenden da Wert gefunden
				}
			}
		}
	}
	return val
}
