package main

import (
	"fmt"

	"github.com/karrick/golf"
)

func main() {
	// 'Type' functions are exactly like their corresponding functions from
	// the 'flag' standard library. However the length of the argument string
	// determines whether a single or double-hyphen prefix is used to change
	// its value. When this argument is a single rune long string then a
	// single-hyphen prefix is used to change its value, for example, "V" will
	// configure the parser to associate '-V' with the option. Otherwise a
	// double-hyphen prefix is used to change its value, for example,
	// "version" will configure the parser to associate '--version' with the
	// option.
	optType1 := golf.Bool("b", false, "can take a one rune flag name and return a variable pointer")
	otpType2 := golf.String("id", "", "can also take a multiple rune flag name and return a variable pointer")

	// 'TypeP' functions take two arguments. The first argument is a rune to
	// be associated with a single-hyphen prefix, for example, '-V'. The
	// second argument is a string to be associated with a double-hyphen
	// prefix, for example, '--version'.
	optTypeP := golf.DurationP('d', "duration", 0, "takes a one rune and a multiple rune flag name and returns a pointer to a variable")

	// The difference between 'Type' and 'TypeVar' is 'Type' returns a pointer
	// to a variable, whereas 'TypeVar' accepts a pointer to a variable.
	var optTypeVar float64
	golf.FloatVar(&optTypeVar, "f", 6.02e-23, "optTypeVar takes a pointer to a variable and a single flag")

	// For completeness, a 'TypeVarP' set of functions is also provided.
	var optTypeVarP int64
	golf.Int64VarP(&optTypeVarP, 'i', "int64", 13, "optTypeVarP takes a pointer to a variable, a rune, and a string")

	golf.Parse()

	fmt.Println("optType: ", *optType1)
	fmt.Println("optTypeP: ", *optTypeP)

	fmt.Println("optTypeVar: ", optTypeVar)
	fmt.Println("optTypeVarP: ", optTypeVarP)
}
