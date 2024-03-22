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
	expected := `{
  "Name": "John",
  "Vorname": "Doe"
}`
	result := PrintJson(input)
	if result != expected {
		t.Errorf("PrintJson() returned unexpected result: got %v want %v", result, expected)
	}
}

func TestSaveJSON(t *testing.T) {
	// Test SaveJSON function
	data := struct {
		Key   string
		Value int
	}{
		Key:   "example",
		Value: 123,
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
