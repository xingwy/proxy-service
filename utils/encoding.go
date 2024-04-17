package utils

import "encoding/json"

func JSONEncoding(v any) (string, error) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func JSONEncodingWithoutErr(v any) string {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(jsonData)
}
