package main

import "time"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		//time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- "two"
	}()

	for i := 0; i < 4; i++ {
		select {
		case val := <-ch1:
			println(val)
		case val := <-ch2:
			println(val)
		case <-time.After(1 * time.Second):
			println("timeout")
		default:
			println("continue!!")

		}
	}

}
