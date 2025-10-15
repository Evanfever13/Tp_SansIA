package main

import (
	"fmt"
)

type Package struct {
	Number    int
	Recipient string
	Address   string
	Content   string
	Weight    float64
}

type Wagon struct {
	Number    string
	Loading   []Package
	NextWagon *Wagon
}

type ResultSearch struct {
	NumberWagon     string
	PositionPackage int
	PackageFind     Package
}

func main() {
	// Exemple de données
	w1 := &Wagon{
		Number:  "A",
		Loading: []Package{{Number: 10}, {Number: 11}},
	}
	w2 := &Wagon{
		Number:  "B",
		Loading: []Package{{Number: 12}, {Number: 13}, {Number: 42}},
	}

	w1.NextWagon = w2

	fmt.Println(FindPackage(42, w1))
	// Résultat attendu : Wagon B, position 2, package 42
}

func FindPackage(number int, startWagon *Wagon) ResultSearch {
	for i := 0; i <= len(startWagon.Loading)-1; i++ {
		if startWagon.Loading[i].Number == number {
			result := ResultSearch{startWagon.Number, i, startWagon.Loading[i]}
			return result
		}
	}
	return FindPackage(number, startWagon.NextWagon)
}
