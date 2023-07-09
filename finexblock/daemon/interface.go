package daemon

type Daemon interface {
	run()
	sleep()
	stop()
	task() error
	state() State
	setState(State)
	insertErrLog(err error) error
	log(v ...any)
}

type State int

const (
	Stopped State = iota
	Running
)
