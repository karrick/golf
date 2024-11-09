package main

import (
	"fmt"

	"github.com/karrick/golf"
)

func main() {
	// The difference between Type and TypeP is Type accepts a single flag,
	// whereas TypeP accepts two flags, one as a rune and the other as a
	// string.
	//
	// NOTE: When using Type, use double-quotes for the flag name even when
	// providing a single rune. When using TypeP, use single quotes for the
	// run and double-quotes for the long flag name.
	optType := golf.Bool("b", false, "optType takes a single flag and returns a pointer to a variable")
	optTypeP := golf.DurationP('d', "duration", 0, "optTypeP takes a rune and a string and returns a pointer to a variable")

	// The difference between Type and TypeVar is Type returns a pointer to a
	// variable, whereas TypeVar accepts a pointer to a variable.
	var optTypeVar float64
	golf.FloatVar(&optTypeVar, "f", 6.02e-23, "optTypeVar takes a pointer to a variable and a single flag")

	var optTypeVarP int64
	golf.Int64VarP(&optTypeVarP, 'i', "int64", 13, "optTypeVarP takes a pointer to a variable, a rune, and a string")

	fmt.Println("optType: ", *optType)
	fmt.Println("optTypeP: ", *optTypeP)

	fmt.Println("optTypeVar: ", optTypeVar)
	fmt.Println("optTypeVarP: ", optTypeVarP)
}
