package helpers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func PrintJson(input interface{}) (string, error) {
	data, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(data), nil
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
