package sysproxy

type SystemProxyStatus struct {
	State bool `json:"state"`
}

type SystemProxy interface {
	TurnOn(addrPort string) error
	TurnOff() error

	Status() (*SystemProxyStatus, error)
}
