//a go program to illustrate of the Dining Philosophers Problem

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
		fmt.Println("Philosopher", i, " is thinking")
		time.Sleep(time.Second)

		left := fork[i]
		right := fork[(i+1)%(len(fork))]
		left.Lock()
		right.Lock()

		state[i] = 1
		fmt.Println("Philosopher", i, " is eating")
		time.Sleep(time.Second)

		left.Unlock()
		right.Unlock()
		fmt.Println("Philosopher", i, " is finished eating")

		if state[i] == 1 {
			break
		}
	}
}

func main() {
	fmt.Println("enter the no.of philosopher")
	var n int
	fmt.Scanln(&n)
	fork = make([]*sync.Mutex, n)
	for i := 0; i < n; i++ {
		fork[i] = &sync.Mutex{}
	}

	state = make([]int, n)
	for i := 0; i < n; i++ {
		go phil(i)
	}

	for {

	}
}
