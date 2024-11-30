package golf

import "fmt"

type slurpType uint

const (
	nothingToSlurp slurpType = iota
	slurpDuration
	slurpFloat
	slurpInt
	slurpInt64
	slurpString
	slurpUint
	slurpUint64
)

func (slurp slurpType) String() string {
	switch slurp {
	case nothingToSlurp:
		return "nothing to slurp"
	case slurpDuration:
		return "slurp duration"
	case slurpFloat:
		return "slurp float"
	case slurpInt:
		return "slurp int"
	case slurpInt64:
		return "slurp int64"
	case slurpString:
		return "slurp string"
	case slurpUint:
		return "slurp uint"
	case slurpUint64:
		return "slurp uint64"
	default:
		return fmt.Sprintf("unknown slurp type value: %d", int(slurp))
	}
}

// runeParserStateType represents the possible states of the parser, implemented as a state machine.
type runeParserStateType uint

const (
	beginArgument runeParserStateType = iota
	consumedHyphen
	wantLongName
	wantShortFlagsOnly
	wantText
	wantArgument
)

// String returns text representation of parser state.
func (state runeParserStateType) String() string {
	switch state {
	case beginArgument:
		return "new argument"
	case consumedHyphen:
		return "consumed single hyphen"
	case wantLongName:
		return "want long name"
	case wantShortFlagsOnly:
		return "want short flags only"
	case wantText:
		return "want text"
	case wantArgument:
		return "want argument"
	default:
		return fmt.Sprintf("unknown rune parser state value: %d", int(state))
	}
}
