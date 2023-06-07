package objects_test

import (
	"reflect"
	"testing"

	"github.com/5aaee9/utils/pkgs/json"
	"github.com/5aaee9/utils/pkgs/objects"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Data string `json:"test"`
}

func TestRemarshJSON(t *testing.T) {
	out, err := objects.RemarshJSON[TestStruct](json.JSON{
		"test": "data test #1",
	})

	assert.NoError(t, err)
	assert.Equal(t, out.Data, "data test #1")
}

func TestRemarshJSONReflect(t *testing.T) {
	out, err := objects.RemarshJSONReflect[TestStruct](json.JSON{
		"test": "data test #1",
	}, reflect.TypeOf(TestStruct{}))

	assert.NoError(t, err)
	assert.Equal(t, out.Data, "data test #1")

	ptr := new(TestStruct)

	out, err = objects.RemarshJSONReflect[TestStruct](json.JSON{
		"test": "data test #1",
	}, reflect.ValueOf(ptr).Type())
	assert.NoError(t, err)
	assert.Equal(t, out.Data, "data test #1")
	assert.Equal(t, ptr.Data, "")
}

func BenchmarkRemarshJSON(b *testing.B) {
	data := json.JSON{
		"test": "data test #1",
	}

	for n := 0; n < b.N; n++ {
		objects.RemarshJSON[TestStruct](data)
	}
}

func BenchmarkRemarshJSONReflect(b *testing.B) {
	data := json.JSON{
		"test": "data test #1",
	}
	kind := reflect.TypeOf(TestStruct{})

	for n := 0; n < b.N; n++ {
		objects.RemarshJSONReflect[TestStruct](data, kind)
	}
}
