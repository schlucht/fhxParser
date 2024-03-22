package helpers

import (
	"os"
	"testing"
)

func TestPrintJson(t *testing.T) {
	// Test PrintJson function
	input := struct {
		Name    string
		Vorname string
	}{
		Name:    "John",
		Vorname: "Doe",
	}
	res := PrintJson(i)
	if res != `{"Name":"lothar","Vorname":"schmid"}` {
		t.Errorf("%s json string ist falsch", res)
	}

	// Provide a temporary file path for testing
	filePath := "testdata/test.json"

	err := SaveJSON(filePath, data)
	if err != nil {
		t.Errorf("SaveJSON() returned an error: %v", err)
	}

	// Read the file to check if the data was saved correctly
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	expected := `{"Key":"example","Value":123}`
	if string(fileData) != expected {
		t.Errorf("SaveJSON() did not write the correct data to the file. Got: %v, Want: %v", string(fileData), expected)
	}
}
