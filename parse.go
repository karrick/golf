package golf

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	argIndex      int  // keep track of os.Args index; increment counter when encounter space
	argsProcessed int  // keep track of how many arguments have been set
	parsed        bool // keep track of whether command line arguments have been parsed
)

func init() {
	resetParser()
}

func resetParser() {
	argsProcessed = 0
	argIndex = 1
	flags = nil
	parsed = false
}

// Parse parses the command line. On error, displays the usage of the command
// line and exits the program with status code 2.
func Parse() {
	// combine os.Args to a single string
	if err := parse(strings.Join(os.Args[1:], " ")); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		Usage()
		os.Exit(2)
	}
	parsed = true
}

// parserState represents the possible states of the parser, implemented as a state machine.
type parserState uint

// String returns text representation of parser state.
func (state parserState) String() string {
	switch state {
	case anything:
		return "nothing"
	case consumedSingleHyphen:
		return "consumed single hyphen"
	case wantLongFlagName:
		return "want long option name"
	case wantShortFlagName:
		return "want short option name"
	case ignorePossibleSpace:
		return "ignore possible space"
	case wantFloat:
		return "want float"
	case wantInt:
		return "want int"
	case wantInt64:
		return "want int64"
	case wantUint:
		return "want uint"
	case wantUint64:
		return "want uint64"
	case wantString:
		return "want string"
	case wantBool:
		return "want bool"
	case wantDuration:
		return "want duration"
	default:
		return fmt.Sprintf("unknown parser state value: %d", int(state))
	}
}

const (
	anything parserState = iota
	consumedSingleHyphen
	ignorePossibleSpace
	wantBool
	wantDuration
	wantFloat
	wantInt
	wantInt64
	wantLongFlagName
	wantShortFlagName
	wantString
	wantUint
	wantUint64
)

func parse(line string) error {
	var state, stateAfterPossibleSpace parserState
	var longFlagName, flagText string
	var f option
	var escaped bool

	for _, r := range line {
		// fmt.Fprintf(os.Stderr, "state: %v; rune: %c\n", state, r)

		if escaped {
			escaped = false
			switch state {
			case wantDuration, wantFloat, wantInt, wantInt64, wantUint, wantUint64, wantString:
				flagText += string(r)
			}
		}

		if r == '\\' {
			escaped = true
			continue
		}

		isSpace := unicode.IsSpace(r)
		if isSpace {
			argIndex++
		}

		if state == ignorePossibleSpace {
			state = stateAfterPossibleSpace
			if isSpace {
				continue // ignore this rune
			}
		}

		switch state {

		case anything:
			if isSpace {
				// no-op
			} else if r == '-' {
				state = consumedSingleHyphen
			} else {
				return nil
			}

		case consumedSingleHyphen:
			if isSpace {
				return errors.New("hyphen without flags")
			} else if r == '-' {
				longFlagName = ""
				state = wantLongFlagName
			} else {
				f = flagFromShortName(r)
				if f == nil {
					return fmt.Errorf("unknown flag: %q", r)
				}
				switch next := f.NextState(); next {
				case wantBool:
					*f.(*optionBool).pv = true
					argsProcessed++
					state = wantShortFlagName
				default:
					flagText = ""
					state = ignorePossibleSpace
					stateAfterPossibleSpace = next
				}
			}

		case wantShortFlagName:
			if isSpace {
				state = anything
			} else if r == '-' {
				return fmt.Errorf("cannot parse argument: %q", os.Args[argIndex])
			} else {
				f = flagFromShortName(r)
				if f == nil {
					return fmt.Errorf("unknown flag: %q", r)
				}
				switch next := f.NextState(); next {
				case wantBool:
					*f.(*optionBool).pv = true
					argsProcessed++
				default:
					flagText = ""
					state = ignorePossibleSpace
					stateAfterPossibleSpace = next
				}
			}

		case wantLongFlagName:
			if isSpace {
				if longFlagName == "" {
					// NOTE: equivalent to reading double hyphen followed by
					// space: done processing arguments
					return nil
				}

				f = flagFromLongName(longFlagName)
				if f == nil {
					return fmt.Errorf("unknown flag: %q", longFlagName)
				}
				switch next := f.NextState(); next {
				case wantBool:
					*f.(*optionBool).pv = true
					argsProcessed++
					state = anything
				default:
					flagText = ""
					state = ignorePossibleSpace
					stateAfterPossibleSpace = next
				}
			} else {
				longFlagName += string(r)
			}

		case wantDuration:
			if isSpace {
				value, err := time.ParseDuration(flagText)
				if err != nil {
					return err
				}
				*f.(*optionDuration).pv = value
				argsProcessed++
				state = anything
			} else {
				flagText += string(r)
			}

		case wantFloat:
			if isSpace {
				value, err := strconv.ParseFloat(flagText, 64)
				if err != nil {
					return err
				}
				*f.(*optionFloat).pv = value
				argsProcessed++
				state = anything
			} else {
				flagText += string(r)
			}

		case wantInt:
			if isSpace {
				value, err := strconv.ParseInt(flagText, 10, 64)
				if err != nil {
					return err
				}
				*f.(*optionInt).pv = int(value)
				argsProcessed++
				state = anything
			} else {
				flagText += string(r)
			}

		case wantInt64:
			if isSpace {
				value, err := strconv.ParseInt(flagText, 10, 64)
				if err != nil {
					return err
				}
				*f.(*optionInt64).pv = value
				argsProcessed++
				state = anything
			} else {
				flagText += string(r)
			}

		case wantUint:
			if isSpace {
				value, err := strconv.ParseUint(flagText, 10, 64)
				if err != nil {
					return err
				}
				*f.(*optionUint).pv = uint(value)
				argsProcessed++
				state = anything
			} else {
				flagText += string(r)
			}

		case wantUint64:
			if isSpace {
				value, err := strconv.ParseUint(flagText, 10, 64)
				if err != nil {
					return err
				}
				*f.(*optionUint64).pv = value
				argsProcessed++
				state = anything
			} else {
				flagText += string(r)
			}

		case wantString:
			if isSpace {
				*f.(*optionString).pv = flagText
				argsProcessed++
				state = anything
			} else {
				flagText += string(r)
			}

		}
	}

	// if state != anything {
	// 	fmt.Fprintf(os.Stderr, "state: %v; opt: %v\n", state, opt)
	// }

	// we might have hit the end of the command line while doing stuff
	switch state {
	case anything, wantShortFlagName:
		// nothing left to do
	case consumedSingleHyphen:
		return errors.New("hyphen without flags")
	case ignorePossibleSpace:
		if long := f.Long(); long != "" {
			return fmt.Errorf("flag requires argument: %q", long)
		}
		return fmt.Errorf("flag requires argument: %q", f.Short())
	case wantLongFlagName:
		f = flagFromLongName(longFlagName)
		if f == nil {
			return fmt.Errorf("unknown flag: %q", longFlagName)
		}
		switch next := f.NextState(); next {
		case wantBool:
			*f.(*optionBool).pv = true
			argsProcessed++
		default:
			return fmt.Errorf("flag requires argument: %q", longFlagName)
		}
	case wantDuration:
		if flagText == "" {
			return errors.New("flag requires argument")
		}
		value, err := time.ParseDuration(flagText)
		if err != nil {
			return err
		}
		*f.(*optionDuration).pv = value
		argsProcessed++
	case wantFloat:
		if flagText == "" {
			return fmt.Errorf("flag requires argument")
		}
		value, err := strconv.ParseFloat(flagText, 64)
		if err != nil {
			return err
		}
		*f.(*optionFloat).pv = value
		argsProcessed++
	case wantInt:
		if flagText == "" {
			return fmt.Errorf("flag requires argument")
		}
		value, err := strconv.ParseInt(flagText, 10, 64)
		if err != nil {
			return err
		}
		*f.(*optionInt).pv = int(value)
		argsProcessed++
	case wantInt64:
		if flagText == "" {
			return fmt.Errorf("flag requires argument")
		}
		value, err := strconv.ParseInt(flagText, 10, 64)
		if err != nil {
			return err
		}
		*f.(*optionInt64).pv = value
		argsProcessed++
	case wantUint:
		if flagText == "" {
			return fmt.Errorf("flag requires argument")
		}
		value, err := strconv.ParseUint(flagText, 10, 64)
		if err != nil {
			return err
		}
		*f.(*optionUint).pv = uint(value)
		argsProcessed++
	case wantUint64:
		if flagText == "" {
			return fmt.Errorf("flag requires argument")
		}
		value, err := strconv.ParseUint(flagText, 10, 64)
		if err != nil {
			return err
		}
		*f.(*optionUint64).pv = value
		argsProcessed++
	case wantString:
		if flagText == "" {
			return fmt.Errorf("flag requires argument")
		}
		*f.(*optionString).pv = flagText
		argsProcessed++
	default:
		return fmt.Errorf("unexpected parser state: %v", state)
	}

	return nil
}

// Parsed reports whether the command-line flags have been parsed.
func Parsed() bool {
	return parsed
}
