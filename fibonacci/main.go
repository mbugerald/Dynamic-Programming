package main

import "log"

func main() {

	memo := map[int]int{}

	// log.Println(fib(50))
	log.Println(fibMemoized(50, memo))

}

// ordinary not memoized.
func fib(n int) int {

	if n <= 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

// memoized
func fibMemoized(n int, memo map[int]int) int {

	if _, ok := memo[n]; ok {
		return memo[n]
	}

	if n <= 2 {
		return 1
	}

	memo[n] = fibMemoized(n-1, memo) + fibMemoized(n-2, memo)
	return memo[n]
}
