package main

import "fmt"

func main() {
	PrintCenteredPyramid(4)
}

func PrintCenteredPyramid(n int) {
	ns := n
	for i := 0; i < n; i++ {
		ns -= 1
		fmt.Println(Space(ns) + Star(i))
	}
}

func Space(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += " "
	}
	return s
}

func Star(n int) string {
	c := "*"
	for i := 0; i < n; i++ {
		c += "**"
	}
	return c
}
