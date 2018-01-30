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
	options = nil
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

func (state parserState) String() string {
	switch state {
	case anything:
		return "nothing"
	case consumedSingleHyphen:
		return "consumed single hyphen"
	case wantShortOptionName:
		return "want short option name"
	case wantLongOptionName:
		return "want long option name"
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
		return fmt.Sprintf("unknown state value: %d", int(state))
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
	wantLongOptionName
	wantShortOptionName
	wantString
	wantUint
	wantUint64
)

func parse(line string) error {
	var state, stateAfterPossibleSpace parserState
	var longOptionName, optionValue string
	var opt option

	for _, r := range line {
		// fmt.Fprintf(os.Stderr, "state: %v; rune: %c\n", state, r)

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
				return errors.New("hyphen without options")
			} else if r == '-' {
				longOptionName = ""
				state = wantLongOptionName
			} else {
				opt = optionFromShort(r)
				if opt == nil {
					return fmt.Errorf("unknown option: %q", r)
				}
				switch next := opt.NextState(); next {
				case wantBool:
					*opt.(*optionBool).pv = true
					argsProcessed++
					state = wantShortOptionName
				default:
					optionValue = ""
					state = ignorePossibleSpace
					stateAfterPossibleSpace = next
				}
			}

		case wantShortOptionName:
			if isSpace {
				state = anything
			} else if r == '-' {
				return fmt.Errorf("cannot parse argument: %q", os.Args[argIndex])
			} else {
				opt = optionFromShort(r)
				if opt == nil {
					return fmt.Errorf("unknown option: %q", r)
				}
				switch next := opt.NextState(); next {
				case wantBool:
					*opt.(*optionBool).pv = true
					argsProcessed++
				default:
					optionValue = ""
					state = ignorePossibleSpace
					stateAfterPossibleSpace = next
				}
			}

		case wantLongOptionName:
			if isSpace {
				if longOptionName == "" {
					// NOTE: equivalent to reading double hyphen followed by
					// space: done processing args
					return nil
				}

				opt = optionFromLong(longOptionName)
				if opt == nil {
					return fmt.Errorf("unknown option: %q", longOptionName)
				}
				switch next := opt.NextState(); next {
				case wantBool:
					*opt.(*optionBool).pv = true
					argsProcessed++
					state = anything
				default:
					optionValue = ""
					state = ignorePossibleSpace
					stateAfterPossibleSpace = next
				}
			} else {
				longOptionName += string(r)
			}

		case wantDuration:
			if isSpace {
				value, err := time.ParseDuration(optionValue)
				if err != nil {
					return err
				}
				*opt.(*optionDuration).pv = value
				argsProcessed++
				state = anything
			} else {
				optionValue += string(r)
			}

		case wantFloat:
			if isSpace {
				value, err := strconv.ParseFloat(optionValue, 64)
				if err != nil {
					return err
				}
				*opt.(*optionFloat).pv = value
				argsProcessed++
				state = anything
			} else {
				optionValue += string(r)
			}

		case wantInt:
			if isSpace {
				value, err := strconv.ParseInt(optionValue, 10, 64)
				if err != nil {
					return err
				}
				*opt.(*optionInt).pv = int(value)
				argsProcessed++
				state = anything
			} else {
				optionValue += string(r)
			}

		case wantInt64:
			if isSpace {
				value, err := strconv.ParseInt(optionValue, 10, 64)
				if err != nil {
					return err
				}
				*opt.(*optionInt64).pv = value
				argsProcessed++
				state = anything
			} else {
				optionValue += string(r)
			}

		case wantUint:
			if isSpace {
				value, err := strconv.ParseUint(optionValue, 10, 64)
				if err != nil {
					return err
				}
				*opt.(*optionUint).pv = uint(value)
				argsProcessed++
				state = anything
			} else {
				optionValue += string(r)
			}

		case wantUint64:
			if isSpace {
				value, err := strconv.ParseUint(optionValue, 10, 64)
				if err != nil {
					return err
				}
				*opt.(*optionUint64).pv = value
				argsProcessed++
				state = anything
			} else {
				optionValue += string(r)
			}

		case wantString:
			if isSpace {
				*opt.(*optionString).pv = optionValue
				argsProcessed++
				state = anything
			} else {
				optionValue += string(r)
			}

		}
	}

	// if state != anything {
	// 	fmt.Fprintf(os.Stderr, "state: %v; opt: %v\n", state, opt)
	// }

	// we might have hit the end of the command line while doing stuff
	switch state {
	case anything, wantShortOptionName:
		// nothing left to do
	case consumedSingleHyphen:
		return errors.New("hyphen without options")
	case ignorePossibleSpace:
		if long := opt.Long(); long != "" {
			return fmt.Errorf("option requires argument: %q", long)
		}
		return fmt.Errorf("option requires argument: %q", opt.Short())
	case wantLongOptionName:
		opt = optionFromLong(longOptionName)
		if opt == nil {
			return fmt.Errorf("unknown option: %q", longOptionName)
		}
		switch next := opt.NextState(); next {
		case wantBool:
			*opt.(*optionBool).pv = true
			argsProcessed++
		default:
			return fmt.Errorf("option requires argument: %q", longOptionName)
		}
	case wantDuration:
		if optionValue == "" {
			return errors.New("option requires argument")
		}
		value, err := time.ParseDuration(optionValue)
		if err != nil {
			return err
		}
		*opt.(*optionDuration).pv = value
		argsProcessed++
	case wantFloat:
		if optionValue == "" {
			return fmt.Errorf("option requires argument")
		}
		value, err := strconv.ParseFloat(optionValue, 64)
		if err != nil {
			return err
		}
		*opt.(*optionFloat).pv = value
		argsProcessed++
	case wantInt:
		if optionValue == "" {
			return fmt.Errorf("option requires argument")
		}
		value, err := strconv.ParseInt(optionValue, 10, 64)
		if err != nil {
			return err
		}
		*opt.(*optionInt).pv = int(value)
		argsProcessed++
	case wantInt64:
		if optionValue == "" {
			return fmt.Errorf("option requires argument")
		}
		value, err := strconv.ParseInt(optionValue, 10, 64)
		if err != nil {
			return err
		}
		*opt.(*optionInt64).pv = value
		argsProcessed++
	case wantUint:
		if optionValue == "" {
			return fmt.Errorf("option requires argument")
		}
		value, err := strconv.ParseUint(optionValue, 10, 64)
		if err != nil {
			return err
		}
		*opt.(*optionUint).pv = uint(value)
		argsProcessed++
	case wantUint64:
		if optionValue == "" {
			return fmt.Errorf("option requires argument")
		}
		value, err := strconv.ParseUint(optionValue, 10, 64)
		if err != nil {
			return err
		}
		*opt.(*optionUint64).pv = value
		argsProcessed++
	case wantString:
		if optionValue == "" {
			return fmt.Errorf("option requires argument")
		}
		*opt.(*optionString).pv = optionValue
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
