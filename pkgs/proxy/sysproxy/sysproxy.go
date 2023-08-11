package sysproxy

type SystemProxy interface {
	TurnOn(addrport string) error
	TurnOff() error
}
