package sysproxy

import (
	"errors"
	"fmt"
	"math/big"
	"syscall"
	"unsafe"

	"github.com/samber/do"
)

type WindowsSystemProxy struct{}

func NewSystemProxy() SystemProxy {
	return &WindowsSystemProxy{}
}

var _ SystemProxy = (*WindowsSystemProxy)(nil)

var (
	wininet, _           = syscall.LoadLibrary("Wininet.dll")
	internetSetOption, _ = syscall.GetProcAddress(wininet, "InternetSetOptionW")
)

const (
	internetOptionPerConnectionOption  = 75
	internetOptionProxySettingsChanged = 95
	internetOptionRefresh              = 37
)

const (
	proxyTypeDirect = 0x00000001 // direct to net
	proxyTypeProxy  = 0x00000002 // via named proxy
)

const (
	internetPerConnFlags       = 1
	internetPerConnProxyServer = 2
)

type internetPerConnOptionList struct {
	dwSize        uint32
	pszConnection *uint16
	dwOptionCount uint32
	dwOptionError uint32
	pOptions      uintptr
}

type internetPreConnOption struct {
	dwOption uint32
	value    uint64
}

func stringPtrAddr(str string) (uint64, error) {
	scriptLocPtr, err := syscall.UTF16PtrFromString(str)
	if err != nil {
		return 0, err
	}
	n := new(big.Int)
	n.SetString(fmt.Sprintf("%x\n", scriptLocPtr), 16)
	return n.Uint64(), nil
}

func newParam(n int) internetPerConnOptionList {
	return internetPerConnOptionList{
		dwSize:        4,
		pszConnection: nil,
		dwOptionCount: uint32(n),
		dwOptionError: 0,
		pOptions:      0,
	}
}

func (p *WindowsSystemProxy) TurnOn(addrport string) error {
	proxyServerPtrAddr, err := stringPtrAddr(addrport)
	if err != nil {
		return err
	}

	param := newParam(2)
	options := []internetPreConnOption{
		{dwOption: internetPerConnFlags, value: proxyTypeProxy | proxyTypeDirect},
		{dwOption: internetPerConnProxyServer, value: proxyServerPtrAddr},
	}

	param.pOptions = uintptr(unsafe.Pointer(&options[0]))
	ret, _, infoPtr := syscall.SyscallN(internetSetOption,
		4,
		0,
		internetOptionPerConnectionOption,
		uintptr(unsafe.Pointer(&param)),
		unsafe.Sizeof(param),
		0, 0)

	if ret != 1 {
		return errors.New(fmt.Sprintf("%s", infoPtr))
	}

	return p.Flush()
}
func (p *WindowsSystemProxy) TurnOff() error {
	param := newParam(1)
	option := internetPreConnOption{
		dwOption: internetPerConnFlags,
		//value:    _PROXY_TYPE_AUTO_DETECT | _PROXY_TYPE_DIRECT}
		value: proxyTypeDirect}
	param.pOptions = uintptr(unsafe.Pointer(&option))
	ret, _, infoPtr := syscall.SyscallN(internetSetOption,
		4,
		0,
		internetOptionPerConnectionOption,
		uintptr(unsafe.Pointer(&param)),
		unsafe.Sizeof(param),
		0, 0)

	if ret != 1 {
		return errors.New(fmt.Sprintf("%s", infoPtr))
	}

	return p.Flush()
}

func (p *WindowsSystemProxy) Flush() error {
	ret, _, infoPtr := syscall.SyscallN(internetSetOption,
		4,
		0,
		internetOptionProxySettingsChanged,
		0, 0,
		0, 0)

	if ret != 1 {
		return errors.New(fmt.Sprintf("%s", infoPtr))
	}

	ret, _, infoPtr = syscall.SyscallN(internetSetOption,
		4,
		0,
		internetOptionRefresh,
		0, 0,
		0, 0)

	if ret != 1 {
		return errors.New(fmt.Sprintf("%s", infoPtr))
	}
	return nil
}
