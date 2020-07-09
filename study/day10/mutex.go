package main

import (
"fmt"
"sync"
"time"
)

func main() {
	mutex := make([]sync.Mutex,10)

	go func() {
		mutex[1].Lock()
		defer mutex[1].Unlock()
		fmt.Println("i sleep")
		time.Sleep(time.Second*3)
		fmt.Println("i sleep done")
	}()

	go func() {
		fmt.Println("second goroutine")
		mutex[1].Lock()
		defer mutex[1].Unlock()
		fmt.Println("second go routine end")
	}()

	time.Sleep(time.Second*5)
}
