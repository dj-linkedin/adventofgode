// cat day3.input | go run day3.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	X int
	Y int
}

func list_of_positions(ops []string) []Pair {
	// first position (the origin)
	x := 0
	y := 0
	var positions []Pair
	positions = append(positions, Pair{x, y})
	for _, op := range ops {
		if 'L' == op[0] {
			value, _ := strconv.Atoi(op[1:])
			for i := 0; i < value; i++ {
				x--
				positions = append(positions, Pair{x, y})
			}
			continue
		}
		if 'R' == op[0] {
			value, _ := strconv.Atoi(op[1:])
			for i := 0; i < value; i++ {
				x++
				positions = append(positions, Pair{x, y})
			}
			continue
		}
		if 'D' == op[0] {
			value, _ := strconv.Atoi(op[1:])
			for i := 0; i < value; i++ {
				y--
				positions = append(positions, Pair{x, y})
			}
			continue
		}
		if 'U' == op[0] {
			value, _ := strconv.Atoi(op[1:])
			for i := 0; i < value; i++ {
				y++
				positions = append(positions, Pair{x, y})
			}
			continue
		}
		panic("invalid direction")
	}
	return positions
}

func manhattan_distance(p1 Pair, p2 Pair) int {
	xdist := p1.X - p2.X
	if xdist < 0 {
		xdist = -1 * xdist
	}
	ydist := p1.Y - p2.Y
	if ydist < 0 {
		ydist = -1 * ydist
	}
	return xdist + ydist
}

func main() {
	var wires = [2][]string{}
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		//		fmt.Println(strings.Split(scanner.Text(), ","))
		wires[count] = strings.Split(scanner.Text(), ",")
		//		fmt.Println(list_of_positions(wires[count]))
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(len(wires))
	p0 := Pair{0, 0}
	for _, p1 := range list_of_positions(wires[0]) {
		for _, p2 := range list_of_positions(wires[1]) {
			if p1 == p2 {
				fmt.Println(manhattan_distance(p0, p1), p1, p2)
			}
		}
	}
}
