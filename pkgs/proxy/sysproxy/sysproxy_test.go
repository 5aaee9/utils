//go:build windows || darwin

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

	err = proxy.TurnOff()
	assert.NoError(t, err)

}

// func Test_CloseSystemProxy(t *testing.T) {
// 	proxy := NewSystemProxy()
// 	err := proxy.TurnOff()
// 	assert.NoError(t, err)
// }

// func Test_OpenSystemProxy(t *testing.T) {
// 	proxy := NewSystemProxy()
// 	err := proxy.TurnOn("127.0.0.1:9090")

// 	assert.NoError(t, err)
// }
