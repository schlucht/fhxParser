package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Fhx struct {
	Recipes []Recipe          `json:"recipes,omitempty"`
	Units   []Unit            `json:"units,omitempty"`
	OPs     []Unit            `json:"ops,omitempty"`
	regFhx  map[string]string `json:"-"`
}

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
		Recipes: []Recipe{},
		Units:   []Unit{},
		OPs:     []Unit{},
		regFhx:  regFhx,
	}
	var fs = []Fhx{}
	err := IsFhxFile(path)

	if err != nil {
		return fs, err
	}

	fileText, err := ReadFhx(path)
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
func NewFhxString(fhxText string) ([]Fhx, error) {
	fhx = Fhx{
		Recipes: []Recipe{},
		Units:   []Unit{},
		OPs:     []Unit{},
		regFhx:  regFhx,
	}
	if fhxText == "" {
		return nil, errors.New("NewFHXString, no file i")
	}
	lines, err := ReadFhxText(fhxText, "")
	if err != nil {
		return nil, err
	}

	fhx, err := fhx.readFhx(lines)
	if err != nil {
		return nil, err
	}

	return fhx, nil
}

/*
Wird ausgeführt wenn ein FHX Datei eingelesen wird diese Eingelesen und neu gespeichert
*/
func (m *Fhx) readFhx(fileText []string) ([]Fhx, error) {
	// log.Println(fileText)
	var fhxs = []Fhx{}
	block, err := ReadBlock("BATCH_RECIPE", fileText)
	// fmt.Printf("Erstes File: %v\n", block[0])
	if err != nil {
		return nil, err
	}

	for _, b := range block {
		var fhx Fhx
		// Type der Datei PROCEDURE OPERATION UNITPROCEDURE
		unitType, err := ReadParam(b, m.regFhx["Type"])
		if err != nil {
			return nil, err
		}

		if unitType[0] != "" {
			if unitType[0] == "UNIT_PROCEDURE" {
				units, err := m.readUnit(b, "UP")
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
				// Noch nicht getestet
			} else if unitType[0] == "OPERATION" {
				ops, err := m.readUnit(b, "OP")
				if err != nil {
					return nil, err
				}

				fhx.OPs = ops
				fhxs = append(fhxs, fhx)
			}
		}
	}
	return fhxs, nil
}

func (m *Fhx) readUnit(fileText []string, unit_type string) ([]Unit, error) {

	var units = []Unit{}

	block, err := ReadBlock("BATCH_RECIPE", fileText)
	if err != nil {
		return nil, err
	}

	for _, b := range block {
		var unit Unit
		// Names of Unit
		unit.Type = unit_type
		if unit.UnitName == "" {
			unitName, err := ReadParam(b, m.regFhx["Recipe"])
			if err != nil {
				return nil, err
			}
			if len(unitName) > 0 {
				unit.UnitName = unitName[0]
			}
		}
		if unit.UnitPosition == "" {
			unitPos, err := ReadParam(b, m.regFhx["Unit"])
			if err != nil {
				return nil, err
			}
			if len(unitPos) > 0 {
				unit.UnitPosition = unitPos[0]
			}
		}

		procedures, err := m.readUps(b, unit)

		if err != nil {
			return nil, err
		}
		unit = procedures
		// unit.Procedures = procedures

		units = append(units, unit)
	}
	return units, nil
}

func (m *Fhx) readRecipe(fileText []string) ([]Recipe, error) {

	var recipes = []Recipe{}

	block, err := ReadBlock("BATCH_RECIPE", fileText)
	if err != nil {
		return nil, err
	}

	for _, b := range block {
		var recipe Recipe
		// Names of Unit
		if recipe.RecipeName == "" {
			recipeName, err := ReadParam(b, m.regFhx["Recipe"])
			if err != nil {
				return nil, err
			}
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
	stepBlocks, err := ReadBlock("STEP", lines)
	if err != nil {
		return nil, err
	}
	for _, b := range stepBlocks {
		for _, l := range b {
			name, err := ReadRegex(m.regFhx["Step"], l)
			if err != nil {
				return nil, err
			}
			if name != "" {
				step.Name = name
				// loadUP(name)
			}
			key, err := ReadRegex(m.regFhx["Definition"], l)
			if err != nil {
				return nil, err
			}
			if key != "" {
				step.Key = key
			}
			if step.Description == "" {
				desc, err := ReadRegex(m.regFhx["StepDesc"], l)

				if err != nil {
					return nil, err
				}
				if desc != "" {
					step.Description = desc
				}
			}
			if step.Rect == "" {
				rect, err := ReadRegex(m.regFhx["Rect"], l)

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
		attrBlocks, err := ReadBlock("ATTRIBUTE_INSTANCE", b)
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
func (m *Fhx) readUps(lines []string, unit Unit) (Unit, error) {
	for _, l := range lines {
		if unit.Description == "" {
			desc, err := ReadRegex(m.regFhx["Desc"], l)
			if err != nil {
				return Unit{}, err
			}
			if desc != "" {
				unit.Description = desc
			}
		}

		if unit.Time == 0 {
			time, err := ReadRegex(m.regFhx["Time"], l)
			if err != nil {
				return Unit{}, err
			}

			if time != "" {
				t, err := strconv.Atoi(time)
				if err != nil {
					return Unit{}, err
				}
				unit.Time = t
			}
		}

		if unit.Author == "" {
			author, err := ReadRegex(m.regFhx["Author"], l)
			if err != nil {
				return Unit{}, err
			}
			if author != "" {
				unit.Author = author
			}
		}
	}
	paramBlocks, err := ReadBlock("FORMULA_PARAMETER", lines)

	if err != nil {
		return Unit{}, err
	}
	attrBlocks, err := ReadBlock("ATTRIBUTE_INSTANCE", lines)
	if err != nil {
		return Unit{}, err
	}
	params, err := m.readParameters(paramBlocks, attrBlocks)
	if err != nil {
		return Unit{}, err
	}
	unit.Parameters = params

	// ups = append(ups, up)
	return unit, nil
}

/*
Auslesen der Parameter und hinzufügen ihrer Werte
*/
func (m *Fhx) readParameters(paramBlock [][]string, attrBlock [][]string) ([]Parameter, error) {

	var params = []Parameter{}
	var param = Parameter{}
	for _, b := range paramBlock {
		for _, l := range b {
			name, err := ReadRegex(m.regFhx["Params"], l)
			if err != nil {
				return nil, err
			}
			desc, err := ReadRegex(m.regFhx["Desc"], l)
			if err != nil {
				return nil, err
			}
			if desc != "" {
				param.Description = desc
			}
			if name != "" {
				param.Name = name
				val, err := m.readAttribute(attrBlock, name)
				if err != nil {
					return nil, err
				}
				param.Value = val
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
					u, err := ReadRegex(m.regFhx["ValueSet"], b[4])
					if err != nil {
						return val, err
					}
					val.Set = u
					strVal, err := ReadRegex(m.regFhx["ValueString"], b[5])
					if err != nil {
						return val, err
					}
					val.StringValue = strVal
					return val, nil
				}

				// Nur bei Texten wie Melden
				//fmt.Println(len(b), i+2)

				parseLine := b[2]

				// d, err := ReadRegex(m.regFhx["ValueDesc"], parseLine)
				// if err != nil {
				// 	return val, err
				// }
				// if d != "" {
				// 	val.Cv = d
				// 	return val, nil
				// }

				// Werte für Zahlen
				v, err := ReadRegexSubexp(m.regFhx["Value"], parseLine)
				if len(v) == 0 {
					return val, nil
				}

				if err != nil {
					return val, err
				}
				h, err := strconv.Atoi(v["s1"])
				if err != nil {
					return val, fmt.Errorf("ReadRegexSubexp v[s1]. %v", err)
				}
				val.High = h
				l, err := strconv.Atoi(v["s2"])
				if err != nil {
					return val, fmt.Errorf("ReadRegexSubexp v[s2]. %v", err)
				}
				val.Low = l
				cv, err := strconv.Atoi(v["s3"])
				if err != nil {
					return val, fmt.Errorf("ReadRegexSubexp v[s3]. %v", err)
				}
				val.Cv = cv
				val.Unit = v["s4"]
				return val, nil
			}
		}
	}
	//fmt.Println(val)
	return val, nil
}
