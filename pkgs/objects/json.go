package objects

import "encoding/json"

func RemarshJSON[T any](input any) (*T, error) {
	bytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	r := new(T)
	if err = json.Unmarshal(bytes, r); err != nil {
		return nil, err
	}

	return r, nil
}
