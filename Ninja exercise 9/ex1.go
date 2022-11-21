package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("in func main")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("in go routine 1")
		fmt.Println("Goroutines in  go routine 1:", runtime.NumGoroutine())
		wg.Done()
	}()

	// wg.Add(1)
	go func() {
		fmt.Println("in go routine 2")
		fmt.Println("Goroutines in go routine 2", runtime.NumGoroutine())
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Goroutines in main", runtime.NumGoroutine())
	fmt.Println("go routines finished")
}
