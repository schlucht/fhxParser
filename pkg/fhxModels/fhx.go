package fhxModels

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/schlucht/fhxreader/pkg/fhxReader"
)

var regFhx = map[string]string{
	"Author":      `AUTHOR="(?P<s>.*)"`,
	"Time":        `.*time=(?P<s>\d{10})/*`,
	"VERSION":     `VERSION="(?P<s>.*)"`,
	"Recipe":      `BATCH_RECIPE NAME="(?P<s>.*)" T`,
	"Params":      `FORMULA_PARAMETER NAME="(?P<s>.*)" T`,
	"Type":        `.*TYPE=(?P<s>.*) C`,
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

	for _, b := range block {
		// Names of Unit
		unitName := fhxReader.ReadParam(b, regFhx["Unit"])
		if len(unitName) > 0 {
			fhx.Unitname = unitName[0]
		}
		fhx.Procedures = fhx.ReadUps(b)
		fhxs = append(fhxs, fhx)
	}
	return fhxs
}

/*
Liest die Unit Proceduren aus der FHX Datei
*/
func (m *Fhx) ReadUps(lines []string) []Procedure {
	var ups = []Procedure{}
	up := Procedure{}
	var fhxType string
	// Names of UP
	for _, l := range lines {

		u, err := fhxReader.ReadRegex(regFhx["Recipe"], l)
		if err != nil {
			log.Fatal("ReadUps():", err)
		}
		if u != "" {
			up.Name = u
		}
		if fhxType == "" {
			fhxType, err = fhxReader.ReadRegex(regFhx["Type"], l)
		}
		if err != nil {
			log.Fatal("ReadUps():", err)
		}

		time, err := fhxReader.ReadRegex(regFhx["Time"], l)
		if err != nil {
			log.Fatal("ReadUps():", err)
		}

		if time != "" {
			t, err := strconv.Atoi(time)
			if err != nil {
				log.Fatal("ReadUps():", err)
			}
			up.Time = t
		}

		author, err := fhxReader.ReadRegex(regFhx["Author"], l)
		if err != nil {
			log.Fatal("ReadUps():", err)
		}
		if author != "" {
			up.Author = author
		}

	}
	if strings.Trim(fhxType, " ") == "PROCEDURE" {
		stepBlocks, err := fhxReader.ReadBlock("STEP NAME", lines)
		if err != nil {
			log.Fatal("ReadUps():", err)
		}
		steps := m.ReadSteps(stepBlocks)
		fmt.Println(steps)
	}

	paramBlocks, err := fhxReader.ReadBlock("FORMULA_PARAMETER", lines)
	if err != nil {
		log.Fatal("ReadUps():", err)
	}
	attrBlocks, err := fhxReader.ReadBlock("ATTRIBUTE_INSTANCE", lines)
	if err != nil {
		log.Fatal("ReadUps():", err)
	}
	up.Parameters = m.ReadParameters(paramBlocks, attrBlocks)

	ups = append(ups, up)
	return ups
}

/*
Liest die Steps aus einer Unitprozedure. Parameters ein Array mit Steps Block. Return ein Array von Steps
*/
func (m *Fhx) ReadSteps(paramBlock [][]string) []Step {
	var step = Step{}
	var steps = []Step{}

	steps = append(steps, step)

	return steps
}

/*
Auslesen der Parameter und hinzufügen ihrer Werte
*/
func (m *Fhx) ReadParameters(paramBlock [][]string, attrBlock [][]string) []Parameter {
	var params = []Parameter{}
	var param = Parameter{}
	for _, b := range paramBlock {
		for _, l := range b {
			name, err := fhxReader.ReadRegex(regFhx["Params"], l)
			if err != nil {
				log.Fatal("ReadParameters():", err)
			}
			if name != "" {
				param.Name = name
				param.Value = m.ReadAttribute(attrBlock, name)
			}

			desc, err := fhxReader.ReadRegex(regFhx["Desc"], l)
			if err != nil {
				log.Fatal("ReadParameters():", err)
			}
			if desc != "" {
				param.Description = desc
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
		for _, l := range b {
			// Wenn der Name vom Parameter gefunden wir
			if strings.Contains(l, paramName) {
				// Auslesen ob ein Set vorhanden ist
				if len(b) > 4 {
					u, err := fhxReader.ReadRegex(regFhx["ValueSet"], b[4])
					if err != nil {
						log.Fatal("ReadAttribute():", err)
					}
					val.Set = u
					strVal, err := fhxReader.ReadRegex(regFhx["ValueString"], b[5])
					if err != nil {
						log.Fatal("ReadAttribute():", err)
					}
					val.StringValue = strVal
					return val
				}

				// Nur bei Texten wie Melden
				//fmt.Println(len(b), i+2)

				parseLine := b[2]

				d, err := fhxReader.ReadRegex(regFhx["ValueDesc"], parseLine)
				if err != nil {
					log.Fatal("ReadAttribute():", err)
				}
				if d != "" {
					val.Cv = d
					return val
				}

				// Werte für Zahlen
				v, err := fhxReader.ReadRegexSubexp(regFhx["Value"], parseLine)
				if err != nil {
					log.Fatal("ReadAttribute():", err)
				}
				val.High = v["s1"]
				val.Low = v["s2"]
				val.Cv = v["s3"]
				val.Unit = v["s4"]
				return val
			}
		}
	}
	return val
}
