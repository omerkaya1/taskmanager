package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	taskSlice = []func() error{func() error {
		//time.Sleep(time.Second)
		fmt.Println("First finished")
		return nil
	}, func() error {
		time.Sleep(time.Second * 3)
		fmt.Println("Second finished")
		return nil
	}, func() error {
		time.Sleep(time.Second * 2)
		fmt.Println("Second and half finished")
		return nil
	}, func() error {
		time.Sleep(time.Second * 3)
		fmt.Println("Third finished")
		return nil
	}, func() error {
		time.Sleep(time.Second * 5)
		fmt.Println("Fourth finished")
		return nil
	}, func() error {
		time.Sleep(time.Second * 4)
		fmt.Println("Fifth, with the error")
		return fmt.Errorf("error")
	}, func() error {
		time.Sleep(time.Second * 10)
		fmt.Println("Final, with the error")
		return fmt.Errorf("error")
	}}
)

var wg sync.WaitGroup

// We need to use channels and wait groups!
func DoCool(tasks []func() error, perTime int, errLimit int) {
	if perTime <= 0 {
		return
	}

	taskChan := make(chan func() error, perTime)
	errChan := make(chan error)
	stop := make(chan struct{})

	go fillTasks(tasks, taskChan)

	wg.Add(1)
	go errs(taskChan, errChan)

	// Main routine
Label:
	for {
		select {
		case logErr := <-errChan:
			fmt.Printf("Error: %#v; Left: %d\n", logErr, errLimit)
			if logErr != nil {
				fmt.Printf("Error: %#v; Left: %d\n", logErr, errLimit)
				errLimit--
			}
			if errLimit == 0 {
				fmt.Printf("Or here?\n")
				stop <- struct{}{}
				break Label
			}
		case <-stop:
			break Label
		}
	}
	fmt.Printf("Done!\n")
	wg.Done()
	return
}

func fillTasks(innerTasks []func() error, taskChannel chan func() error) {
	for _, t := range innerTasks {
		taskChannel <- t
	}
	close(taskChannel)
	return
}

func errs(tasks chan func() error, errors chan error) {
	for t := range tasks {
		errors <- t()
	}
	wg.Done()
	close(errors)
}

func main() {
	wg.Add(1)
	go DoCool(taskSlice, 2, 1)
	wg.Wait()
	return
}
