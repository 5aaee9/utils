//go:build !(windows || darwin)

package sysproxy

import "errors"

type OtherSystemProxy struct{}

var _ SystemProxy = (*OtherSystemProxy)(nil)

func (p *OtherSystemProxy) TurnOff() error {
	return errors.New("not implemented")
}

func (p *OtherSystemProxy) TurnOn(addrport string) error {
	return errors.New("not implemented")
}

func NewSystemProxy() SystemProxy {
	return &OtherSystemProxy{}
}

func (p *OtherSystemProxy) Status() (*SystemProxyStatus, error) {
	return nil, errors.New("not implemented")
}
