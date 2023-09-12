package common

import "encoding/json"

func TypeConvert[T any](data any) (*T, error) {
	var result T
	jsonedData, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonedData, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
