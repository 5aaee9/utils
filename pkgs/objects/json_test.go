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
