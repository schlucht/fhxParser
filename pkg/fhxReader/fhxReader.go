package fhxReader

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

/*
	Liest einen ganzen Block aus einer fhx Datei. Es wird alles zwischen einem öffnenden und Schliessenden {} in ein Array gelesen.

Return ist ein Array mit dem Blocknamen und dann mit allen Zeilen als Array.
Parameter ein String zum identifizieren des Blocks und den zu durchsuchenden Text als Arra
*/
func ReadBlock(startString string, lines []string) ([][]string, error) {

	results := [][]string{}
	if strings.Trim(startString, "") == "" {
		return results, errors.New("kein suchparameter vorhanden")
	}
	if len(lines) == 0 {
		return results, errors.New("kein text übergeben")
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

/* Durchläuft ein Textarray und sucht mittels Regex den gesuchten Eintrag. Der Eintrag wird ausgelesen mit einem Schlüssel Wert paar zurück gegeben*/
func ReadRegexMap(regex map[string]string, txt []string) map[string]interface{} {
	res := make(map[string]interface{})
	for _, l := range txt {
		for key, r := range regex {
			rCompile := regexp.MustCompile(r)
			matches := rCompile.FindStringSubmatch(l)

			if len(matches) > 0 {
				if rCompile.SubexpIndex("s") > -1 {
					res[key] = matches[rCompile.SubexpIndex("s")]
				}
				if rCompile.SubexpIndex("i") > -1 {
					i, err := strconv.ParseInt(matches[rCompile.SubexpIndex("i")], 10, 32)
					if err != nil {
						log.Printf("%v\n", err)
					}
					res[key] = i
				}
				if rCompile.SubexpIndex("b") > -1 {
					b := matches[rCompile.SubexpIndex("b")]
					res[key] = b == "T"
				}
			}
		}
	}
	return res
}

/* Durchsucht einen String ob das Pattern in der Zeile vorhanden ist. Es gibt den gesuchten Wert als String zurück*/
func ReadRegex(regex string, txt string) (string, error) {
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
func ReadRegexSubexp(regex string, txt string) (map[string]string, error) {
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

/* Liest ein fhx File im UTF16L in ein Array ein*/
func ReadFhxFile16(filePath string) ([]string, error) {
	var res []string
	if filePath == "" {
		return res, errors.New("filePath is empty")
	}
	file, err := os.Open(filePath)
	if err != nil {!
		return res, err
	}

	scanner := bufio.NewScanner(transform.NewReader(file, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))
	for scanner.Scan() {
		res = append(res, (scanner.Text()))
	}
	return res, nil
}

/* Testet ein fhx Pfad ob es ein*/
func IsFhxFile(pathStr string) string {
	ext := path.Ext(pathStr)

	return strings.ToUpper(ext)
}

/* Splittet ein Text beim Zeilen Umbruch*/
func SplitLines(text string) []string {
	return strings.Split(text, "\n")
}

/*
	Matcht den Wert aus einem einer Liste mit fhx Zeilen

Parameter: lines ein Array mit den Zeilen
regex ein Regex Suchstring
Rückgabe alle Wert die auf das Regex gepasst haben
*/
func ReadParam(lines []string, regex string) []string {
	var results []string
	for _, l := range lines {
		u, _ := ReadRegex(regex, l)
		if u != "" {
			results = append(results, u)
		}
	}
	return results
}
