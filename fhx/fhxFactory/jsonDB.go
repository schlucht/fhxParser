package fhx

import (
	"fmt"
	"io/fs"
	"os"
)

func SaveFhxFile(pathName string) {

	// Alle Ordner auslesen
	// Einlesen der Vorhanden FHX
	// Druchlauf der FHX und die alten FHX durch die neuen
	// Ersetzen

	// fhxs := fhxModels.NewFhxPath(pathName)

	// for _, f := range fhxs {
	// 	fmt.Println(f.Unitname)
	// }
	m := loadFolder()
	for k, v := range m {
		fmt.Printf("%s\n", k)
		for _, f := range v {
			fmt.Printf("\t%s\n", f)
		}
	}
	// b, err := json.Marshal(obj)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// p := filepath.Join("./", "jsonDb", pathName+".json")
	// f, err := os.Create(p)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// f.Write(b)
}

func loadFolder() (sys map[string][]string) {
	fsys := os.DirFS("./jsonDb")
	sys = make(map[string][]string)
	folder := ""
	var s []string
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			folder = d.Name()
			s = nil
		} else {
			s = append(s, d.Name())
			sys[folder] = s
		}
		return nil
	})
	return sys
}

func LoadAllStandardFilename() (map[string][]string, error) {
	return make(map[string][]string), nil
}
