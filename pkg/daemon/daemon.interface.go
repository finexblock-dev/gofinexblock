package daemon

type Daemon interface {
	Run()
	Sleep()
	Stop()
	Task() error
	State() State
	SetState(State)
	InsertErrLog(err error) error
	Log(v ...any)
}

type State int

const (
	Stopped State = iota
	Running
)
