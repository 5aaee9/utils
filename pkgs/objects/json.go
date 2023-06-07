package objects

import (
	"encoding/json"
	"errors"
	"reflect"
)

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

var ErrTypeError = errors.New("type error")

func RemarshJSONReflect[T any](input any, kind reflect.Type) (*T, error) {
	data, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	if kind.Kind() == reflect.Ptr {
		typ := kind.Elem()
		dst := reflect.New(typ).Elem()
		err := json.Unmarshal(data, dst.Addr().Interface())
		if err != nil {
			return nil, err
		}

		res, ok := dst.Addr().Interface().(*T)
		if !ok {
			return nil, ErrTypeError
		}

		return res, nil
	}

	dst := reflect.New(kind).Elem()
	err = json.Unmarshal(data, dst.Addr().Interface())
	if err != nil {
		return nil, err
	}

	res, ok := dst.Interface().(T)
	if !ok {
		return nil, ErrTypeError
	}

	return &res, nil
}
