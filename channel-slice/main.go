package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := []int{}
	ch := make(chan int)
	var wg sync.WaitGroup

	go func() {
		for val := range ch {
			slice = append(slice, val)
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			ch <- val
		}(i)
	}
	wg.Wait()
	close(ch)

	fmt.Println("Final slice:", slice)
}
