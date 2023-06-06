package objects_test

import (
	"testing"

	"github.com/5aaee9/utils/pkgs/objects"
	"github.com/stretchr/testify/assert"
)

func TestRemarshGob(t *testing.T) {
	// Copy with same field name
	a := struct {
		A string `json:"test"`
		B *string
	}{A: "test"}

	b, err := objects.RemarshGob[struct {
		A string `json:"another"`
	}](a)

	assert.NoError(t, err)
	assert.Equal(t, "test", b.A)

	// Copy without same field name
	c := struct {
		Test string `json:"test"`
	}{Test: "test"}

	_, err = objects.RemarshGob[struct {
		A string `json:"another"`
	}](c)
	assert.Error(t, err)
}
