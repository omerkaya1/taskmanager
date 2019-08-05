package taskmanager

func TaskManager(tasks []func() error, numInParallel, errLimit int) {
	if numInParallel < 0 || len(tasks) <= 0 {
		return
	}

	taskChan := make(chan func() error, len(tasks))
	errChan := make(chan error)
	doneChan := make(chan bool)
	toFinish, errCount := 0, 0
	//defer close(errChan)

	// Initialise worker/workers
	if numInParallel == 0 {
		go worker(taskChan, doneChan, errChan)
		toFinish++
	} else {
		for i := 0; i < numInParallel; i++ {
			go worker(taskChan, doneChan, errChan)
			toFinish++
		}
	}

	// Populate the task channel and then close it
	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan)

	// Main work cycle, where we wait for all the workers to finish,
	// simultaneously listening for what we get from the errChan
WORK:
	for {
		select {
		case err := <-errChan:
			if err != nil {
				errCount++
			}
			if errLimit < errCount {
				break WORK
			}
		case <-doneChan:
			if toFinish--; toFinish == 0 {
				break WORK
			}
		}
	}
	return
}

func worker(tasksChan chan func() error, doneCh chan bool, errCh chan<- error) {
	for task := range tasksChan {
		errCh <- task()
	}
	doneCh <- true
	return
}
