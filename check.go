// a go program to illustrate of Checkpoint Synchronization
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	fork  []*sync.Mutex
	state []int
)

func phil(i int) {
	for {
		state[i] = 0
		fmt.Println("Phil", i, " is thinking")
		time.Sleep(time.Second)

		left := fork[i]
		right := fork[(i+1)%(len(fork))]
		left.Lock()
		right.Lock()

		state[i] = 1
		fmt.Println("Phil", i, " is eating")
		time.Sleep(time.Second)
		left.Unlock()
		right.Unlock()

		state[i] = 2
		fmt.Println("Phil", i, " has reached the check point")
		time.Sleep(time.Second)
		if state[i] == 2 {
			break
		}
	}
}

func main() {
	var n int
	fmt.Println("Enter the no.of phil")
	fmt.Scanln(&n)
	fork = make([]*sync.Mutex, n)
	state = make([]int, n)
	for i := 0; i < n; i++ {
		fork[i] = &sync.Mutex{}
	}
	for i := 0; i < n; i++ {
		go phil(i)
	}
	for {

	}
}
