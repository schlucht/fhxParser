package fhx

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/schlucht/fhxreader/fhx/fhxReader"
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
var fhx = Fhx{}

/*
Einstiegspunkt zum laden der Daten aus einer Fhx Datei
*/
func NewFhxPath(path string) ([]Fhx, error) {
	fhx = Fhx{
		UnitName: "",
		Recipes:  []Recipe{},
		Units:    []Unit{},
		regFhx:   regFhx,
	}
	var fs = []Fhx{}
	err := fhxReader.IsFhxFile(path)

	if err != nil {
		return fs, err
	}

	fileText, err := fhxReader.ReadFhx(path)
	if err != nil {
		log.Fatal("FHX: Line 48, readFhx()", err)
	}
	fs = fhx.readFhx(fileText)
	return fs, nil
}

// Ein FHX String einlesen
func NewFhxString(fhxText string) error {
	fhx = Fhx{
		UnitName: "",
		Recipes:  []Recipe{},
		Units:    []Unit{},
		regFhx:   regFhx,
	}

	var fs = []Fhx{}
	if fhxText == "" {
		return errors.New("file is empty")
	}
	lines, err := fhxReader.ReadFhxText(fhxText)
	if err != nil {
		return err
	}
	fs = fhx.readFhx(lines)
	fhx.saveFhx(fs)
	return nil
}

/*
Save Object in a json File Structur
*/
func (m *Fhx) saveFhx(lines []Fhx) {

}

/*
Wird ausgeführt wenn ein FHX Datei eingelesen wird diese Eingelesen und neu gespeichert
*/

func (m *Fhx) readFhx(fileText []string) []Fhx {

	var fhxs = []Fhx{}

	block, err := fhxReader.ReadBlock("BATCH_RECIPE", fileText)
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range block {
		// Names of Unit
		if fhx.Unitname == "" {
			unitName := fhxReader.ReadParam(b, m.regFhx["Unit"])
			if len(unitName) > 0 {
				fhx.Unitname = unitName[0]
			}
		}
		// Type der Datei PROCEDURE OPERATION UNITPROCEDURE
		
		unitType := fhxReader.ReadParam(b, m.regFhx["Type"])
		if unitType[1] != "" {
			if unitType[1] == "UNIT_PROCEDURE" {
				unitType[1] = fhx.readUps(b)
			} else if unitType[1] == "PROCEDURE" {
				fhx.Step = fhx.readStep(b)				
			}
		}		}
		fhxs = append(fhxs, fhx)
	}
	return fhxs
}

func (m *Fhx) readStep(lines []string) []Step {
	var steps = []Step{}
	step := Step{}
	stepBlocks, err := fhxReader.ReadBlock("STEP", lines)
	if err != nil {
		log.Panicln("Read Step Block: ", err)
	}
	for _, b := range stepBlocks {
		for _, l := range b {
			name, err := fhxReader.ReadRegex(m.regFhx["Step"], l)
			if err != nil {
				log.Panic("ReadSteps Name: ", err)
			}
			if name != "" {
				step.Name = name
				// loadUP(name)
			}
			key, err := fhxReader.ReadRegex(m.regFhx["Definition"], l)
			if err != nil {
				log.Panic("ReadSteps Def: ", err)
			}
			if key != "" {
				step.Key = key
			}
			if step.Description == "" {
				desc, err := fhxReader.ReadRegex(m.regFhx["StepDesc"], l)

				if err != nil {
					log.Panic("ReadSteps Desc: ", err)
				}
				if desc != "" {
					step.Description = desc
				}
			}
			if step.Rect == "" {
				rect, err := fhxReader.ReadRegex(m.regFhx["Rect"], l)

				if err != nil {
					log.Panic("ReadSteps Rect: ", err)
				}
				if rect != "" {
					step.Description = rect
				}
			}
		}
		steps = append(steps, step)
		//
		// log.Println(stepparams)
		attrBlocks, err := fhxReader.ReadBlock("ATTRIBUTE_INSTANCE", b)
		if err != nil {
			log.Fatal("ReadStep():", err)
		}
		step.Parameters = m.stepParameters(attrBlocks)

		// log.Println(steps)
	}

	return steps
}
func (m *Fhx) stepParameters(attrBlock [][]string) []Parameter {
	parameters := []Parameter{}

	return parameters
}

/*
Liest die Unit Proceduren aus der FHX Datei
*/
func (m *Fhx) readUps(lines []string) []Procedure {
	var ups = []Procedure{}
	up := Procedure{}

	for _, l := range lines {
		if up.Description == "" {
			desc, err := fhxReader.ReadRegex(m.regFhx["Desc"], l)
			if err != nil {
				log.Fatal("ReadUps():", err)
			}
			if desc != "" {
				up.Description = desc
			}
		}
		if up.Name == "" {
			u, err := fhxReader.ReadRegex(m.regFhx["Recipe"], l)
			if err != nil {
				log.Fatal("ReadUps():", err)
			}
			if u != "" {
				up.Name = u
			}
		}
		if up.Time == 0 {
			time, err := fhxReader.ReadRegex(m.regFhx["Time"], l)
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
			author, err := fhxReader.ReadRegex(m.regFhx["Author"], l)
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
	up.Parameters = m.readParameters(paramBlocks, attrBlocks)

	ups = append(ups, up)
	return ups
}

/*
Auslesen der Parameter und hinzufügen ihrer Werte
*/
func (m *Fhx) readParameters(paramBlock [][]string, attrBlock [][]string) []Parameter {
	var params = []Parameter{}
	var param = Parameter{}
	for _, b := range paramBlock {
		for _, l := range b {
			name, err := fhxReader.ReadRegex(m.regFhx["Params"], l)
			if err != nil {
				log.Fatal("ReadParameters():", err)
			}
			if name != "" {
				param.Name = name
				param.Value = m.readAttribute(attrBlock, name)
			}
			if param.Description == "" {
				desc, err := fhxReader.ReadRegex(m.regFhx["Desc"], l)
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
func (m *Fhx) readAttribute(block [][]string, paramName string) Value {
	var val = Value{}
	for _, b := range block {
		for _, l := range b {
			// Wenn der Name vom Parameter gefunden wir
			if strings.Contains(l, paramName) {
				// Auslesen ob ein Set vorhanden ist
				if len(b) > 4 {
					u, err := fhxReader.ReadRegex(m.regFhx["ValueSet"], b[4])
					if err != nil {
						log.Fatal("ReadAttribute():", err)
					}
					val.Set = u
					strVal, err := fhxReader.ReadRegex(m.regFhx["ValueString"], b[5])
					if err != nil {
						log.Fatal("ReadAttribute():", err)
					}
					val.StringValue = strVal
					return val
				}

				// Nur bei Texten wie Melden
				//fmt.Println(len(b), i+2)

				parseLine := b[2]

				d, err := fhxReader.ReadRegex(m.regFhx["ValueDesc"], parseLine)
				if err != nil {
					log.Fatal("ReadAttribute():", err)
				}
				if d != "" {
					val.Cv = d
					return val
				}

				// Werte für Zahlen
				v, err := fhxReader.ReadRegexSubexp(m.regFhx["Value"], parseLine)
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
