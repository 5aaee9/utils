package json_test

import (
	"testing"

	"github.com/5aaee9/utils/pkgs/json"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Data string `json:"test"`
}

func TestIntoJSON(t *testing.T) {
	out, err := json.IntoJSON(TestStruct{Data: "test random str"})
	assert.NoError(t, err)
	assert.Equal(t, "test random str", out["test"])
}

func TestIntoJSONPointer(t *testing.T) {
	out, err := json.IntoJSON(&TestStruct{Data: "test random str"})
	assert.NoError(t, err)
	assert.Equal(t, "test random str", out["test"])
}

func TestMergeJSON(t *testing.T) {
	out, err := json.IntoJSON(TestStruct{Data: "test random str"})
	assert.NoError(t, err)

	out = json.MergeJSON(out, json.JSON{"data": 123})
	assert.Equal(t, "test random str", out["test"])
	assert.Equal(t, 123, out["data"])
}
