package sysproxy

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

var firmwareType = []string{
	"webproxy",
	"securewebproxy",
	"socksfirewallproxy",
}

type DarwinSystemProxy struct {
}

func RunBashCmd(cmd string) (string, error) {
	c := exec.Command("sh", "-c", cmd)
	out, err := c.CombinedOutput()
	if err != nil {
		return "", errors.New(string(out) + err.Error())
	}
	return strings.TrimSpace(string(out)), nil
}

func GetNetworkInterface() (string, error) {
	return RunBashCmd("networksetup -listnetworkserviceorder | grep -B 1 $(route -n get default | grep interface | awk '{print $2}') | head -n 1 | sed 's/.*) //'")
}

var _ SystemProxy = (*DarwinSystemProxy)(nil)

type RequestResponse struct {
	Error string `json:"error,omitempty"`
	Code  uint   `json:"code,omitempty"`
}

func (p *DarwinSystemProxy) TurnOff() error {
	s, err := GetNetworkInterface()
	if err != nil {
		return err
	}

	for _, t := range firmwareType {
		if _, err := RunBashCmd(fmt.Sprintf("networksetup -set%sstate %s off", t, s)); err != nil {
			return err
		}
	}

	return nil
}

func (p *DarwinSystemProxy) Status() (*SystemProxyStatus, error) {
	cmd := exec.Command("scutil", "--proxy")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	info := strings.ReplaceAll(string(output), " ", "")

	status := new(SystemProxyStatus)
	status.State = strings.Contains(info, "HTTPEnable:1") || strings.Contains(info, "HTTPSEnable:1") || strings.Contains(info, "SOCKSEnable:1")

	return status, nil
}

func (p *DarwinSystemProxy) TurnOn(addrport string) error {
	host, port, err := net.SplitHostPort(addrport)
	if err != nil {
		return err
	}

	netInterface, err := GetNetworkInterface()
	if err != nil {
		return err
	}

	for _, t := range firmwareType {
		cmd := fmt.Sprintf(`networksetup -set%s "%s" "%s" %s && networksetup -set%sstate "%s" on`, t, netInterface, host, port, t, netInterface)
		if _, err := RunBashCmd(cmd); err != nil {
			return err
		}
	}

	cmd := fmt.Sprintf(`networksetup -setproxybypassdomains "%s" "192.168.0.0/16" "10.0.0.0/8" "172.16.0.0/12" "127.0.0.1" "localhost" "*.local" "timestamp.apple.com"`, netInterface)
	if _, err := RunBashCmd(cmd); err != nil {
		return err
	}

	return nil
}

func NewSystemProxy() SystemProxy {
	return &DarwinSystemProxy{}
}
