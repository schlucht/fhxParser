package parser

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Fhx struct {
	Recipes  []Recipe          `json:"recipes,omitempty"`
	Units    []Unit            `json:"units,omitempty"`
	OPs      []Unit            `json:"ops,omitempty"`
	regFhx   map[string]string `json:"-"`
	UnitType string            `json:"unit_type,omitempty"`
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
	"Origin":      `ORIGIN=(?P<s>.*)`,
	"DeferTo":     `DEFERRED_TO="(?P<s>.*)"`,
	"Group":       `GROUP="(?P<s>.*)"`,
	"StepParams":  `STEP_PARAMETER NAME="(?P<s>.*)"`,
}

// var fhx = Fhx{}

/*
Einstiegspunkt zum laden der Daten aus einer Fhx Datei
*/

// Ein FHX String einlesen
func NewFhxString(fhxText string) ([]Fhx, error) {
	var fhx = Fhx{regFhx: regFhx}
	if fhxText == "" {
		return nil, errors.New("NewFHXString, no file i")
	}
	lines, err := readFhxText(fhxText, "")
	if err != nil {
		return nil, err
	}

	fhxs, err := fhx.readFhx(lines)

	if err != nil {
		return nil, err
	}
	return fhxs, nil
}

/*
Wird ausgeführt wenn ein FHX Datei eingelesen wird diese Eingelesen und neu gespeichert
*/
func (m *Fhx) readFhx(fileText []string) ([]Fhx, error) {

	var fhxs = []Fhx{}
	block, err := readBlock("BATCH_RECIPE", fileText)
	// fmt.Printf("Erstes File: %v\n", block[0])
	if err != nil {
		return nil, err
	}

	for _, b := range block {
		var fhx Fhx
		// Type der Datei PROCEDURE OPERATION UNITPROCEDURE
		unitType, err := readParam(b, m.regFhx["Type"])
		if err != nil {
			return nil, err
		}

		if unitType[0] != "" {
			fhx.UnitType = unitType[0]

			if unitType[0] == "UNIT_PROCEDURE" {
				units, err := m.readUnit(b, "UP")
				if err != nil {
					return nil, err
				}
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

	block, err := readBlock("BATCH_RECIPE", fileText)
	if err != nil {
		return nil, err
	}

	for _, b := range block {
		var unit Unit
		// Names of Unit
		unit.Type = unit_type
		if unit.UnitName == "" {
			unitName, err := readParam(b, m.regFhx["Recipe"])
			if err != nil {
				return nil, err
			}
			if len(unitName) > 0 {
				unit.UnitName = unitName[0]
			}
		}
		if unit.UnitPosition == "" {
			unitPos, err := readParam(b, m.regFhx["Unit"])
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
		// steps, err := m.readStep(b)
		// if err != nil {
		// 	return nil, err
		// }
		// unit.Steps = steps

		units = append(units, unit)
	}
	return units, nil
}

func (m *Fhx) readRecipe(fileText []string) ([]Recipe, error) {

	var recipes = []Recipe{}

	block, err := readBlock("BATCH_RECIPE", fileText)
	if err != nil {
		return nil, err
	}

	for _, b := range block {
		var recipe Recipe
		// Names of Unit
		if recipe.RecipeName == "" {
			recipeName, err := readParam(b, m.regFhx["Recipe"])
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

		parameters, err := m.readAttributeParamBlock(b)
		if err != nil {
			return nil, err
		}

		recipe.Parameters = parameters
		recipes = append(recipes, recipe)

	}
	return recipes, nil
}

func (m *Fhx) readStep(lines []string) ([]Step, error) {
	var steps = []Step{}
	step := Step{}
	stepBlocks, err := readBlock("STEP", lines)
	if err != nil {
		return nil, err
	}
	for _, b := range stepBlocks {
		for _, l := range b {
			name, err := readRegex(m.regFhx["Step"], l)
			if err != nil {
				return nil, err
			}
			if name != "" {
				step.Name = name
				// loadUP(name)
			}
			key, err := readRegex(m.regFhx["Definition"], l)
			if err != nil {
				return nil, err
			}
			if key != "" {
				step.Key = key
			}
			if step.Description == "" {
				desc, err := readRegex(m.regFhx["StepDesc"], l)

				if err != nil {
					return nil, err
				}
				if desc != "" {
					step.Description = desc
				}
			}
			if step.Rect == "" {
				rect, err := readRegex(m.regFhx["Rect"], l)

				if err != nil {
					return nil, err
				}
				if rect != "" {
					step.Description = rect
				}
			}
		}
		//
		// log.Println(stepparams)
		attrBlocks, err := readBlock("STEP_PARAMETER", b)
		if err != nil {
			return nil, err
		}
		params, err := m.stepParameters(attrBlocks)
		if err != nil {
			return nil, err
		}
		step.StepParameters = params
		steps = append(steps, step)

		// fmt.Println(steps)
	}

	return steps, err
}

func (m *Fhx) stepParameters(attrBlock [][]string) ([]StepParameter, error) {
	params := []StepParameter{}
	for _, b := range attrBlock {
		param := StepParameter{}
		for _, l := range b {
			// fmt.Printf("%v", l)
			name, err := readRegex(m.regFhx["StepParams"], l)
			// fmt.Println(name)
			if err != nil {
				return nil, err
			}
			if name != "" {
				param.Name = name
			}
			origin, err := readRegex(m.regFhx["Origin"], l)
			if err != nil {
				return nil, err
			}
			if origin != "" {
				param.Origin = origin
			}
			deferTo, err := readRegex(m.regFhx["DeferTo"], l)
			if err != nil {
				return nil, err
			}
			if deferTo != "" {
				param.DeferTo = deferTo
			}
			group, err := readRegex(m.regFhx["Group"], l)
			if err != nil {
				return nil, err
			}
			if group != "" {
				param.Group = group
			}
		}
		// fmt.Println(param)
		params = append(params, param)
	}
	// TODO: READ STEP Paramter
	return params, nil
}

/*
Liest die Unit Proceduren aus der FHX Datei
*/
func (m *Fhx) readUps(lines []string, unit Unit) (Unit, error) {
	for _, l := range lines {
		if unit.Description == "" {
			desc, err := readRegex(m.regFhx["Desc"], l)
			if err != nil {
				return Unit{}, err
			}
			if desc != "" {
				unit.Description = desc
			}
		}

		if unit.Time == 0 {
			time, err := readRegex(m.regFhx["Time"], l)
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
			author, err := readRegex(m.regFhx["Author"], l)
			if err != nil {
				return Unit{}, err
			}
			if author != "" {
				unit.Author = author
			}
		}
	}
	paramBlocks, err := readBlock("FORMULA_PARAMETER", lines)

	if err != nil {
		return Unit{}, err
	}
	attrBlocks, err := readBlock("ATTRIBUTE_INSTANCE", lines)
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

// Liest die Paramter aus den Attribute Blöcken und
// Paramtern Blöcken
func (m *Fhx) readAttributeParamBlock(lines []string) ([]Parameter, error) {
	paramBlocks, err := readBlock("FORMULA_PARAMETER", lines)
	if err != nil {
		return []Parameter{}, err
	}
	log.Println(paramBlocks)
	attrBlocks, err := readBlock("ATTRIBUTE_INSTANCE", lines)
	if err != nil {
		return []Parameter{}, err
	}
	params, err := m.readParameters(paramBlocks, attrBlocks)
	if err != nil {
		return []Parameter{}, err
	}

	return params, nil
}

/*
Auslesen der Parameter und hinzufügen ihrer Werte
*/
func (m *Fhx) readParameters(paramBlock [][]string, attrBlock [][]string) ([]Parameter, error) {

	var params = []Parameter{}
	var param = Parameter{}
	for _, b := range paramBlock {
		for _, l := range b {
			name, err := readRegex(m.regFhx["Params"], l)
			if err != nil {
				return nil, err
			}
			desc, err := readRegex(m.regFhx["Desc"], l)
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
					u, err := readRegex(m.regFhx["ValueSet"], b[4])
					if err != nil {
						return val, err
					}
					val.Set = u
					strVal, err := readRegex(m.regFhx["ValueString"], b[5])
					if err != nil {
						return val, err
					}
					val.StringValue = strVal
					return val, nil
				}

				parseLine := b[2]

				v, err := readRegexSubexp(m.regFhx["Value"], parseLine)
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
