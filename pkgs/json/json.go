package json

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

func PettryJSON(input JSON) string {
	data, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func MergeJSON(a, b JSON) JSON {
	for k, v := range b {
		a[k] = v
	}

	return a
}

func IntoJSON[T any](input T) (JSON, error) {
	var data JSON
	b, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		str, ok := value.(string)
		if !ok {
			return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
		}

		bytes = []byte(str)
	}

	result := JSON{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j JSON) GormDataType() string {
	return "JSON"
}
