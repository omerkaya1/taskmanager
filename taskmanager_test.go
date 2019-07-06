package main

import (
	"testing"
)

//var (
//	taskSlice = []func() error { func() error {
//		//time.Sleep(time.Second)
//		fmt.Println("First finished")
//		return nil
//	}, func() error {
//		time.Sleep(time.Second * 3)
//		fmt.Println("Second finished")
//		return nil
//	}, func() error {
//		time.Sleep(time.Second * 2)
//		fmt.Println("Second and half finished")
//		return nil
//	}, func() error {
//		time.Sleep(time.Second * 3)
//		fmt.Println("Third finished")
//		return nil
//	}, func() error {
//		time.Sleep(time.Second * 5)
//		fmt.Println("Fourth finished")
//		return nil
//	}, func() error {
//		time.Sleep(time.Second * 4)
//		fmt.Println("Fifth, with the error")
//		return fmt.Errorf("error")
//	}, func() error {
//		time.Sleep(time.Second * 10)
//		fmt.Println("Final, with the error")
//		return fmt.Errorf("error")
//	}}
//)

func TestTaskManager_DoCool(t *testing.T) {
	DoCool(taskSlice, 3, 1)
}
