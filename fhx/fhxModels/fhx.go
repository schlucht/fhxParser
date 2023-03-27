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
		return nil, err
	}
	fs, err = fhx.readFhx(fileText)
	if err != nil {
		return nil, err
	}
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
	fs, err = fhx.readFhx(lines)
	if err != nil {
		return err
	}
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
func (m *Fhx) readFhx(fileText []string) ([]Fhx, error) {

	var fhxs = []Fhx{}

	block, err := fhxReader.ReadBlock("BATCH_RECIPE", fileText)
	if err != nil {
		return nil, err
	}

	for _, b := range block {
		var fhx Fhx
		// Type der Datei PROCEDURE OPERATION UNITPROCEDURE
		unitType := fhxReader.ReadParam(b, m.regFhx["Type"])
		if unitType[0] != "" {
			if unitType[0] == "UNIT_PROCEDURE" {
				units, err := m.readUnit(b)
				if err != nil {
					return nil, err
				}
				//log.Println("Anzahl:", units)
				fhx.Units = units
				fhxs = append(fhxs, fhx)
			} else if unitType[0] == "PROCEDURE" {
				recipes, err := m.readRecipe(b)
				if err != nil {
					return nil, err
				}
				fhx.Recipes = recipes
				fhxs = append(fhxs, fhx)
			}

		}
	}
	return fhxs, nil
}

func (m *Fhx) readUnit(fileText []string) ([]Unit, error) {

	var units = []Unit{}

	block, err := fhxReader.ReadBlock("BATCH_RECIPE", fileText)
	if err != nil {
		return nil, err
	}

	for _, b := range block {
		var unit Unit
		// Names of Unit
		if unit.UnitName == "" {
			unitName := fhxReader.ReadParam(b, m.regFhx["Recipe"])
			if len(unitName) > 0 {
				fhx.UnitName = unitName[0]
			}
		}
		if unit.UnitPosition == "" {
			unitPos := fhxReader.ReadParam(b, m.regFhx["Unit"])
			if len(unitPos) > 0 {
				unit.UnitPosition = unitPos[0]
			}
		}

		procedures, err := m.readUps(b)
		if err != nil {
			return nil, err
		}
		unit.Procedures = procedures
		units = append(units, unit)
	}
	return units, nil
}

func (m *Fhx) readRecipe(fileText []string) ([]Recipe, error) {

	var recipes = []Recipe{}

	block, err := fhxReader.ReadBlock("BATCH_RECIPE", fileText)
	if err != nil {
		return nil, err
	}

	for _, b := range block {
		var recipe Recipe
		// Names of Unit
		if recipe.RecipeName == "" {
			recipeName := fhxReader.ReadParam(b, m.regFhx["Recipe"])
			if len(recipeName) > 0 {
				recipe.RecipeName = recipeName[0]
			}
		}
		steps, err := m.readStep(b)
		if err != nil {
			return nil, err
		}
		recipe.Steps = steps
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func (m *Fhx) readStep(lines []string) ([]Step, error) {
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
				return nil, err
			}
			if name != "" {
				step.Name = name
				// loadUP(name)
			}
			key, err := fhxReader.ReadRegex(m.regFhx["Definition"], l)
			if err != nil {
				return nil, err
			}
			if key != "" {
				step.Key = key
			}
			if step.Description == "" {
				desc, err := fhxReader.ReadRegex(m.regFhx["StepDesc"], l)

				if err != nil {
					return nil, err
				}
				if desc != "" {
					step.Description = desc
				}
			}
			if step.Rect == "" {
				rect, err := fhxReader.ReadRegex(m.regFhx["Rect"], l)

				if err != nil {
					return nil, err
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
			return nil, err
		}
		params, err := m.stepParameters(attrBlocks)
		if err != nil {
			return nil, err
		}
		step.Parameters = params

		// log.Println(steps)
	}

	return steps, err
}

func (m *Fhx) stepParameters(attrBlock [][]string) ([]Parameter, error) {
	parameters := []Parameter{}

	return parameters, nil
}

/*
Liest die Unit Proceduren aus der FHX Datei
*/
func (m *Fhx) readUps(lines []string) ([]Procedure, error) {
	var ups = []Procedure{}
	up := Procedure{}

	for _, l := range lines {
		if up.Description == "" {
			desc, err := fhxReader.ReadRegex(m.regFhx["Desc"], l)
			if err != nil {
				return nil, err
			}
			if desc != "" {
				up.Description = desc
			}
		}
		if up.Name == "" {
			u, err := fhxReader.ReadRegex(m.regFhx["Recipe"], l)
			if err != nil {
				return nil, err
			}
			if u != "" {
				up.Name = u
			}
		}
		if up.Time == 0 {
			time, err := fhxReader.ReadRegex(m.regFhx["Time"], l)
			if err != nil {
				return nil, err
			}

			if time != "" {
				t, err := strconv.Atoi(time)
				if err != nil {
					return nil, err
				}
				up.Time = t
			}
		}

		if up.Author == "" {
			author, err := fhxReader.ReadRegex(m.regFhx["Author"], l)
			if err != nil {
				return nil, err
			}
			if author != "" {
				up.Author = author
			}
		}
	}
	paramBlocks, err := fhxReader.ReadBlock("FORMULA_PARAMETER", lines)
	if err != nil {
		return nil, err
	}
	attrBlocks, err := fhxReader.ReadBlock("ATTRIBUTE_INSTANCE", lines)
	if err != nil {
		return nil, err
	}
	params, err := m.readParameters(paramBlocks, attrBlocks)
	if err != nil {
		return nil, err
	}
	up.Parameters = params

	ups = append(ups, up)
	return ups, nil
}

/*
Auslesen der Parameter und hinzufügen ihrer Werte
*/
func (m *Fhx) readParameters(paramBlock [][]string, attrBlock [][]string) ([]Parameter, error) {
	var params = []Parameter{}
	var param = Parameter{}
	for _, b := range paramBlock {
		for _, l := range b {
			name, err := fhxReader.ReadRegex(m.regFhx["Params"], l)
			if err != nil {
				return nil, err
			}
			if name != "" {
				param.Name = name
				val, err := m.readAttribute(attrBlock, name)
				if err != nil {
					return nil, err
				}
				param.Value = val
			}
			if param.Description == "" {
				desc, err := fhxReader.ReadRegex(m.regFhx["Desc"], l)
				if err != nil {
					return nil, err
				}
				param.Description = desc
			}
		}
		params = append(params, param)
	}
	return params, nil
}

/*
Auslesen der Wert aus der ATTRIBUTE_INSTANCE die Werte werden den Parameter hinzugefügt.
*/
func (m *Fhx) readAttribute(block [][]string, paramName string) (Value, error) {
	var val Value
	for _, b := range block {
		for _, l := range b {
			// Wenn der Name vom Parameter gefunden wir
			if strings.Contains(l, paramName) {
				// Auslesen ob ein Set vorhanden ist
				if len(b) > 4 {
					u, err := fhxReader.ReadRegex(m.regFhx["ValueSet"], b[4])
					if err != nil {
						return val, err
					}
					val.Set = u
					strVal, err := fhxReader.ReadRegex(m.regFhx["ValueString"], b[5])
					if err != nil {
						return val, err
					}
					val.StringValue = strVal
					return val, nil
				}

				// Nur bei Texten wie Melden
				//fmt.Println(len(b), i+2)

				parseLine := b[2]

				d, err := fhxReader.ReadRegex(m.regFhx["ValueDesc"], parseLine)
				if err != nil {
					return val, err
				}
				if d != "" {
					val.Cv = d
					return val, nil
				}

				// Werte für Zahlen
				v, err := fhxReader.ReadRegexSubexp(m.regFhx["Value"], parseLine)
				if err != nil {
					return val, err
				}
				val.High = v["s1"]
				val.Low = v["s2"]
				val.Cv = v["s3"]
				val.Unit = v["s4"]
				return val, nil
			}
		}
	}
	return val, nil
}
