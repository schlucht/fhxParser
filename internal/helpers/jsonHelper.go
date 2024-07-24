package helpers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func PrintJson(input interface{}) string {
	data, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return ""
	}
	return string(data)
}

func SaveJSON(path string, data interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	_, err = fmt.Fprintf(w, "%v", data)
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}

func OpenJSON(path string) ([]byte, error) {

	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteJson, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	return byteJson, nil
}
