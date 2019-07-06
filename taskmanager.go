package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	taskSlice = []func() error{func() error {
		time.Sleep(time.Second)
		fmt.Println("First finished")
		return nil
	}, func() error {
		time.Sleep(time.Second)
		fmt.Println("Second finished")
		return nil
	}, func() error {
		time.Sleep(time.Second)
		fmt.Println("Third finished")
		return nil
	}, func() error {
		time.Sleep(time.Second)
		fmt.Println("Third and half finished")
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
		//return fmt.Errorf("error")
		return nil
	}, func() error {
		time.Sleep(time.Second * 10)
		fmt.Println("Final, with the error")
		return fmt.Errorf("error")
	}}
)

// We need to use channels and wait groups!
func DoCool(tasks []func() error, perTime, errLimit int) {
	if perTime <= 0 {
		return
	}
	wg := sync.WaitGroup{}
	taskChan := make(chan func() error, perTime)
	errChan := make(chan error, perTime)

	go func(taskSlice []func() error, errCh chan func() error) {
		for _, t := range taskSlice {
			errCh <- t
		}
	}(tasks, taskChan)

LABEL:
	for {
		select {
		case t := <-taskChan:
			wg.Add(1)
			go func(task func() error, errCh chan error) {
				defer wg.Done()
				errCh <- t()
			}(t, errChan)
			break
		case e := <-errChan:
			if e != nil {
				fmt.Printf(" error: %v\n", e)
				errLimit--
			}
			fmt.Printf("Limit: %d\n", errLimit)
			if errLimit == 0 {
				fmt.Printf("reached the limit!\n")
				close(taskChan)
				close(errChan)
				break LABEL
			}
			break
		}
	}
	return
}

func main() {
	DoCool(taskSlice, 4, 1)
	return
}
