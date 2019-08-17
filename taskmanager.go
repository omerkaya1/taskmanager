package taskmanager

func TaskManager(tasks []func() error, numInParallel, errLimit int) {
	if numInParallel < 0 || len(tasks) <= 0 {
		return
	}

	taskChan := make(chan func() error)
	errChan := make(chan error)
	readyChan := make(chan bool, numInParallel)

	//Initialise worker/workers
	for i := 0; i < numInParallel; i++ {
		go worker(taskChan, errChan)
		readyChan <- true
	}

WORK:
	for i := 0; i < len(tasks); i++ {
		select {
		case e := <-errChan:
			if e != nil {
				errLimit--
			}
			readyChan <- true
		case <-readyChan:
			if errLimit < 0 {
				break WORK
			}
			taskChan <- tasks[i]
		}
	}
	return
}

func worker(tasksChan chan func() error, errCh chan<- error) {
	for task := range tasksChan {
		errCh <- task()
	}
}
