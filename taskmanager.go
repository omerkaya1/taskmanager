package taskmanager

func Worker(tasks []func() error, perTime, errLimit int, report chan<- int) {
	if perTime <= 0 || len(tasks) == 0 {
		return
	}

	taskChan := make(chan func() error, perTime)
	errChan := make(chan error)

	go enqueue(tasks, taskChan)
	go manager(taskChan, errChan)

	// The loop is used to handle cases when we either don't have errors or
	// the allowed number of errors is equal to the number of
	// errors encountered during the execution
	for i := 0; i < len(tasks); i++ {
		if e := <-errChan; e != nil {
			report <- errLimit
			errLimit--
		}
		if errLimit <= 0 {
			return
		}
	}
}

func enqueue(taskSlice []func() error, taskCh chan func() error) {
	for _, t := range taskSlice {
		taskCh <- t
	}
}

func manager(tasks chan func() error, errCh chan error) {
	defer close(errCh)
	// Do the tasks concurrently; the range will take care of the taskChan closure
	for t := range tasks {
		go func(task func() error, errChannel chan error) {
			if e := task(); e != nil {
				errChannel <- e
			}
		}(t, errCh)
	}
}
