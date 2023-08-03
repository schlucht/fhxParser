package parser

import "testing"

func Test_printJson(t *testing.T) {
	i := struct {
		Name    string
		Vorname string
	}{
		Name:    "lothar",
		Vorname: "schmid",
	}
	res, err := PrintJson(i)
	if res != `{"Name":"lothar","Vorname":"schmid"}` {
		t.Errorf("%s json string ist falsch", res)
	}
	if err != nil {
		t.Error(err)
	}

}
