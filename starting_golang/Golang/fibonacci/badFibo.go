package main

import (
	"fmt"
	"time"
)

/*
Function to calculate fib to m.
*/
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2) // very very slow
}

/*
Function to print a loading visual indication that the program is still running.
*/
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func main() {
	go spinner(100 * time.Millisecond) // create a ne goroutine to execute spinner()
	const n = 45
	fibN := fib(n) // executed by main
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
