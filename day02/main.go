package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("day02/input.txt")
	ss := strings.Split(string(f), "\n")

	// Part 1
	h := 0
	d := 0
	for _, s := range ss {
		command := strings.Split(s, " ")
		if len(command) == 2 {
			val, _ := strconv.Atoi(command[1])
			switch command[0] {
			case "forward":
				h += val
			case "down":
				d += val
			case "up":
				d -= val
			}
		}
	}
	fmt.Println("Part 1 =", h*d)

	// Part 2
	h = 0
	d = 0
	a := 0
	for _, s := range ss {
		command := strings.Split(s, " ")
		if len(command) == 2 {
			val, _ := strconv.Atoi(command[1])
			switch command[0] {
			case "forward":
				h += val
				d += a * val
			case "down":
				a += val
			case "up":
				a -= val
			}
		}
	}
	fmt.Println("Part 2 =", h*d)
}
