package errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError(t *testing.T) {
	assert.NoError(t, WithStack(nil))
	assert.True(t, HasStack(WithStack(errors.New("test"))))
	assert.False(t, HasStack(errors.New("test")))

	//helper.InitDebugSentry()
	//sentry.CaptureException(WithStack(errors.New("test")))
	//sentry.Flush(time.Second * 10)
}
