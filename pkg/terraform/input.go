package terraform

import (
	"encoding/json"
	"os"
)

type Variable struct {
	Sensitive bool        `json:"sensitive"`
	Type      string      `json:"type"`
	Value     interface{} `json:"value"`
}

type Data struct {
	Variables map[string]Variable `json:"-"` // skip the first-level key during unmarshalling
}

func GetOutputValues(file string) (*Data, error) {
	data := &Data{}
	jsonData, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &data.Variables)
	if err != nil {
		return nil, err
	}
	return data, nil
}
