package main

import "fmt"

func PrintHollowDiamond(r int) {
	nbr := r
	if nbr%2 == 0 {
		nbr++
	}
	espace := nbr / 2
	for toto := 0; toto < nbr; toto++ {
		for tata := 0; tata < nbr; tata++ {
			if toto+tata == espace || tata-toto == espace || toto-tata == espace || toto+tata == nbr+espace-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	PrintHollowDiamond(4)
}
