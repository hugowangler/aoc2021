package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("day01/input.txt")
	ss := strings.Split(string(data), "\n")

	// Part 1
	inc := 0
	for i := 1; i < len(ss); i++ {
		earlier, _ := strconv.Atoi(ss[i-1])
		now, _ := strconv.Atoi(ss[i])
		if now > earlier {
			inc++
		}
	}
	fmt.Println("Part 1 =", inc)

	// Part 2
	inc = 0
	for i := 0; i < len(ss)-3; i++ {
		e := ss[i : i+3]
		n := ss[i+1 : i+4]
		eSum := 0
		nSum := 0
		for j := 0; j < len(e); j++ {
			x, _ := strconv.Atoi(e[j])
			y, _ := strconv.Atoi(n[j])
			eSum += x
			nSum += y
		}
		if nSum > eSum {
			inc++
		}
	}
	fmt.Println("Part 2 =", inc)
}
