package parser

import (
	"errors"
	"regexp"
	"strings"
)

/*
	Liest einen ganzen Block aus einer fhx Datei. Es wird alles zwischen einem öffnenden und Schliessenden {} in ein Array gelesen.

Return ist ein Array mit dem Blocknamen und dann mit allen Zeilen als Array.
Parameter ein String zum identifizieren des Blocks und den zu durchsuchenden Text als Arra
*/
func readBlock(startString string, lines []string) ([][]string, error) {

	results := [][]string{}
	if strings.Trim(startString, "") == "" {
		return nil, errors.New("kein suchparameter vorhanden")
	}
	if len(lines) == 0 {
		return nil, errors.New("kein text übergeben")
	}

	regParam, _ := regexp.Compile(startString)

	var start = false
	var curlybreak = 0
	var blockLines = []string{}

	for _, l := range lines {

		if start {
			blockLines = append(blockLines, l)
			if strings.Contains(l, "{") {
				sl := strings.Trim(l, " ")
				if len(sl) <= 2 {
					curlybreak++
				}
			}
			if strings.Contains(l, "}") {
				sl := strings.Trim(l, " ")
				if len(sl) <= 2 {
					curlybreak--
					if curlybreak == 0 {
						start = false
						results = append(results, blockLines)
						blockLines = []string{}
					}
				}
			}
		} else {
			start = regParam.MatchString(l)
			if start {
				blockLines = append(blockLines, l)
			}
		}
	}
	return results, nil
}

/* Durchsucht einen String ob das Pattern in der Zeile vorhanden ist. Es gibt den gesuchten Wert als String zurück*/
func readRegex(regex string, txt string) (string, error) {
	res := ""
	if regex == "" {
		return res, errors.New("no regexpattern for search")
	}
	compile := regexp.MustCompile(regex)
	matches := compile.FindStringSubmatch(txt)
	if len(matches) > 0 {
		if compile.SubexpIndex("s") > -1 {
			res = matches[compile.SubexpIndex("s")]
		}
	}
	return res, nil
}

// readRegexSubexp reads the named subexpressions in the regular expression and returns a map of named subexpressions to their matches.
//
// It takes a regex string and a txt string as parameters and returns a map[string]string and an error.
func readRegexSubexp(regex string, txt string) (map[string]string, error) {
	res := make(map[string]string)
	if regex == "" {
		return res, errors.New("no regexpattern for search")
	}
	compile := regexp.MustCompile(regex)
	matches := compile.FindStringSubmatch(txt)

	if len(matches) > 0 {
		for _, s := range compile.SubexpNames() {
			if s != "" {
				res[s] = matches[compile.SubexpIndex(s)]
			}
		}
	}
	return res, nil
}

// /* Read fhx File in UTF-8 Format*/
// func readFhx(path string) ([]string, error) {
// 	f, err := os.Open(path)
// 	var lines []string
// 	if err != nil {
// 		return lines, err
// 	}
// 	r := bufio.NewScanner(f)
// 	r.Split(bufio.ScanLines)
// 	for r.Scan() {
// 		lines = append(lines, r.Text())
// 	}
// 	return lines, nil
// }

// Liest ein Text aus einer FHX Datei ein
func readFhxText(text string, sep string) ([]string, error) {

	var lines []string
	if text == "" {
		return lines, errors.New("file is empty")
	}
	if sep != "" {
		lines = splitLinesSep(text, sep)
	} else {
		lines = splitLines(text)
	}

	if len(lines) == 0 {
		return lines, errors.New("file is empty")
	}

	return lines, nil
}

// /* Testet ein fhx Pfad ob es ein*/
// func isFhxFile(pathStr string) error {
// 	ext := path.Ext(pathStr)
// 	if strings.ToUpper(ext) != ".FHX" {
// 		return errors.New("file is not a fhx file")
// 	}
// 	return nil
// }

/* Splittet ein Text beim Zeilen Umbruch*/
func splitLines(text string) []string {
	return strings.Split(text, "\n")
}

func splitLinesSep(text string, sep string) []string {
	return strings.Split(text, sep)
}

/*
	Matcht den Wert aus einem einer Liste mit fhx Zeilen

Parameter: lines ein Array mit den Zeilen
regex ein Regex Suchstring
Rückgabe alle Wert die auf das Regex gepasst haben
*/
func readParam(lines []string, regex string) ([]string, error) {
	var results []string
	for _, l := range lines {
		u, err := readRegex(regex, l)
		if err != nil {
			return nil, err
		}
		if u != "" {
			results = append(results, u)
		}
	}
	return results, nil
}
