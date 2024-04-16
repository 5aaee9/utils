package sysproxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SystemProxyStatus(t *testing.T) {
	proxy := NewSystemProxy()

	status, err := proxy.Status()
	assert.NoError(t, err)
	assert.NotNil(t, status)
}
