package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMeasureFuncTime(t *testing.T) {
	ret := MeasureFuncTime(func() {
		time.Sleep(time.Second * 2)
	})

	assert.Equal(t, 2000, int(ret/time.Millisecond))
}
