// cat day1.input | go run day1.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fuel(mass int) int {
	intermediate_result := ((mass / 3) - 2)
	if intermediate_result > 0 {
		return intermediate_result
	} else {
		return 0
	}
}

func main() {
	total := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		incremental_fuel := fuel(num)
		for incremental_fuel > 0 {
			total = total + incremental_fuel
			incremental_fuel = fuel(incremental_fuel)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(total)
}
