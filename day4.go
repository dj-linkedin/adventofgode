package main

import (
	"fmt"
)

func adjacent_digits(n int) bool {
	run_length := 1
	p := -1
	for n > 0 {
		if p == n%10 {
			run_length++
		} else {
			if 2 == run_length {
				return true
			} else {
				run_length = 1
			}
		}
		p = n % 10
		n = n / 10
	}
	if 2 == run_length {
		return true
	}
	return false
}

func not_decreasing(n int) bool {
	p := 10
	for n > 0 {
		if p < n%10 {
			return false
		}
		p = n % 10
		n = n / 10
	}
	return true
}

func main() {
	lo := 145852
	hi := 616942
	for i := lo; i <= hi; i++ {
		if adjacent_digits(i) && not_decreasing(i) {
			fmt.Println(i)
		}
	}
}
