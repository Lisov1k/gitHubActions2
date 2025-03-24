package main

import (
	"fmt"
	"sync"
)

func main() {

}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func Race() [10]int {
	counter := [10]int{}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			counter[i] = i
			wg.Done()
		}()
	}

	wg.Wait()
	return counter
}
