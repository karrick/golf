package golf

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

var options []option

// optionFromLong performs linear search for the option with a matching long
// option name in the list of options. It returns the option found, or nil if
// the requested long option name was not found.
func optionFromLong(long string) option {
	for _, option := range options {
		if option.Long() == long {
			return option
		}
	}
	return nil
}

// optionFromShort performs linear search for the option with a matching short
// option name in the list of options. It returns the option found, or nil if
// the requested short option name was not found.
func optionFromShort(short rune) option {
	for _, option := range options {
		if option.Short() == short {
			return option
		}
	}
	return nil
}

func parseAndCheckFlags(short, long string) (rune, error) {
	var r rune
	if short == "" && long == "" {
		return r, errors.New("cannot add option without either short, long, or both flags")
	}
	if short != "" {
		if r, _ = utf8.DecodeRuneInString(short); r == utf8.RuneError {
			return r, fmt.Errorf("cannot decode first rune of short flag: %q", short)
		}
		if r == '-' {
			return r, fmt.Errorf("cannot set short flag to a hyphen: %q", short)
		}
	}
	if strings.HasPrefix(long, "-") {
		return r, fmt.Errorf("cannot start long flag with a hyphen: %q", long)
	}
	if err := redefinition(r, long); err != nil {
		return r, err
	}
	return r, nil
}

func redefinition(short rune, long string) error {
	var zeroValue rune
	for _, opt := range options {
		if long != "" && long == opt.Long() {
			return fmt.Errorf("cannot add option that duplicates long flag: %q", long)
		}
		if short != zeroValue && short == opt.Short() {
			return fmt.Errorf("cannot add option that duplicates short flag: %q", short)
		}
	}
	return nil
}

// option is list of methods any concrete option needs to have for use by parser.
type option interface {
	Default() interface{}   // default value of command line option
	Description() string    // describes the command line option
	Long() string           // long flag
	NextState() parserState // next state for state machine
	Short() rune            // short flag
}

type optionBool struct {
	pv          *bool
	long        string
	description string
	short       rune
	def         bool
}

func (o optionBool) Default() interface{}   { return o.def }
func (o optionBool) Description() string    { return o.description }
func (o optionBool) Long() string           { return o.long }
func (o optionBool) NextState() parserState { return wantBool }
func (o optionBool) Short() rune            { return o.short }

// Bool returns a pointer to a bool command line option.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func Bool(short, long string, value bool, description string) *bool {
	var v bool
	BoolVar(&v, short, long, value, description)
	return &v
}

// BoolVar binds an existing boolean variable to a flag.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func BoolVar(pv *bool, short, long string, value bool, description string) {
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	*pv = value
	options = append(options, &optionBool{
		description: description,
		long:        long,
		short:       r,
		pv:          pv,
		def:         value,
	})
}

type optionDuration struct {
	pv          *time.Duration
	description string
	long        string
	short       rune
	def         time.Duration
}

func (o optionDuration) Default() interface{}   { return o.def }
func (o optionDuration) Description() string    { return o.description }
func (o optionDuration) Long() string           { return o.long }
func (o optionDuration) NextState() parserState { return wantDuration }
func (o optionDuration) Short() rune            { return o.short }

// Duration returns a pointer to a duration command line option.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func Duration(short, long string, value time.Duration, description string) *time.Duration {
	var v time.Duration
	DurationVar(&v, short, long, value, description)
	return &v
}

// DurationVar binds an existing time.Duration variable to a flag.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func DurationVar(pv *time.Duration, short, long string, value time.Duration, description string) {
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	*pv = value
	options = append(options, &optionDuration{
		description: description,
		long:        long,
		short:       r,
		pv:          pv,
		def:         value,
	})
}

type optionFloat struct {
	pv          *float64
	description string
	short       rune
	long        string
	def         float64
}

func (o optionFloat) Default() interface{}   { return o.def }
func (o optionFloat) Description() string    { return o.description }
func (o optionFloat) Long() string           { return o.long }
func (o optionFloat) NextState() parserState { return wantFloat }
func (o optionFloat) Short() rune            { return o.short }

// Float returns a pointer to a float64 command line option.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func Float(short, long string, value float64, description string) *float64 {
	var v float64
	FloatVar(&v, short, long, value, description)
	return &v
}

// FloatVar binds an existing float64 variable to a flag.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func FloatVar(pv *float64, short, long string, value float64, description string) {
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	*pv = value
	options = append(options, &optionFloat{
		description: description,
		long:        long,
		short:       r,
		def:         value,
		pv:          pv,
	})
}

type optionInt struct {
	pv          *int
	description string
	short       rune
	long        string
	def         int
}

func (o optionInt) Default() interface{}   { return o.def }
func (o optionInt) Description() string    { return o.description }
func (o optionInt) Long() string           { return o.long }
func (o optionInt) NextState() parserState { return wantInt }
func (o optionInt) Short() rune            { return o.short }

// Int returns a pointer to an int command line option.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func Int(short, long string, value int, description string) *int {
	var v int
	IntVar(&v, short, long, value, description)
	return &v
}

// IntVar binds an existing int variable to a flag.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func IntVar(pv *int, short, long string, value int, description string) {
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	*pv = value
	options = append(options, &optionInt{
		description: description,
		long:        long,
		short:       r,
		def:         value,
		pv:          pv,
	})
}

type optionInt64 struct {
	pv          *int64
	description string
	short       rune
	long        string
	def         int64
}

func (o optionInt64) Default() interface{}   { return o.def }
func (o optionInt64) Description() string    { return o.description }
func (o optionInt64) Long() string           { return o.long }
func (o optionInt64) NextState() parserState { return wantInt64 }
func (o optionInt64) Short() rune            { return o.short }

// Int64 returns a pointer to an int64 command line option.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func Int64(short, long string, value int64, description string) *int64 {
	var v int64
	Int64Var(&v, short, long, value, description)
	return &v
}

// Int64Var binds an existing int64 variable to a flag.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func Int64Var(pv *int64, short, long string, value int64, description string) {
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	*pv = value
	options = append(options, &optionInt64{
		description: description,
		long:        long,
		short:       r,
		def:         value,
		pv:          pv,
	})
}

type optionUint struct {
	pv          *uint
	description string
	short       rune
	long        string
	def         uint
}

func (o optionUint) Default() interface{}   { return o.def }
func (o optionUint) Description() string    { return o.description }
func (o optionUint) Long() string           { return o.long }
func (o optionUint) NextState() parserState { return wantUint }
func (o optionUint) Short() rune            { return o.short }

// Uint returns a pouinter to an uint command line option.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func Uint(short, long string, value uint, description string) *uint {
	var v uint
	UintVar(&v, short, long, value, description)
	return &v
}

// UintVar binds an existing uint variable to a flag.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func UintVar(pv *uint, short, long string, value uint, description string) {
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	*pv = value
	options = append(options, &optionUint{
		description: description,
		long:        long,
		short:       r,
		def:         value,
		pv:          pv,
	})
}

type optionUint64 struct {
	pv          *uint64
	description string
	short       rune
	long        string
	def         uint64
}

func (o optionUint64) Default() interface{}   { return o.def }
func (o optionUint64) Description() string    { return o.description }
func (o optionUint64) Short() rune            { return o.short }
func (o optionUint64) Long() string           { return o.long }
func (o optionUint64) NextState() parserState { return wantUint64 }

// Uint64 returns a pouinter to an uint64 command line option.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func Uint64(short, long string, value uint64, description string) *uint64 {
	var v uint64
	Uint64Var(&v, short, long, value, description)
	return &v
}

// Uint64Var binds an existing uint64 variable to a flag.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func Uint64Var(pv *uint64, short, long string, value uint64, description string) {
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	*pv = value
	options = append(options, &optionUint64{
		description: description,
		long:        long,
		short:       r,
		def:         value,
		pv:          pv,
	})
}

type optionString struct {
	pv          *string
	description string
	short       rune
	long        string
	def         string
}

func (o optionString) Default() interface{}   { return o.def }
func (o optionString) Description() string    { return o.description }
func (o optionString) Short() rune            { return o.short }
func (o optionString) Long() string           { return o.long }
func (o optionString) NextState() parserState { return wantString }

// String returns a pointer to a string command line option.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func String(short, long string, value string, description string) *string {
	var v string
	StringVar(&v, short, long, value, description)
	return &v
}

// StringVar binds an existing string variable to a flag.
//
// While invoking panic is never a good idea from a library, neither is calling
// log.Fatal, however, the flag API from the standard library being emulated
// does not allow for returning an error.
func StringVar(pv *string, short, long string, value string, description string) {
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	*pv = value
	options = append(options, &optionString{
		description: description,
		long:        long,
		short:       r,
		def:         value,
		pv:          pv,
	})
}
