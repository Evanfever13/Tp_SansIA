package main

import "fmt"

func CompressRLE(s string) string {
	if len(s) == 0 {
		return ""
	} else {
		var resultat string = ""
		var compteur int = 1
		for i := 1; i < len(s); i++ {
			if s[i] == s[i-1] {
				compteur++
			} else {
				resultat += fmt.Sprintf("%c%d", s[i-1], compteur)
				compteur = 1
			}
		}
		resultat += fmt.Sprintf("%c%d", s[len(s)-1], compteur)
		return resultat
	}
}

func main() {
	fmt.Println(CompressRLE("aaabbc")) // a3b2c1
	fmt.Println(CompressRLE("ab"))     // a1b1
}
