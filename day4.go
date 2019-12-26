package main

import (
	"fmt"
)

func adjacent_digits(n int) bool {
	p := -1
	for n > 0 {
		if p == n%10 {
			return true
		}
		p = n % 10
		n = n / 10
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
