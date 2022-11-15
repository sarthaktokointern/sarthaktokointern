package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Mutex{}
	c := sync.NewCond(&m)
	mp := make(map[string]string)

	go func() {
		c.L.Lock()

		for len(mp) == 0 {
			c.Wait()
		}
	}()

	go func() {
		c.L.Lock()

		for len(mp) < 2 {
			c.Wait()
		}
	}()
	c.L.Lock()
	mp["name"] = "sarthak"
	mp["age"] = "22"
	c.Broadcast()
	c.L.Unlock()

	fmt.Printf("name of the user : %s\nage of the user: %s\n", mp["name"], mp["age"])

	var wg sync.WaitGroup

	load := func() {
		print("loaded only once!!!!")
	}
	var once sync.Once
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(load)
		}()
	}
	wg.Wait()
}
