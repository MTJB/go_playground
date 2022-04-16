package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	var num int

	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&num)
	fibonacciForLoop(num)
	fibonacciNoTemp(num)
	fibonacciResursive(num)
}

func fibonacciForLoop(num int) {
	t1 := 0
	t2 := 1
	nextTerm := 0
	defer timeTrack(time.Now(), "For Loop")
	// fmt.Print("Fibonacci Series:")
	for i := 1; i <= num; i++ {
		if i == 1 {
			// fmt.Print(" ", t1)
			continue
		}

		if i == 2 {
			// fmt.Print(" ", t2)
			continue
		}

		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
		// fmt.Print(" ", nextTerm)
	}
	// fmt.Println()
}

func fibonacciNoTemp(num int) {
	t1 := 0
	t2 := 1
	defer timeTrack(time.Now(), "No temp variable")
	// fmt.Print("Fibonacci Series:")

	for i := 1; i <= num; i++ {
		if i == 1 {
			// fmt.Print(" ", t1)
			continue
		}

		if i == 2 {
			// fmt.Print(" ", t2)
			continue
		}
		// fmt.Print(" ", t1+t2)
		t2 = t1 + t2
		t1 = t2 - t1
	}
	// fmt.Println()
}

func fibonacciResursive(num int) {
	defer timeTrack(time.Now(), "Recursion")
	// fmt.Print("Fibonacci Series: ")
	for i := 0; i < num; i++ {
		fib(i)
		// fmt.Print(" ", fib(i))
	}
	// fmt.Println()
}

func fib(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fib(num-1) + fib(num-2)
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
