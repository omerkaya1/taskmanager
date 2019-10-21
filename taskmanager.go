package taskmanager

func TaskManager(tasks []func() error, numInParallel, errLimit int) {
	if numInParallel < 0 || len(tasks) <= 0 {
		return
	}
	// Workers will be communicating with the main routine via the task and error channels
	taskChan := make(chan func() error)
	errChan := make(chan error, 1)
	//Initialise worker/workers
	for i := 0; i <= numInParallel; i++ {
		go worker(taskChan, errChan)
	}

WORK:
	for i := 0; i < len(tasks); {
		select {
		case err := <-errChan:
			if err != nil {
				errLimit--
			}
		default:
			if errLimit < 0 {
				close(taskChan)
				break WORK
			}
			taskChan <- tasks[i]
			i++
		}
	}
	return
}

func worker(tasksChan chan func() error, errCh chan<- error) {
	for task := range tasksChan {
		errCh <- task()
	}
}
