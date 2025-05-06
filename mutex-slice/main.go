package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		slice []int
		mu    sync.Mutex
		wg    sync.WaitGroup
	)

	// 每个 goroutine 在 append 前先加锁，操作完成后释放锁。
	// sync.Mutex 保证了 append 的原子性，避免了竞态。
	// 	容易死锁，复杂操作难管理
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()

			mu.Lock()
			slice = append(slice, val)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println("Final slice:", slice)
}
