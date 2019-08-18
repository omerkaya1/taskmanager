package taskmanager

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	taskSlice = []func() error{func() error {
		time.Sleep(time.Millisecond)
		fmt.Println("First finished")
		return nil

	}, func() error {
		time.Sleep(time.Millisecond)
		fmt.Println("Second finished")
		return nil

	}, func() error {
		time.Sleep(time.Millisecond)
		fmt.Println("Third finished")
		return nil

	}, func() error {
		time.Sleep(time.Millisecond * 3)
		fmt.Println("Fourth finished")
		return fmt.Errorf("first error")

	}, func() error {
		time.Sleep(time.Millisecond * 6)
		fmt.Println("Fifth finished")
		return nil

	}, func() error {
		time.Sleep(time.Millisecond * 5)
		fmt.Println("Sixth finished with an error")
		return fmt.Errorf("fourth error")

	}, func() error {
		time.Sleep(time.Millisecond * 7)
		fmt.Println("Seventh finished with an error")
		return fmt.Errorf("fifth error")

	}, func() error {
		time.Sleep(time.Millisecond * 8)
		fmt.Println("Semi-final")
		return nil

	}, func() error {
		time.Sleep(time.Millisecond * 10)
		fmt.Println("Almost final")
		return nil
	}, func() error {
		time.Sleep(time.Millisecond * 9)
		fmt.Println("Final, with an error")
		return fmt.Errorf("yet another error")

	}}
)

func TestWorker_ZeroErrorTolerance(t *testing.T) {
	errNum := 0
	start := time.Now()
	TaskManager(taskSlice, 0, errNum)
	stop := time.Now()
	assert.WithinDuration(t, stop, start, time.Millisecond*10)
}

func TestWorker_TwoErrorTolerance(t *testing.T) {
	errNum := 2
	start := time.Now()
	TaskManager(taskSlice, 5, errNum)
	stop := time.Now()
	assert.WithinDuration(t, stop, start, time.Millisecond*15)
}

func TestWorker_FullErrorTolerance(t *testing.T) {
	errNum := 4
	start := time.Now()
	TaskManager(taskSlice, 5, errNum)
	stop := time.Now()
	assert.WithinDuration(t, stop, start, time.Millisecond*12)
}
