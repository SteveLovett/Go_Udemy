package main

func main() {

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

// General Notes:
// unused variables will throw a compiler error - .\main.go:10:6: <var> declared but not used
//
// Variable Notes:
// variables can be declared outside of functions but cannot be initialized
// var card string - longform variable declaration
// card := "Ace of Spades" - shortform declaration
// := is only used on declaration
// = used to reassign after declaration
//
// Function Notes:
// default function return type is void
// function return type declared after list of parameters
// functions can be used in other files using the same package without an import statement
// 		^cont. - but VS Code is throwing an error even though it compiled and ran properly
//
// Array Notes:
// two types of arrays Array and Slice - Array is fixed length and Slice is not fixed
// slices are declared like other variables but with []<type>{val1,val2,...} after the :=
// append slices by var = append(var, "value")
// append returns a new slice, that is why the variable reassignment is necessary
//
// Loop Notes:
// only one loop in Go - for loop
// variables declared in the loop are local and destroyed at the end of the loop
// for i, x := range <array/slice> {...} - valid for loop - seems to be shorthand?
// for i := 0; i < 10; i++ {...} - valid for loop - longhand?
// parentheses not required, curly braced are required
