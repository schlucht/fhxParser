package fhx

import (
	"log"
	"strings"
	
)

var list map[string][]string

type UnitList struct {
	UnitNames []string
	Units     []Procedure
}

func (m *UnitList) NewUnitList() {
	var err error
	list, err = LoadAllStandardFilename()
	var names []string
	if err != nil {
		m.UnitNames = names
		log.Fatalln(err)
	}
	for k := range list {
		names = append(names, k)
	}
	m.UnitNames = names
}

func (m *UnitList) UPNames(unit string) []string {
	var s []string

	for _, n := range list[unit] {
		p := strings.TrimRight(n, ".json")
		s = append(s, p)
	}
	return s
}
