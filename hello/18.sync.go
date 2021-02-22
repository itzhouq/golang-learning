// 并发编程
// 协程 goroutine
// 多个并发协程之间不需要通信时可以使用 sync.WaitGroup

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second) // 模拟耗时操作
	wg.Done()               // 减去一个计时
}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)                             // 添加一个计时
		go download("a.com/" + string(i+'0')) // 启动新的协程并发执行download函数
	}
	wg.Wait() // 等待所有的协程执行结束
	fmt.Println("Done")
}

// start to download a.com/2
// start to download a.com/1
// start to download a.com/0
// Done
