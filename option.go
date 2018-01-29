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
	description string
	long        string
	short       rune
	def, value  bool
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
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	o := &optionBool{
		description: description,
		long:        long,
		short:       r,
		value:       value,
		def:         value,
	}
	options = append(options, o)
	return &o.value
}

type optionDuration struct {
	description string
	long        string
	short       rune
	def, value  time.Duration
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
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	o := &optionDuration{
		description: description,
		long:        long,
		short:       r,
		def:         value,
		value:       value,
	}
	options = append(options, o)
	return &o.value
}

type optionFloat struct {
	description string
	short       rune
	long        string
	def, value  float64
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
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	o := &optionFloat{
		description: description,
		long:        long,
		short:       r,
		def:         value,
		value:       value,
	}
	options = append(options, o)
	return &o.value
}

type optionInt struct {
	description string
	short       rune
	long        string
	def, value  int
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
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	o := &optionInt{
		description: description,
		long:        long,
		short:       r,
		def:         value,
		value:       value,
	}
	options = append(options, o)
	return &o.value
}

type optionInt64 struct {
	description string
	short       rune
	long        string
	def, value  int64
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
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	o := &optionInt64{
		description: description,
		long:        long,
		short:       r,
		def:         value,
		value:       value,
	}
	options = append(options, o)
	return &o.value
}

type optionUint struct {
	description string
	short       rune
	long        string
	def, value  uint
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
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	o := &optionUint{
		description: description,
		long:        long,
		short:       r,
		def:         value,
		value:       value,
	}
	options = append(options, o)
	return &o.value
}

type optionUint64 struct {
	description string
	short       rune
	long        string
	def, value  uint64
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
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	o := &optionUint64{
		description: description,
		long:        long,
		short:       r,
		def:         value,
		value:       value,
	}
	options = append(options, o)
	return &o.value
}

type optionString struct {
	description string
	short       rune
	long        string
	def, value  string
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
	r, err := parseAndCheckFlags(short, long)
	if err != nil {
		panic(err)
	}
	o := &optionString{
		description: description,
		long:        long,
		short:       r,
		def:         value,
		value:       value,
	}
	options = append(options, o)
	return &o.value
}
