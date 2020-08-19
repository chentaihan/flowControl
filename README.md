# flowControl

## 流量控制器
### 滑动窗口实现流量控制

```golang
//qps = 1000/5 = 200
rw := NewRollWindow(5, 1000)
startTime := time.Now()
const count = 10000
for i := 0; i < count; i++ {
    rw.Wait()
}
useTime := time.Now().Sub(startTime)
fmt.Println(useTime)
```