package ast

import (
	"encoding/json"
	"fmt"
)

func toJSON(i interface{}) (string, error) {
	m := map[string]interface{}{}
	m[fmt.Sprintf("%T", i)] = i
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
