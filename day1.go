// cat day1.input | go run day1.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fuel(mass int) int {
	return ((mass / 3) - 2)
}

func main() {
	total := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		total = total + fuel(num)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(total)
}
