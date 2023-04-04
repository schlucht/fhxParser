package fhxFactory

import (
	"strings"

	"github.com/schlucht/fhxreader/database"
)

func (ff *FhxFactory) SaveUnits() error {
	db := database.New()
	for _, fhx := range ff.Fhx {
		for _, u := range fhx.Units {
			err := db.InsertUnit(u)
			if err != nil {
				errNr := strings.Split(err.Error(), " ")[1]
				if errNr == "1062" {
					continue
				} else {
					return err
				}
			}
		}
	}
	return nil
}

// func (ff *FhxFactory) SaveProcedures() error {
// 	db := database.New()
// 	for _, fhx := range ff.Fhx {
// 		for _, u := range fhx.Units {
// 			for _, p := range u.Procedures {
// 				i, err := db.IdFromUnitname(p.Name)
// 				if err != nil {
// 					return err
// 				}
// 				log.Println("Die ID ist:", i)
// 			}
// 		}
// 	}
// 	return nil
// }
