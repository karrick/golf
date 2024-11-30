package golf

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// option is list of methods any concrete option needs to have for use by
// parser.
type option interface {
	Callback(string) error // Callback parses input text stores value when valid
	Default() interface{}  // default value of command line option
	Description() string   // describes the command line option
	Long() string          // long flag
	NextSlurp() slurpType  // next state for state machine
	Short() string         // short flag
}

type optionBool struct {
	callback    func(string) (interface{}, error)
	description string
	long        string
	short       string
	pv          *bool
	def         bool
}

func (o optionBool) Callback(text string) error {
	if o.callback != nil {
		v, err := o.callback(text)
		if err != nil {
			return err
		}
		return o.store(v)
	}

	switch strings.ToUpper(text) {
	case "0", "F", "FALSE", "N", "NO":
		return o.store(false)
	case "1", "T", "TRUE", "Y", "YES":
		return o.store(true)
	default:
		return fmt.Errorf("cannot parse as bool: %q", text)
	}
}

func (o optionBool) store(value interface{}) error {
	v, ok := value.(bool)
	if !ok {
		return fmt.Errorf("cannot store bool: %T(%v)", value, value)
	}
	*o.pv = v
	return nil
}

func (o optionBool) Default() interface{} { return o.def }
func (o optionBool) Description() string  { return o.description }
func (o optionBool) Long() string         { return o.long }
func (o optionBool) NextSlurp() slurpType { return nothingToSlurp }
func (o optionBool) Short() string        { return o.short }

type optionDuration struct {
	callback    func(string) (interface{}, error)
	description string
	long        string
	short       string
	pv          *time.Duration
	def         time.Duration
}

func (o optionDuration) Callback(text string) error {
	if o.callback != nil {
		v, err := o.callback(text)
		if err != nil {
			return err
		}
		return o.store(v)
	}

	v, err := time.ParseDuration(text)
	if err != nil {
		return err
	}
	return o.store(v)
}

func (o optionDuration) store(value interface{}) error {
	v, ok := value.(time.Duration)
	if !ok {
		return fmt.Errorf("cannot store time.Duration: %T(%v)", value, value)
	}
	*o.pv = v
	return nil
}

func (o optionDuration) Default() interface{} { return o.def }
func (o optionDuration) Description() string  { return o.description }
func (o optionDuration) Long() string         { return o.long }
func (o optionDuration) NextSlurp() slurpType { return slurpDuration }
func (o optionDuration) Short() string        { return o.short }

type optionFloat struct {
	callback    func(string) (interface{}, error)
	description string
	long        string
	short       string
	pv          *float64
	def         float64
}

func (o optionFloat) Callback(text string) error {
	if o.callback != nil {
		v, err := o.callback(text)
		if err != nil {
			return err
		}
		return o.store(v)
	}

	v, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return err
	}
	return o.store(v)
}

func (o optionFloat) store(value interface{}) error {
	v, ok := value.(float64)
	if !ok {
		return fmt.Errorf("cannot store float64: %T(%v)", value, value)
	}
	*o.pv = v
	return nil
}

func (o optionFloat) Default() interface{} { return o.def }
func (o optionFloat) Description() string  { return o.description }
func (o optionFloat) Long() string         { return o.long }
func (o optionFloat) NextSlurp() slurpType { return slurpFloat }
func (o optionFloat) Short() string        { return o.short }

type optionInt struct {
	callback    func(string) (interface{}, error)
	description string
	long        string
	short       string
	pv          *int
	def         int
}

func (o optionInt) Callback(text string) error {
	if o.callback != nil {
		v, err := o.callback(text)
		if err != nil {
			return err
		}
		return o.store(v)
	}

	v, err := strconv.Atoi(text)
	if err != nil {
		return err
	}
	return o.store(v)
}

func (o optionInt) store(value interface{}) error {
	v, ok := value.(int)
	if !ok {
		return fmt.Errorf("cannot store int: %T(%v)", value, value)
	}
	*o.pv = v
	return nil
}

func (o optionInt) Default() interface{} { return o.def }
func (o optionInt) Description() string  { return o.description }
func (o optionInt) Long() string         { return o.long }
func (o optionInt) NextSlurp() slurpType { return slurpInt }
func (o optionInt) Short() string        { return o.short }

type optionInt64 struct {
	callback    func(string) (interface{}, error)
	description string
	long        string
	short       string
	pv          *int64
	def         int64
}

func (o optionInt64) Callback(text string) error {
	if o.callback != nil {
		v, err := o.callback(text)
		if err != nil {
			return err
		}
		return o.store(v)
	}

	v, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return err
	}
	return o.store(v)
}

func (o optionInt64) store(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return fmt.Errorf("cannot store int64: %T(%v)", value, value)
	}
	*o.pv = v
	return nil
}

func (o optionInt64) Default() interface{} { return o.def }
func (o optionInt64) Description() string  { return o.description }
func (o optionInt64) Long() string         { return o.long }
func (o optionInt64) NextSlurp() slurpType { return slurpInt64 }
func (o optionInt64) Short() string        { return o.short }

type optionString struct {
	callback    func(string) (interface{}, error)
	description string
	long        string
	short       string
	pv          *string
	def         string
}

func (o optionString) Callback(text string) error {
	if o.callback != nil {
		v, err := o.callback(text)
		if err != nil {
			return err
		}
		return o.store(v)
	}

	return o.store(text)
}

func (o optionString) store(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return fmt.Errorf("cannot store string: %T(%v)", value, value)
	}
	*o.pv = v
	return nil
}

func (o optionString) Default() interface{} { return o.def }
func (o optionString) Description() string  { return o.description }
func (o optionString) Long() string         { return o.long }
func (o optionString) NextSlurp() slurpType { return slurpString }
func (o optionString) Short() string        { return o.short }

type optionUint struct {
	callback    func(string) (interface{}, error)
	description string
	long        string
	short       string
	pv          *uint
	def         uint
}

func (o optionUint) Callback(text string) error {
	if o.callback != nil {
		v, err := o.callback(text)
		if err != nil {
			return err
		}
		return o.store(v)
	}

	v, err := strconv.ParseUint(text, 10, 0)
	if err != nil {
		return err
	}
	return o.store(v)
}

func (o optionUint) store(value interface{}) error {
	v, ok := value.(uint)
	if !ok {
		return fmt.Errorf("cannot store uint: %T(%v)", value, value)
	}
	*o.pv = v
	return nil
}

func (o optionUint) Default() interface{} { return o.def }
func (o optionUint) Description() string  { return o.description }
func (o optionUint) Long() string         { return o.long }
func (o optionUint) NextSlurp() slurpType { return slurpUint }
func (o optionUint) Short() string        { return o.short }

type optionUint64 struct {
	callback    func(string) (interface{}, error)
	description string
	long        string
	short       string
	pv          *uint64
	def         uint64
}

func (o optionUint64) Callback(text string) error {
	if o.callback != nil {
		v, err := o.callback(text)
		if err != nil {
			return err
		}
		return o.store(v)
	}

	v, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		return err
	}
	return o.store(v)
}

func (o optionUint64) store(value interface{}) error {
	v, ok := value.(uint64)
	if !ok {
		return fmt.Errorf("cannot store uint64: %T(%v)", value, value)
	}
	*o.pv = v
	return nil
}

func (o optionUint64) Default() interface{} { return o.def }
func (o optionUint64) Description() string  { return o.description }
func (o optionUint64) Long() string         { return o.long }
func (o optionUint64) NextSlurp() slurpType { return slurpUint64 }
func (o optionUint64) Short() string        { return o.short }
