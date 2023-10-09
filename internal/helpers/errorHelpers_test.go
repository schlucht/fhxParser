package helpers

import (
	"errors"
	"testing"
)

func Test_FatalError(t *testing.T) {
	res := []struct {
		name   string
		input  error
		output bool
	}{

		{name: "keine Fehler", input: nil},
		{name: "mit Fehler", input: errors.New("kein wert vorhanden")},
	}

	res[0].output = FatalError(res[0].input)
	if !res[0].output {
		t.Errorf("Test nicht bestanden %s", res[0].name)
	}

}
