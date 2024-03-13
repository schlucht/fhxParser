package helpers

import "testing"

func Test_printJson(t *testing.T) {
	i := struct {
		Name    string
		Vorname string
	}{
		Name:    "lothar",
		Vorname: "schmid",
	}
	res := PrintJson(i)
	if res != `{"Name":"lothar","Vorname":"schmid"}` {
		t.Errorf("%s json string ist falsch", res)
	}

}
