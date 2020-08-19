package flowControl

import (
	"sync"
	"testing"
	"time"
)

func TestRollWindow_Wait(t *testing.T) {
	rw := NewRollWindow(5, 1000)
	startTime := time.Now()
	const count = 10000
	for i := 0; i < count; i++ {
		rw.Wait()
	}
	useTime := time.Now().Sub(startTime)
	t.Log(useTime)
}

func TestRollWindow_WaitLock(t *testing.T) {
	rw := NewRollWindow(5, 1000)
	var wg sync.WaitGroup
	startTime := time.Now()
	const count = 10000
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < count/10; j++ {
				rw.Wait()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	useTime := time.Now().Sub(startTime)
	t.Log(useTime)
}
