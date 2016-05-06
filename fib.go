package main

import "fmt"

func main() {
	dp := newdp()
	fmt.Println(dp.fib(500))
}

// dp is a dynamic prog wrapper that uses memoization technique to improve
// T(n) from exponential to polynomial
type dp struct {
	memo map[int]int
}

func newdp() *dp {
	return &dp{memo: make(map[int]int)}
}

func (d *dp) fib(n int) int {
	if _, ok := d.memo[n]; ok {
		return d.memo[n]
	}

	if n <= 1 {
		return 1
	}
	v := d.fib(n-1) + d.fib(n-2)
	d.memo[n] = v
	return v
}

func (d *dp) fib_raw(n int) int {
	if n <= 1 {
		return 1
	}
	return d.fib_raw(n-1) + d.fib_raw(n-2)
}
