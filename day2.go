// cat day2.input | go run day2.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func str_array_to_int_array(string_array []string) []int {
	var i_a = []int{}
	for _, s := range string_array {
		i, _ := strconv.Atoi(s)
		i_a = append(i_a, i)
	}
	return i_a
}

func int_array_to_str_array(int_array []int) []string {
	var s_a = []string{}
	for _, i := range int_array {
		s := strconv.Itoa(i)
		s_a = append(s_a, s)
	}
	return s_a
}

func process(i_a []int) []int {
	cur_pos := 0
	for {
		cur_val := i_a[cur_pos]
		if 99 == cur_val {
			return i_a
		}
		if 1 == cur_val {
			i_a[i_a[cur_pos+3]] =
				i_a[i_a[cur_pos+1]] + i_a[i_a[cur_pos+2]]
			cur_pos += 4
			continue
		}
		if 2 == cur_val {
			i_a[i_a[cur_pos+3]] =
				i_a[i_a[cur_pos+1]] * i_a[i_a[cur_pos+2]]
			cur_pos += 4
			continue
		}
		fmt.Println(cur_val)
		panic("error - invalid opcode")
	}
	return i_a
}

func main() {
	var program = []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		program = strings.Split(scanner.Text(), ",")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			values := program
			values[1] = strconv.Itoa(noun)
			values[2] = strconv.Itoa(verb)
			fmt.Println(noun, verb, process(str_array_to_int_array(values))[0])
		}
	}
}
