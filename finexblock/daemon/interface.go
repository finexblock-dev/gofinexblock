package daemon

type Daemon interface {
	Run()
	Sleep()
	Stop()
	Task() error
}

type State int

const (
	Stopped State = iota
	Running
)