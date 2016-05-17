package main

import (
	"fmt"
	"sync"
	"time"
)

func plusOne(n int) {
	time.Sleep(1 * time.Second)
	fmt.Println(n + 1)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			plusOne(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
