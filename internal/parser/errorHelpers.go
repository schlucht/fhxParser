package parser

import (
	"log"
)

func FatalError(err error) bool {
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func ErrFatal(linenumber string, module string, method string, err error) {
	log.Fatalf("Error Message:\n%s: (%s: '%s'), %s\n", linenumber, module, method, err)
}
