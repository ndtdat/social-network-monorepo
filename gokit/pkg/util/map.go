package util

import (
	"encoding/json"
	"fmt"
)

func UnmarshalMapString(m map[string]string) []byte {
	if m == nil {
		m = map[string]string{}
	}
	mb, _ := json.Marshal(m)

	return mb
}

func StructToMap(object any) (map[string]any, error) {
	jsonBytes, err := json.Marshal(object)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal map due to: %v", err)
	}
	var results map[string]any
	if err = json.Unmarshal(jsonBytes, &results); err != nil {
		return nil, fmt.Errorf("cannot unmarshal json string due to: %v", err)
	}

	return results, nil
}

func MapToStruct(obj map[string]any) (any, error) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal map due to: %v", err)
	}
	var results any
	if err = json.Unmarshal(jsonBytes, &results); err != nil {
		return nil, fmt.Errorf("cannot unmarshal json string due to: %v", err)
	}

	return results, nil
}
