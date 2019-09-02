package taskmanager

func TaskManager(tasks []func() error, numInParallel, errLimit int) {
	if numInParallel < 0 || len(tasks) <= 0 {
		return
	}
	// Workers will be communicating with the main routine via the task and error channels
	taskChan := make(chan func() error)
	errChan := make(chan error)
	//Initialise worker/workers
	for i := 0; i <= numInParallel; i++ {
		go worker(taskChan, errChan)
	}

	// Tasks are added to the task channel in a separate goroutine
	go func() {
		for _, t := range tasks {
			taskChan <- t
		}
	}()

WORK:
	for i := 0; i < len(tasks); {
		select {
		case e := <-errChan:
			i++
			if e != nil {
				errLimit--
			}
			if errLimit < 0 {
				break WORK
			}
		default:
			break
		}
	}
	return
}

func worker(tasksChan chan func() error, errCh chan<- error) {
	for task := range tasksChan {
		errCh <- task()
	}
}
