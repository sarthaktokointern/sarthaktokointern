package main

import (
	"sync"
	"time"
)

func print(s string) {
	for i := 0; i < 3; i++ {
		println(s)
		time.Sleep(1 * time.Millisecond)
	}
}
func main() {
	//direct call
	print("direct call")
	//go routine function call
	go print("goroutine 1")
	//goroutine with anonymous function call
	go func() {
		print("goroutine 2")
	}()
	//goroutine with function value call
	go func(s string) {
		go print(s)
	}("goroutine 3")
	time.Sleep(3 * time.Millisecond)
	//waitgroups in action
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			println(i)
		}(i)
	}
	wg.Wait()
}
