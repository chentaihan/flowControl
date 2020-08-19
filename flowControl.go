package flowControl

import (
	"sync"
	"time"
)

//滑动窗口实现流量控制

type RollWindow struct {
	buf        []int64
	totalQps   int64
	currentQps int64
	index      int
	lastSecond int64
	sync.Mutex
}

//生成一个流量控制对象
//gridCount：滑动窗口格子数量，一个格子就是1秒
//totalQps：滑动窗口总qps，如gridCount=10，totalQps=50，qps = totalQps / gridCount = 5
func NewRollWindow(gridCount int, totalQps int64) *RollWindow {
	return &RollWindow{
		buf:        make([]int64, gridCount),
		totalQps:   totalQps,
		currentQps: 0,
		index:      0,
	}
}

//无锁流量控制，不支持多协程调用
func (fc *RollWindow) Wait() {
	curTime := time.Now().UnixNano()
	curSecond := curTime / int64(time.Second)
	var startValue int64
	if curSecond > fc.lastSecond {
		fc.lastSecond = curSecond
		fc.index++
		if fc.index >= len(fc.buf) {
			fc.index = 0
		}
		startValue = fc.buf[fc.index]
		fc.buf[fc.index] = 0
	}
	fc.buf[fc.index]++
	fc.currentQps++
	fc.currentQps -= startValue
	if fc.currentQps > fc.totalQps {
		//睡到下一秒
		time.Sleep(time.Second - time.Duration(curTime)%time.Second)
	}
}

//有锁流量控制，支持多协程调用
func (fc *RollWindow) WaitLock() {
	fc.Lock()
	fc.Wait()
	fc.Unlock()
}
