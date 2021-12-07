package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Type Notes:
// new type declaration requires a base type
// for receiver arguments a one or two letter variable is used by convention
