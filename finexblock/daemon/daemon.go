package daemon

func Run(daemon Daemon) {
	var err error
	daemon.Run()

	for {
		if daemon.State() == Running {
			if err = Task(daemon); err != nil {
				if nextErr := daemon.InsertErrLog(err); nextErr != nil {
					daemon.Log(nextErr)
					Stop(daemon)
				}
			}
		}
		Sleep(daemon)
	}
}

func Stop(daemon Daemon) {
	daemon.Stop()
}

func Sleep(daemon Daemon) {
	daemon.Sleep()
}

func Task(daemon Daemon) error {
	return daemon.Task()
}
