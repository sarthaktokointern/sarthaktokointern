package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func SequentialADD(numbers []int) int {
	var sum int
	sum = 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func ConcurrentADD(numbers []int) int {
	numOfCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numOfCores)

	max := len(numbers)

	lengthOfEachPart := max / numOfCores
	var sum int
	sum = 0
	var wg sync.WaitGroup
	wg.Add(numOfCores)
	for i := 0; i < numOfCores; i++ {

		st := i * lengthOfEachPart
		en := st + lengthOfEachPart
		part := numbers[st:en]

		go func(nums []int) {
			var partsum int
			partsum = 0
			for _, num := range nums {
				partsum += num
			}
			sum += partsum
			wg.Done()
		}(part)

	}

	wg.Wait()

	return sum
}

func main() {
	numbers := rand.Perm(1e7)
	st := time.Now()
	sum1 := SequentialADD(numbers)
	fmt.Printf("the sum is : %d and time taken is %d\n", sum1, time.Since(st).Milliseconds())
	st = time.Now()
	sum2 := ConcurrentADD(numbers)
	fmt.Printf("the sum is : %d and the time taken is %d\n", sum2, time.Since(st).Milliseconds())

}
