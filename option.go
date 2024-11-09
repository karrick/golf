package golf

import "time"

// option is list of methods any concrete option needs to have for use by
// parser.
type option interface {
	Default() interface{} // default value of command line option
	Description() string  // describes the command line option
	Long() string         // long flag
	NextSlurp() slurpType // next state for state machine
	Short() string        // short flag

}

type optionBool struct {
	pv          *bool
	description string
	long        string
	short       string
	def         bool
}

func (o optionBool) Default() interface{} { return o.def }
func (o optionBool) Description() string  { return o.description }
func (o optionBool) Long() string         { return o.long }
func (o optionBool) NextSlurp() slurpType { return nothingToSlurp }
func (o optionBool) Short() string        { return o.short }

type optionDuration struct {
	pv          *time.Duration
	description string
	long        string
	short       string
	def         time.Duration
}

func (o optionDuration) Default() interface{} { return o.def }
func (o optionDuration) Description() string  { return o.description }
func (o optionDuration) Long() string         { return o.long }
func (o optionDuration) NextSlurp() slurpType { return slurpDuration }
func (o optionDuration) Short() string        { return o.short }

type optionFloat struct {
	pv          *float64
	description string
	long        string
	short       string
	def         float64
}

func (o optionFloat) Default() interface{} { return o.def }
func (o optionFloat) Description() string  { return o.description }
func (o optionFloat) Long() string         { return o.long }
func (o optionFloat) NextSlurp() slurpType { return slurpFloat }
func (o optionFloat) Short() string        { return o.short }

type optionInt struct {
	pv          *int
	description string
	long        string
	short       string
	def         int
}

func (o optionInt) Default() interface{} { return o.def }
func (o optionInt) Description() string  { return o.description }
func (o optionInt) Long() string         { return o.long }
func (o optionInt) NextSlurp() slurpType { return slurpInt }
func (o optionInt) Short() string        { return o.short }

type optionInt64 struct {
	pv          *int64
	description string
	long        string
	short       string
	def         int64
}

func (o optionInt64) Default() interface{} { return o.def }
func (o optionInt64) Description() string  { return o.description }
func (o optionInt64) Long() string         { return o.long }
func (o optionInt64) NextSlurp() slurpType { return slurpInt64 }
func (o optionInt64) Short() string        { return o.short }

type optionString struct {
	pv          *string
	description string
	long        string
	short       string
	def         string
}

func (o optionString) Default() interface{} { return o.def }
func (o optionString) Description() string  { return o.description }
func (o optionString) Long() string         { return o.long }
func (o optionString) NextSlurp() slurpType { return slurpString }
func (o optionString) Short() string        { return o.short }

type optionUint struct {
	pv          *uint
	description string
	long        string
	short       string
	def         uint
}

func (o optionUint) Default() interface{} { return o.def }
func (o optionUint) Description() string  { return o.description }
func (o optionUint) Long() string         { return o.long }
func (o optionUint) NextSlurp() slurpType { return slurpUint }
func (o optionUint) Short() string        { return o.short }

type optionUint64 struct {
	pv          *uint64
	description string
	long        string
	short       string
	def         uint64
}

func (o optionUint64) Default() interface{} { return o.def }
func (o optionUint64) Description() string  { return o.description }
func (o optionUint64) Long() string         { return o.long }
func (o optionUint64) NextSlurp() slurpType { return slurpUint64 }
func (o optionUint64) Short() string        { return o.short }
