package taskmanager

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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
		time.Sleep(time.Second * 3)
		fmt.Println("Fourth finished")
		return nil
	}, func() error {
		time.Sleep(time.Second * 6)
		fmt.Println("Fifth finished")
		return nil
	}, func() error {
		time.Sleep(time.Second * 5)
		fmt.Println("Sixth finished with an error")
		return fmt.Errorf("fourth error")
	}, func() error {
		time.Sleep(time.Second * 7)
		fmt.Println("Seventh finished with an error")
		return fmt.Errorf("fifth error")
	}, func() error {
		time.Sleep(time.Second * 8)
		fmt.Println("Semi-final")
		return nil
	}, func() error {
		time.Sleep(time.Second * 10)
		fmt.Println("Almost final, with an error")
		return nil
	}, func() error {
		time.Sleep(time.Second * 10)
		fmt.Println("Final, with an error")
		return fmt.Errorf("yet another error")
	}}
)

func TestWorker_ZeroErrorTolerance(t *testing.T) {
	errNum := 0
	test := make(chan int, errNum)
	Worker(taskSlice[4:], 5, 0, test)
	assert.NoError(t, nil)

}

func TestWorker_OneErrorTolerance(t *testing.T) {
	errNum := 2
	test := make(chan int, errNum)
	Worker(taskSlice, 5, errNum, test)
}
