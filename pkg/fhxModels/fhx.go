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
	"Step":        `STEP NAME="(?P<s>.*)" D`,
	"Definition":  `.*DEFINITION="(?P<s>.*)"`,
	"StepDesc":    `.*DESCRIPTION="(?P<s>.*)"`,
	"Rect":        `RECTANGLE= (?P<s>.*)`,
	"StepParams":  `STEP_PARAMETER NAME="(?P<s>.*)"`,
}

type Fhx struct {
	Unitname   string      `json:"unitname"`
	Type       string      `json:"type"`
	Procedures []Procedure `json:"unitprocedure,omitempty"`
	Step       []Step      `json:"steps,omitempty"`
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
		Unitname: "",
		Type:     "",
	}
	var fhxs = []Fhx{}

	fileText, err := fhxReader.ReadFhx(path)
	if err != nil {
		log.Fatal("FHX: Line 48, readFhx()", err)
	}

	block, err := fhxReader.ReadBlock("BATCH_RECIPE", fileText)
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range block {
		// Names of Unit
		if fhx.Unitname == "" {
			unitName := fhxReader.ReadParam(b, regFhx["Unit"])
			if len(unitName) > 0 {
				fhx.Unitname = unitName[0]
			}
		}
		// Type der Datei PROCEDURE OPERATION UNITPROCEDURE
		if fhx.Type == "" {
			unitType := fhxReader.ReadParam(b, regFhx["Type"])
			if len(unitType) > 0 {
				fhx.Type = unitType[0]
			}
		}
		if fhx.Type == "UNIT_PROCEDURE" {
			fhx.Procedures = fhx.ReadUps(b)
		} else if fhx.Type == "PROCEDURE" {
			fhx.Step = fhx.ReadStep(b)
			// log.Println(fhx.Step)
		}
		fhxs = append(fhxs, fhx)
	}
	return fhxs
}
func loadUP(upName string) {
	fmt.Println(upName)
}

func (m *Fhx) ReadStep(lines []string) []Step {
	var steps = []Step{}
	step := Step{}
	stepBlocks, err := fhxReader.ReadBlock("STEP", lines)
	if err != nil {
		log.Panicln("Read Step Block: ", err)
	}
	for _, b := range stepBlocks {
		for _, l := range b {
			name, err := fhxReader.ReadRegex(regFhx["Step"], l)
			if err != nil {
				log.Panic("ReadSteps Name: ", err)
			}
			if name != "" {
				step.Name = name
				// loadUP(name)
			}
			key, err := fhxReader.ReadRegex(regFhx["Definition"], l)
			if err != nil {
				log.Panic("ReadSteps Def: ", err)
			}
			if key != "" {
				step.Key = key
			}
			if step.Description == "" {
				desc, err := fhxReader.ReadRegex(regFhx["StepDesc"], l)

				if err != nil {
					log.Panic("ReadSteps Desc: ", err)
				}
				if desc != "" {
					step.Description = desc
				}
			}
			if step.Rect == "" {
				rect, err := fhxReader.ReadRegex(regFhx["Rect"], l)

				if err != nil {
					log.Panic("ReadSteps Rect: ", err)
				}
				if rect != "" {
					step.Description = rect
				}
			}
			// stepParam, err := fhxReader.ReadRegex(regFhx["StepParams"], l)

			// if err != nil {
			// 	log.Panic("ReadSteps Paramname: ", err)
			// }
			// if stepParam != "" {
			// 	p := Parameter{Name: stepParam}
			// 	step.Parameters = append(step.Parameters, p)
			// }
			// Step die Parameter duchlaufen in der ATTRIBUTE INSTANCE
			// Wenn der Name vorhanden ist, dann die Values auslesen
			// Wenn die Parameter nicht vorhanden sind dann Muss mit dem STEP Namen die Units durchsucht
			// Werden und dann die Values holen und anbinden.
		}
		steps = append(steps, step)
		//
		// log.Println(stepparams)
		attrBlocks, err := fhxReader.ReadBlock("ATTRIBUTE_INSTANCE", b)
		if err != nil {
			log.Fatal("ReadStep():", err)
		}
		step.Parameters = m.StepParameters(attrBlocks)

		// log.Println(steps)
	}

	return steps
}
func (m *Fhx) StepParameters(attrBlock [][]string) []Parameter {
	parameters := []Parameter{}

	return parameters
}

/*
Liest die Unit Proceduren aus der FHX Datei
*/
func (m *Fhx) ReadUps(lines []string) []Procedure {
	var ups = []Procedure{}
	up := Procedure{}

	for _, l := range lines {
		if up.Description == "" {
			desc, err := fhxReader.ReadRegex(regFhx["Desc"], l)
			if err != nil {
				log.Fatal("ReadUps():", err)
			}
			if desc != "" {
				up.Description = desc
			}
		}
		if up.Name == "" {
			u, err := fhxReader.ReadRegex(regFhx["Recipe"], l)
			if err != nil {
				log.Fatal("ReadUps():", err)
			}
			if u != "" {
				up.Name = u
			}
		}
		if up.Time == 0 {
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
		}

		if up.Author == "" {
			author, err := fhxReader.ReadRegex(regFhx["Author"], l)
			if err != nil {
				log.Fatal("ReadUps():", err)
			}
			if author != "" {
				up.Author = author
			}
		}

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
			if param.Description == "" {
				desc, err := fhxReader.ReadRegex(regFhx["Desc"], l)
				if err != nil {
					log.Fatal("ReadParameters():", err)
				}
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
