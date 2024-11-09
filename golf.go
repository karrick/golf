package golf

import (
	"fmt"
	"io"
	"os"
	"time"
)

var defaultParser *Parser

func init() {
	defaultParser = new(Parser)
}

// dw is default wrapper, used for printing default options.
var dw = LineWrapper{Max: 80, Prefix: "    "}

// Arg returns the i'th command-line argument. Arg(0) is the first remaining
// argument after flags have been processed. Arg returns an empty string if the
// requested element does not exist.
func Arg(i int) string {
	return defaultParser.Arg(i)
}

// Args returns the non-flag command-line arguments.
func Args() []string {
	return defaultParser.Args()
}

// NArg returns the number of arguments remaining after flags have been
// processed.
func NArg() int {
	return defaultParser.NArg()
}

// NFlag returns the number of command-line flags that have been set.
func NFlag() int {
	return defaultParser.NFlag()
}

// Parse parses the command line. On error, displays the usage of the command
// line and exits the program with status code 2.
func Parse() {
	if err := defaultParser.Parse(os.Args[1:]); err != nil {
		// NOTE: Format output then exit similar to how Go standard library
		// "flag" might.
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
		Usage()
		os.Exit(2)
	}
}

// Parsed reports whether the command-line flags have been parsed.
//
// Deprecated
func Parsed() bool {
	return defaultParser.Parsed()
}

// PrintDefaults prints to standard error, a usage message showing the default
// settings of all defined command-line flags.
func PrintDefaults() {
	defaultParser.PrintDefaults()
}

// PrintDefaultsTo prints to w, a usage message showing the default settings of
// all defined command-line flags.
func PrintDefaultsTo(w io.Writer) {
	defaultParser.PrintDefaultsTo(w)
}

// Usage prints command line usage to stderr, but may be overridden by programs
// that need to customize the usage information.
var Usage = func() {
	// NOTE: Format output similar to how Go standard library "flag" might.
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	defaultParser.PrintDefaults()
}

// Bool returns a pointer to a bool command line option, allowing for either a
// short or a long flag. If both are desired, use the BoolP function.
func Bool(flag string, value bool, description string) *bool {
	return defaultParser.WithBool(flag, value, description)
}

// BoolP returns a pointer to a bool command line option, allowing for both a
// short and a long flag.
func BoolP(short rune, long string, value bool, description string) *bool {
	return defaultParser.WithBoolP(short, long, value, description)
}

// BoolVar binds an existing boolean variable to a flag, allowing for either a
// short or a long flag. If both are desired, use the BoolVarP function.
func BoolVar(pv *bool, flag string, value bool, description string) {
	*pv = value
	defaultParser.WithBoolVar(pv, flag, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}

// BoolVarP binds an existing boolean variable to a flag, allowing for both a
// short and a long flag.
func BoolVarP(pv *bool, short rune, long string, value bool, description string) {
	*pv = value
	defaultParser.WithBoolVarP(pv, short, long, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}

// Duration returns a pointer to a time.Duration command line option, allowing
// for either a short or a long flag. If both are desired, use the DurationP
// function.
func Duration(flag string, value time.Duration, description string) *time.Duration {
	return defaultParser.WithDuration(flag, value, description)
}

// DurationP returns a pointer to a time.Duration command line option, allowing
// for both a short and a long flag.
func DurationP(short rune, long string, value time.Duration, description string) *time.Duration {
	return defaultParser.WithDurationP(short, long, value, description)
}

// DurationVar binds an existing time.Duration variable to a flag, allowing for
// either a short or a long flag. If both are desired, use the DurationVarP
// function.
func DurationVar(pv *time.Duration, flag string, value time.Duration, description string) {
	*pv = value
	defaultParser.WithDurationVar(pv, flag, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}

// DurationVarP binds an existing time.Duration variable to a flag, allowing for
// both a short and a long flag.
func DurationVarP(pv *time.Duration, short rune, long string, value time.Duration, description string) {
	*pv = value
	defaultParser.WithDurationVarP(pv, short, long, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}

// Float returns a pointer to a float64 command line option, allowing for either
// a short or a long flag. If both are desired, use the FloatP function.
func Float(flag string, value float64, description string) *float64 {
	return defaultParser.WithFloat(flag, value, description)
}

// FloatP returns a pointer to a float64 command line option, allowing for both
// a short and a long flag.
func FloatP(short rune, long string, value float64, description string) *float64 {
	return defaultParser.WithFloatP(short, long, value, description)
}

// FloatVar binds an existing float64 variable to a flag, allowing for either a
// short or a long flag. If both are desired, use the FloatVarP function.
func FloatVar(pv *float64, flag string, value float64, description string) {
	*pv = value
	defaultParser.WithFloatVar(pv, flag, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}

// FloatVarP binds an existing float64 variable to a flag, allowing for both a
// short and a long flag.
func FloatVarP(pv *float64, short rune, long string, value float64, description string) {
	*pv = value
	defaultParser.WithFloatVarP(pv, short, long, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}

// Int returns a pointer to a int command line option, allowing for either a
// short or a long flag. If both are desired, use the IntP function.
func Int(flag string, value int, description string) *int {
	return defaultParser.WithInt(flag, value, description)
}

// IntP returns a pointer to a int command line option, allowing for both a
// short and a long flag.
func IntP(short rune, long string, value int, description string) *int {
	return defaultParser.WithIntP(short, long, value, description)
}

// IntVar binds an existing int variable to a flag, allowing for either a short
// or a long flag. If both are desired, use the IntVarP function.
func IntVar(pv *int, flag string, value int, description string) {
	*pv = value
	defaultParser.WithIntVar(pv, flag, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}

// IntVarP binds an existing int variable to a flag, allowing for both a short
// and a long flag.
func IntVarP(pv *int, short rune, long string, value int, description string) {
	*pv = value
	defaultParser.WithIntVarP(pv, short, long, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}

// Int64 returns a pointer to a int64 command line option, allowing for either a
// short or a long flag. If both are desired, use the Int64P function.
func Int64(flag string, value int64, description string) *int64 {
	return defaultParser.WithInt64(flag, value, description)
}

// Int64P returns a pointer to a int64 command line option, allowing for both a
// short and a long flag.
func Int64P(short rune, long string, value int64, description string) *int64 {
	return defaultParser.WithInt64P(short, long, value, description)
}

// Int64Var binds an existing int64 variable to a flag, allowing for either a
// short or a long flag. If both are desired, use the Int64VarP function.
func Int64Var(pv *int64, flag string, value int64, description string) {
	*pv = value
	defaultParser.WithInt64Var(pv, flag, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}

// Int64VarP binds an existing int64 variable to a flag, allowing for both a
// short and a long flag.
func Int64VarP(pv *int64, short rune, long string, value int64, description string) {
	*pv = value
	defaultParser.WithInt64VarP(pv, short, long, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}

// String returns a postringer to a string command line option, allowing for either a
// short or a long flag. If both are desired, use the StringP function.
func String(flag string, value string, description string) *string {
	return defaultParser.WithString(flag, value, description)
}

// StringP returns a postringer to a string command line option, allowing for both a
// short and a long flag.
func StringP(short rune, long string, value string, description string) *string {
	return defaultParser.WithStringP(short, long, value, description)
}

// StringVar binds an existing string variable to a flag, allowing for either a short
// or a long flag. If both are desired, use the StringVarP function.
func StringVar(pv *string, flag string, value string, description string) {
	*pv = value
	defaultParser.WithStringVar(pv, flag, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}

// StringVarP binds an existing string variable to a flag, allowing for both a short
// and a long flag.
func StringVarP(pv *string, short rune, long string, value string, description string) {
	*pv = value
	defaultParser.WithStringVarP(pv, short, long, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}

// Uint returns a pouinter to a uint command line option, allowing for either a
// short or a long flag. If both are desired, use the UintP function.
func Uint(flag string, value uint, description string) *uint {
	return defaultParser.WithUint(flag, value, description)
}

// UintP returns a pouinter to a uint command line option, allowing for both a
// short and a long flag.
func UintP(short rune, long string, value uint, description string) *uint {
	return defaultParser.WithUintP(short, long, value, description)
}

// UintVar binds an existing uint variable to a flag, allowing for either a short
// or a long flag. If both are desired, use the UintVarP function.
func UintVar(pv *uint, flag string, value uint, description string) {
	*pv = value
	defaultParser.WithUintVar(pv, flag, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}

// UintVarP binds an existing uint variable to a flag, allowing for both a short
// and a long flag.
func UintVarP(pv *uint, short rune, long string, value uint, description string) {
	*pv = value
	defaultParser.WithUintVarP(pv, short, long, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}

// Uint64 returns a pointer to a uint64 command line option, allowing for either a
// short or a long flag. If both are desired, use the Uint64P function.
func Uint64(flag string, value uint64, description string) *uint64 {
	return defaultParser.WithUint64(flag, value, description)
}

// Uint64P returns a pointer to a uint64 command line option, allowing for both a
// short and a long flag.
func Uint64P(short rune, long string, value uint64, description string) *uint64 {
	return defaultParser.WithUint64P(short, long, value, description)
}

// Uint64Var binds an existing uint64 variable to a flag, allowing for either a
// short or a long flag. If both are desired, use the Uint64VarP function.
func Uint64Var(pv *uint64, flag string, value uint64, description string) {
	*pv = value
	defaultParser.WithUint64Var(pv, flag, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}

// Uint64VarP binds an existing uint64 variable to a flag, allowing for both a
// short and a long flag.
func Uint64VarP(pv *uint64, short rune, long string, value uint64, description string) {
	*pv = value
	defaultParser.WithUint64VarP(pv, short, long, description)
	if err := defaultParser.Err(); err != nil {
		panic(err)
	}
}
