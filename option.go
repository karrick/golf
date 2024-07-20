package golf

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

// flagFromLongName performs linear search for the flag with a matching long
// flag name in the list of flags. It returns the flag found, or nil if the
// requested long flag name was not found.
func flagFromLongName(long string) option {
	for _, option := range flags {
		if option.Long() == long {
			return option
		}
	}
	return nil
}

// flagFromShortName performs linear search for the flag with a matching short
// flag name in the list of flags. It returns the flag found, or nil if the
// requested short flag name was not found.
func flagFromShortName(short rune) option {
	for _, option := range flags {
		if option.Short() == short {
			return option
		}
	}
	return nil
}

// redefinition ensures the specified short and long flags do not redefine any
// existing flag definitions.
func redefinition(short rune, long string) error {
	for _, opt := range flags {
		if long != "" && long == opt.Long() {
			return fmt.Errorf("cannot add option that duplicates long flag: %q", long)
		}
		if short != utf8.RuneError && short == opt.Short() {
			return fmt.Errorf("cannot add option that duplicates short flag: %q", short)
		}
	}
	return nil
}

// parseShortAndLongFlag is called when there is both a short and a long flag to
// validate and ensure there are no duplicates.
func parseShortAndLongFlag(short rune, long string) error {

	switch short {
	case utf8.RuneError:
		return fmt.Errorf("cannot use flag with invalid rune: %q", short)
	case '-':
		return fmt.Errorf("cannot use hyphen as a flag: %q", short)
	}

	switch {
	case long == "":
		return errors.New("cannot use empty flag string")
	case strings.HasPrefix(long, "-"):
		return fmt.Errorf("cannot use flag that starts with a hyphen: %q", long)
	}

	return redefinition(short, long)
}

// parseSingleFlag is called when there is a single flag, and it is not known
// whether the flag is short or long. If validates the flag and ensures it is
// not a duplicate.
func parseSingleFlag(flag string) (rune, string, error) {
	if flag == "" {
		return utf8.RuneError, "", errors.New("cannot use empty flag string")
	}

	var firstRune rune
	var bufIndex, runeCount int

	// Ensure all bytes are valid runes
	buf := []byte(flag)

	for bufIndex < len(buf) {
		thisRune, runeSize := utf8.DecodeRune(buf[bufIndex:])
		if thisRune == utf8.RuneError {
			return thisRune, "", fmt.Errorf("cannot use flag with invalid rune: %q", flag)
		}
		if runeCount == 0 {
			if thisRune == '-' {
				return thisRune, "", fmt.Errorf("cannot use flag that starts with a hyphen: %q", flag)
			}
			firstRune = thisRune
		}
		runeCount++
		bufIndex += runeSize
	}

	if runeCount == 1 {
		return firstRune, "", redefinition(firstRune, "")
	}
	return utf8.RuneError, flag, redefinition(utf8.RuneError, flag)
}

// option is list of methods any concrete option needs to have for use by
// parser.
type option interface {
	Default() interface{} // default value of command line option
	Description() string  // describes the command line option
	Long() string         // long flag
	NextSlurp() slurpType // next state for state machine
	Short() rune          // short flag
	Callback() error      // call the callback function set for this option
}

type optionBool struct {
	pv          *bool
	cb          *func(bool) error
	long        string
	description string
	short       rune
	def         bool
}

func (o optionBool) Default() interface{} { return o.def }
func (o optionBool) Description() string  { return o.description }
func (o optionBool) Long() string         { return o.long }
func (o optionBool) NextSlurp() slurpType { return nothingToSlurp }
func (o optionBool) Short() rune          { return o.short }

func (o optionBool) Callback() error {
	if o.cb != nil {
		return (*o.cb)(*o.pv)
	}
	return nil
}

func (o optionBool) toggleOption() error {
	*o.pv = !o.def
	return o.Callback()
}

// Bool returns a pointer to a bool command line option, allowing for either a
// short or a long flag. If both are desired, use the BoolP function.
func Bool(flag string, value bool, description string) *bool {
	var v bool
	BoolVar(&v, flag, value, description)
	return &v
}

// BoolVar binds an existing boolean variable to a flag, allowing for either a
// short or a long flag. If both are desired, use the BoolVarP function.
func BoolVar(pv *bool, flag string, value bool, description string) {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionBool{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// BoolFunc creates a boolean command line option with either a short or a long
// flag, and sets a callback function that will be called when the option is
// processed. It returns a pointer to the variable.
func BoolFunc(flag string, value bool, description string, cb func(bool) error) *bool {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	return makeBoolFunc(short, long, value, description, cb)
}

// BoolP returns a pointer to a bool command line option, allowing for both a
// short and a long flag.
func BoolP(short rune, long string, value bool, description string) *bool {
	var v bool
	BoolVarP(&v, short, long, value, description)
	return &v
}

// BoolVarP binds an existing boolean variable to a flag, allowing for both a
// short and a long flag.
func BoolVarP(pv *bool, short rune, long string, value bool, description string) {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionBool{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// BoolFuncP creates a boolean command line option with both a short and a long
// flag, and sets a callback function that will be called when the option is
// processed. It returns a pointer to the variable.
func BoolFuncP(short rune, long string, value bool, description string, cb func(bool) error) *bool {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	return makeBoolFunc(short, long, value, description, cb)
}

func makeBoolFunc(short rune, long string, value bool, description string, cb func(bool) error) *bool {
	v := value
	flags = append(flags, &optionBool{
		description: description,
		long:        long,
		short:       short,
		pv:          &v,
		cb:          &cb,
		def:         value,
	})
	return &v
}

type optionDuration struct {
	pv          *time.Duration
	cb          *func(time.Duration) error
	description string
	long        string
	short       rune
	def         time.Duration
}

func (o optionDuration) Default() interface{} { return o.def }
func (o optionDuration) Description() string  { return o.description }
func (o optionDuration) Long() string         { return o.long }
func (o optionDuration) NextSlurp() slurpType { return slurpDuration }
func (o optionDuration) Short() rune          { return o.short }

func (o optionDuration) Callback() error {
	if o.cb != nil {
		return (*o.cb)(*o.pv)
	}
	return nil
}

// Duration returns a pointer to a time.Duration command line option, allowing
// for either a short or a long flag. If both are desired, use the DurationP
// function.
func Duration(flag string, value time.Duration, description string) *time.Duration {
	var v time.Duration
	DurationVar(&v, flag, value, description)
	return &v
}

// DurationVar binds an existing time.Duration variable to a flag, allowing for
// either a short or a long flag. If both are desired, use the DurationVarP
// function.
func DurationVar(pv *time.Duration, flag string, value time.Duration, description string) {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionDuration{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// DurationFunc creates a time.Duration command line option with either a short
// or a long flag, and sets a callback function that will be called when the
// option is processed. It returns a pointer to the variable.
func DurationFunc(flag string, value time.Duration, description string, cb func(time.Duration) error) *time.Duration {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	return makeDurationFunc(short, long, value, description, cb)
}

// DurationP returns a pointer to a time.Duration command line option, allowing
// for both a short and a long flag.
func DurationP(short rune, long string, value time.Duration, description string) *time.Duration {
	var v time.Duration
	DurationVarP(&v, short, long, value, description)
	return &v
}

// DurationVarP binds an existing time.Duration variable to a flag, allowing for
// both a short and a long flag.
func DurationVarP(pv *time.Duration, short rune, long string, value time.Duration, description string) {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionDuration{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// DurationFuncP creates a time.Duration command line option with both a short
// and a long flag, and sets a callback function that will be called when the
// option is processed. It returns a pointer to the variable that contains the
// value.
func DurationFuncP(short rune, long string, value time.Duration, description string, cb func(time.Duration) error) *time.Duration {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	return makeDurationFunc(short, long, value, description, cb)
}

func makeDurationFunc(short rune, long string, value time.Duration, description string, cb func(time.Duration) error) *time.Duration {
	v := value
	flags = append(flags, &optionDuration{
		description: description,
		long:        long,
		short:       short,
		pv:          &v,
		cb:          &cb,
		def:         value,
	})
	return &v
}

type optionFloat struct {
	pv          *float64
	cb          *func(float64) error
	description string
	short       rune
	long        string
	def         float64
}

func (o optionFloat) Default() interface{} { return o.def }
func (o optionFloat) Description() string  { return o.description }
func (o optionFloat) Long() string         { return o.long }
func (o optionFloat) NextSlurp() slurpType { return slurpFloat }
func (o optionFloat) Short() rune          { return o.short }

func (o optionFloat) Callback() error {
	if o.cb != nil {
		return (*o.cb)(*o.pv)
	}
	return nil
}

// Float returns a pointer to a float64 command line option, allowing for either
// a short or a long flag. If both are desired, use the FloatP function.
func Float(flag string, value float64, description string) *float64 {
	var v float64
	FloatVar(&v, flag, value, description)
	return &v
}

// FloatVar binds an existing float64 variable to a flag, allowing for either a
// short or a long flag. If both are desired, use the FloatVarP function.
func FloatVar(pv *float64, flag string, value float64, description string) {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionFloat{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// FloatFunc creates a float64 command line option with either a short
// or a long flag, and sets a callback function that will be called when the
// option is processed. It returns a pointer to the variable.
func FloatFunc(flag string, value float64, description string, cb func(float64) error) *float64 {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	return makeFloatFunc(short, long, value, description, cb)
}

// FloatP returns a pointer to a float64 command line option, allowing for both
// a short and a long flag.
func FloatP(short rune, long string, value float64, description string) *float64 {
	var v float64
	FloatVarP(&v, short, long, value, description)
	return &v
}

// FloatVarP binds an existing float64 variable to a flag, allowing for both a
// short and a long flag.
func FloatVarP(pv *float64, short rune, long string, value float64, description string) {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionFloat{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// FloatFuncP creates a float64 command line option with both a short
// and a long flag, and sets a callback function that will be called when the
// option is processed. It returns a pointer to the variable that contains the
// value.
func FloatFuncP(short rune, long string, value float64, description string, cb func(float64) error) *float64 {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	return makeFloatFunc(short, long, value, description, cb)
}

func makeFloatFunc(short rune, long string, value float64, description string, cb func(float64) error) *float64 {
	v := value
	flags = append(flags, &optionFloat{
		description: description,
		long:        long,
		short:       short,
		pv:          &v,
		cb:          &cb,
		def:         value,
	})
	return &v
}

type optionInt struct {
	pv          *int
	cb          *func(int) error
	description string
	short       rune
	long        string
	def         int
}

func (o optionInt) Default() interface{} { return o.def }
func (o optionInt) Description() string  { return o.description }
func (o optionInt) Long() string         { return o.long }
func (o optionInt) NextSlurp() slurpType { return slurpInt }
func (o optionInt) Short() rune          { return o.short }

func (o optionInt) Callback() error {
	if o.cb != nil {
		return (*o.cb)(*o.pv)
	}
	return nil
}

// Int returns a pointer to a int command line option, allowing for either a
// short or a long flag. If both are desired, use the IntP function.
func Int(flag string, value int, description string) *int {
	var v int
	IntVar(&v, flag, value, description)
	return &v
}

// IntVar binds an existing int variable to a flag, allowing for either a short
// or a long flag. If both are desired, use the IntVarP function.
func IntVar(pv *int, flag string, value int, description string) {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionInt{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// IntFunc creates a int command line option with either a short
// or a long flag, and sets a callback function that will be called when the
// option is processed. It returns a pointer to the variable.
func IntFunc(flag string, value int, description string, cb func(int) error) *int {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	return makeIntFunc(short, long, value, description, cb)
}

// IntP returns a pointer to a int command line option, allowing for both a
// short and a long flag.
func IntP(short rune, long string, value int, description string) *int {
	var v int
	IntVarP(&v, short, long, value, description)
	return &v
}

// IntVarP binds an existing int variable to a flag, allowing for both a short
// and a long flag.
func IntVarP(pv *int, short rune, long string, value int, description string) {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionInt{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// IntFuncP creates a int command line option with both a short
// and a long flag, and sets a callback function that will be called when the
// option is processed. It returns a pointer to the variable that contains the
// value.
func IntFuncP(short rune, long string, value int, description string, cb func(int) error) *int {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	return makeIntFunc(short, long, value, description, cb)
}

func makeIntFunc(short rune, long string, value int, description string, cb func(int) error) *int {
	v := value
	flags = append(flags, &optionInt{
		description: description,
		long:        long,
		short:       short,
		pv:          &v,
		cb:          &cb,
		def:         value,
	})
	return &v
}

type optionInt64 struct {
	pv          *int64
	cb          *func(int64) error
	description string
	short       rune
	long        string
	def         int64
}

func (o optionInt64) Default() interface{} { return o.def }
func (o optionInt64) Description() string  { return o.description }
func (o optionInt64) Long() string         { return o.long }
func (o optionInt64) NextSlurp() slurpType { return slurpInt64 }
func (o optionInt64) Short() rune          { return o.short }

func (o optionInt64) Callback() error {
	if o.cb != nil {
		return (*o.cb)(*o.pv)
	}
	return nil
}

// Int64 returns a pointer to a int64 command line option, allowing for either a
// short or a long flag. If both are desired, use the Int64P function.
func Int64(flag string, value int64, description string) *int64 {
	var v int64
	Int64Var(&v, flag, value, description)
	return &v
}

// Int64Var binds an existing int64 variable to a flag, allowing for either a
// short or a long flag. If both are desired, use the Int64VarP function.
func Int64Var(pv *int64, flag string, value int64, description string) {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionInt64{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// Int64Func creates a int64 command line option with either a short
// or a long flag, and sets a callback function that will be called when the
// option is processed. It returns a pointer to the variable.
func Int64Func(flag string, value int64, description string, cb func(int64) error) *int64 {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	return makeInt64Func(short, long, value, description, cb)
}

// Int64P returns a pointer to a int64 command line option, allowing for both a
// short and a long flag.
func Int64P(short rune, long string, value int64, description string) *int64 {
	var v int64
	Int64VarP(&v, short, long, value, description)
	return &v
}

// Int64VarP binds an existing int64 variable to a flag, allowing for both a
// short and a long flag.
func Int64VarP(pv *int64, short rune, long string, value int64, description string) {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionInt64{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// Int64FuncP creates a int64 command line option with both a short
// and a long flag, and sets a callback function that will be called when the
// option is processed. It returns a pointer to the variable that contains the
// value.
func Int64FuncP(short rune, long string, value int64, description string, cb func(int64) error) *int64 {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	return makeInt64Func(short, long, value, description, cb)
}

func makeInt64Func(short rune, long string, value int64, description string, cb func(int64) error) *int64 {
	v := value
	flags = append(flags, &optionInt64{
		description: description,
		long:        long,
		short:       short,
		pv:          &v,
		cb:          &cb,
		def:         value,
	})
	return &v
}

type optionUint struct {
	pv          *uint
	cb          *func(uint) error
	description string
	short       rune
	long        string
	def         uint
}

func (o optionUint) Default() interface{} { return o.def }
func (o optionUint) Description() string  { return o.description }
func (o optionUint) Long() string         { return o.long }
func (o optionUint) NextSlurp() slurpType { return slurpUint }
func (o optionUint) Short() rune          { return o.short }

func (o optionUint) Callback() error {
	if o.cb != nil {
		return (*o.cb)(*o.pv)
	}
	return nil
}

// Uint returns a pouinter to a uint command line option, allowing for either a
// short or a long flag. If both are desired, use the UintP function.
func Uint(flag string, value uint, description string) *uint {
	var v uint
	UintVar(&v, flag, value, description)
	return &v
}

// UintVar binds an existing uint variable to a flag, allowing for either a short
// or a long flag. If both are desired, use the UintVarP function.
func UintVar(pv *uint, flag string, value uint, description string) {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionUint{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// UintFunc creates a uint command line option with either a short
// or a long flag, and sets a callback function that will be called when the
// option is processed. It returns a pointer to the variable.
func UintFunc(flag string, value uint, description string, cb func(uint) error) *uint {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	return makeUintFunc(short, long, value, description, cb)
}

// UintP returns a pouinter to a uint command line option, allowing for both a
// short and a long flag.
func UintP(short rune, long string, value uint, description string) *uint {
	var v uint
	UintVarP(&v, short, long, value, description)
	return &v
}

// UintVarP binds an existing uint variable to a flag, allowing for both a short
// and a long flag.
func UintVarP(pv *uint, short rune, long string, value uint, description string) {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionUint{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// UintFuncP creates a uint command line option with both a short
// and a long flag, and sets a callback function that will be called when the
// option is processed. It returns a pointer to the variable that contains the
// value.
func UintFuncP(short rune, long string, value uint, description string, cb func(uint) error) *uint {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	return makeUintFunc(short, long, value, description, cb)
}

func makeUintFunc(short rune, long string, value uint, description string, cb func(uint) error) *uint {
	v := value
	flags = append(flags, &optionUint{
		description: description,
		long:        long,
		short:       short,
		pv:          &v,
		cb:          &cb,
		def:         value,
	})
	return &v
}

type optionUint64 struct {
	pv          *uint64
	cb          *func(uint64) error
	description string
	short       rune
	long        string
	def         uint64
}

func (o optionUint64) Default() interface{} { return o.def }
func (o optionUint64) Description() string  { return o.description }
func (o optionUint64) Long() string         { return o.long }
func (o optionUint64) NextSlurp() slurpType { return slurpUint64 }
func (o optionUint64) Short() rune          { return o.short }

func (o optionUint64) Callback() error {
	if o.cb != nil {
		return (*o.cb)(*o.pv)
	}
	return nil
}

// Uint64 returns a pointer to a uint64 command line option, allowing for either a
// short or a long flag. If both are desired, use the Uint64P function.
func Uint64(flag string, value uint64, description string) *uint64 {
	var v uint64
	Uint64Var(&v, flag, value, description)
	return &v
}

// Uint64Var binds an existing uint64 variable to a flag, allowing for either a
// short or a long flag. If both are desired, use the Uint64VarP function.
func Uint64Var(pv *uint64, flag string, value uint64, description string) {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionUint64{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// Uint64Func creates a uint64 command line option with either a short
// or a long flag, and sets a callback function that will be called when the
// option is processed. It returns a pointer to the variable.
func Uint64Func(flag string, value uint64, description string, cb func(uint64) error) *uint64 {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	return makeUint64Func(short, long, value, description, cb)
}

// Uint64P returns a pointer to a uint64 command line option, allowing for both a
// short and a long flag.
func Uint64P(short rune, long string, value uint64, description string) *uint64 {
	var v uint64
	Uint64VarP(&v, short, long, value, description)
	return &v
}

// Uint64VarP binds an existing uint64 variable to a flag, allowing for both a
// short and a long flag.
func Uint64VarP(pv *uint64, short rune, long string, value uint64, description string) {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionUint64{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// Uint64FuncP creates a uint64 command line option with both a short
// and a long flag, and sets a callback function that will be called when the
// option is processed. It returns a pointer to the variable that contains the
// value.
func Uint64FuncP(short rune, long string, value uint64, description string, cb func(uint64) error) *uint64 {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	return makeUint64Func(short, long, value, description, cb)
}

func makeUint64Func(short rune, long string, value uint64, description string, cb func(uint64) error) *uint64 {
	v := value
	flags = append(flags, &optionUint64{
		description: description,
		long:        long,
		short:       short,
		pv:          &v,
		cb:          &cb,
		def:         value,
	})
	return &v
}

type optionString struct {
	pv          *string
	cb          *func(string) error
	description string
	short       rune
	long        string
	def         string
}

func (o optionString) Default() interface{} { return o.def }
func (o optionString) Description() string  { return o.description }
func (o optionString) Long() string         { return o.long }
func (o optionString) NextSlurp() slurpType { return slurpString }
func (o optionString) Short() rune          { return o.short }

func (o optionString) Callback() error {
	if o.cb != nil {
		return (*o.cb)(*o.pv)
	}
	return nil
}

// String returns a postringer to a string command line option, allowing for either a
// short or a long flag. If both are desired, use the StringP function.
func String(flag string, value string, description string) *string {
	var v string
	StringVar(&v, flag, value, description)
	return &v
}

// StringVar binds an existing string variable to a flag, allowing for either a short
// or a long flag. If both are desired, use the StringVarP function.
func StringVar(pv *string, flag string, value string, description string) {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionString{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// StringFunc creates a string command line option with either a short
// or a long flag, and sets a callback function that will be called when the
// option is processed. It returns a pointer to the variable.
func StringFunc(flag string, value string, description string, cb func(string) error) *string {
	short, long, err := parseSingleFlag(flag)
	if err != nil {
		panic(err)
	}
	return makeStringFunc(short, long, value, description, cb)
}

// StringP returns a postringer to a string command line option, allowing for both a
// short and a long flag.
func StringP(short rune, long string, value string, description string) *string {
	var v string
	StringVarP(&v, short, long, value, description)
	return &v
}

// StringVarP binds an existing string variable to a flag, allowing for both a short
// and a long flag.
func StringVarP(pv *string, short rune, long string, value string, description string) {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	*pv = value
	flags = append(flags, &optionString{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         value,
	})
}

// StringFuncP creates a string command line option with both a short
// and a long flag, and sets a callback function that will be called when the
// option is processed. It returns a pointer to the variable that contains the
// value.
func StringFuncP(short rune, long string, value string, description string, cb func(string) error) *string {
	if err := parseShortAndLongFlag(short, long); err != nil {
		panic(err)
	}
	return makeStringFunc(short, long, value, description, cb)
}

func makeStringFunc(short rune, long string, value string, description string, cb func(string) error) *string {
	v := value
	flags = append(flags, &optionString{
		description: description,
		long:        long,
		short:       short,
		pv:          &v,
		cb:          &cb,
		def:         value,
	})
	return &v
}
