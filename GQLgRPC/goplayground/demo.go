package main

import (
	"fmt"
	"time"
)

func server1(option1 *chan string) {
	time.Sleep(6 * time.Second)
	*option1 <- "server1"

}

func server2(option2 *chan string) {
	time.Sleep(3 * time.Second)

	*option2 <- "server2"

}
func main() {
	option1 := make(chan string)
	option2 := make(chan string)

	go server1(&option1)
	go server2(&option2)

	select {
	case op1 := <-option1:
		fmt.Println(op1)
	case op2 := <-option2:
		fmt.Println(op2)
	case <-time.After(5 * time.Second):
		fmt.Println("time exceeded")

	}

}
