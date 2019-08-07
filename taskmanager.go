package taskmanager

func TaskManager(tasks []func() error, numInParallel, errLimit int) {
	if numInParallel < 0 || len(tasks) <= 0 {
		return
	}

	taskChan := make(chan func() error)
	errChan := make(chan error)

	//Initialise worker/workers
	go worker(taskChan, errChan)
	if numInParallel > 0 {
		for i := 0; i < numInParallel; i++ {
			go worker(taskChan, errChan)
		}
	}

	// Populate the task channel
	go taskDispatcher(tasks, taskChan, &errLimit)

WORK:
	for {
		if err, ok := <-errChan; ok {
			if err != nil {
				errLimit--
			}
		}
		if errLimit < 0 {
			break WORK
		}
	}
	return
}

func worker(tasksChan chan func() error, errCh chan<- error) {
	for task := range tasksChan {
		errCh <- task()
	}
}

func taskDispatcher(tasks []func() error, taskCh chan func() error, errLim *int) {
	for _, task := range tasks {
		if *errLim >= 0 {
			taskCh <- task
		}
	}
	close(taskCh)
}
