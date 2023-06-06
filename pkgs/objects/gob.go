package objects

import (
	"bytes"
	"encoding/gob"
)

func RemarshGob[T any](input any) (*T, error) {
	buf := bytes.Buffer{}
	out := new(T)

	if err := gob.NewEncoder(&buf).Encode(input); err != nil {
		return nil, err
	}

	if err := gob.NewDecoder(&buf).Decode(out); err != nil {
		return nil, err
	}

	return out, nil
}

func MustRemarshGob[T any](input any) *T {
	r, err := RemarshGob[T](input)
	if err != nil {
		return nil
	}

	return r
}
