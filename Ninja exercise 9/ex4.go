package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("No of runtimes at function start", runtime.NumGoroutine())
	var wg sync.WaitGroup
	const gs = 25
	counter := 0
	wg.Add(gs)
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter

			v++
			counter = v

			fmt.Println("No of counter", counter)
			mu.Unlock()
			wg.Done()
			fmt.Println("No of runtimes mid", runtime.NumGoroutine())
		}()

	}
	wg.Wait()
	fmt.Println("No of runtimes end", runtime.NumGoroutine())
}
