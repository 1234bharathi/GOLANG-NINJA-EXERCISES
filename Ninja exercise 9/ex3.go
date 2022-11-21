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

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("No of counter", counter)
			wg.Done()
			fmt.Println("No of runtimes mid", runtime.NumGoroutine())
		}()

	}
	wg.Wait()
	fmt.Println("No of runtimes end", runtime.NumGoroutine())
}
