package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("No of runtimes at function start", runtime.NumGoroutine())
	const gs = 25
	var wg sync.WaitGroup
	var counter int64
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			// reads the counter variable
			val := atomic.LoadInt64(&counter)
			fmt.Println("No of counter", val)

			wg.Done()
			fmt.Println("No of runtimes mid", runtime.NumGoroutine())
		}()

	}
	wg.Wait()
	fmt.Println("No of runtimes end", runtime.NumGoroutine())
}
