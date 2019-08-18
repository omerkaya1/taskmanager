package taskmanager

func TaskManager(tasks []func() error, numInParallel, errLimit int) {
	if numInParallel < 0 || len(tasks) <= 0 {
		return
	}
	// Workers will be communicating with the main routine via task and error channels
	taskChan := make(chan func() error)
	errChan := make(chan error)
	// There is also a readyChan which ensures that workers are run only when they are ready
	var readyChan chan bool
	if numInParallel == 0 {
		readyChan = make(chan bool, 1)
		go worker(taskChan, errChan)
		readyChan <- true
	} else {
		readyChan = make(chan bool, numInParallel)
	}
	//Initialise worker/workers
	for i := 0; i < numInParallel; i++ {
		go worker(taskChan, errChan)
		readyChan <- true
	}

WORK:
	for i := 0; i < len(tasks); {
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
