package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)

	var wg sync.WaitGroup

	balance := 0
	wg.Add(100)
	var m sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()
			balance += 1
		}()
	}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()
			balance -= 1
		}()
	}

	var counter int64
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {

				atomic.AddInt64(&counter, 1)
				counter++
			}

		}()
	}

	wg.Wait()
	println(balance)
	println(counter)

}
