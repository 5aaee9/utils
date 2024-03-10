package sysproxy

type SystemProxy interface {
	TurnOn(addrPort string) error
	TurnOff() error
}
