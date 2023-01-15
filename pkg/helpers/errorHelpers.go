package helpers

import "log"

func FatalError(err error) bool {
	if err != nil {
		log.Fatal(err)
	}
	return true
}
