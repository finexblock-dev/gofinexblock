package daemon

func Run(daemon Daemon) {
	var err error
	daemon.run()

	for {
		if daemon.state() == Running {
			if err = Task(daemon); err != nil {
				if nextErr := daemon.insertErrLog(err); nextErr != nil {
					daemon.log(nextErr)
					Stop(daemon)
				}
			}
		}
		Sleep(daemon)
	}
}

func Stop(daemon Daemon) {
	daemon.stop()
}

func Sleep(daemon Daemon) {
	daemon.sleep()
}

func Task(daemon Daemon) error {
	return daemon.task()
}
