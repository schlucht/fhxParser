package helpers

import "encoding/json"

func PrintJson(input interface{}) (string, error) {
	data, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	return string(data), nil

}
