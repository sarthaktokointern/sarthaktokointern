package main

import "sync"

func main() {
	ch := make(chan int)
	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)
	val, ok := <-ch
	if ok {
		println(val)

	}
	ch2 := make(chan int, 6)
	var wg sync.WaitGroup

	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch2 <- i

		}(i)
	}
	wg.Wait()
	close(ch2)
	for val := range ch2 {
		println(val)
	}

	owner := func() <-chan int {
		ch4 := make(chan int, 6)
		go func() {
			for i := 0; i < 6; i++ {
				ch4 <- i
			}
			close(ch4)
		}()
		return ch4

	}
	consumer := func(ch5 <-chan int) {
		for val := range ch5 {
			println(val)
		}

	}

	ch3 := owner()
	consumer(ch3)

}
