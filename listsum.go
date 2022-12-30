package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ListSum(list []int, start int, limit int, c chan int) {
	size := limit - start
	if size == 1 {
		c <- list[start]
	} else {
		sum := 0
		leftchan := make(chan int)
		rightchan := make(chan int)
		go ListSum(list, start, start+size/2, leftchan)
		go ListSum(list, start+size/2, limit, rightchan)
		sum = (<-leftchan + <-rightchan)
		c <- sum
	}

}

func main() {
	slice := []int{}

	const n = 10000000
	for i := 0; i < n; i++ {
		slice = append(slice, rand.Intn(1e3))
	}

	c := make(chan int)

	timeStart := time.Now()
	go ListSum(slice, 0, n, c)
	duration := time.Since(timeStart)
	fmt.Printf("Goroutine sum = %d, time taken = %s\n", <-c, duration)

	timeStart = time.Now()
	seqSum := SeqListSum(slice)
	duration = time.Since(timeStart)
	fmt.Printf("Sequential sum = %d, time taken = %s\n", seqSum, duration)
}

func SeqListSum(list []int) int {
	sum := 0
	for i := 0; i < len(list); i++ {
		sum += list[i]
	}
	return sum
}
