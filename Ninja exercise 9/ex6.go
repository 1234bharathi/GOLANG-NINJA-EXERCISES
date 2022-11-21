package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Num of GO ROUTINES", runtime.NumGoroutine())
	fmt.Println("GO ARCH", runtime.GOARCH)
	fmt.Println("OS", runtime.GOOS)

}
