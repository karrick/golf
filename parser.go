package golf

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

// Parser can parse a series of command line arguments.
type Parser struct {
	options            []option
	remainingArguments []string // keep track of remaining arguments
	err                error
	argsProcessed      int  // keep track of how many arguments have been set
	parsed             bool // keep track of whether command line arguments have been parsed
}

// Arg returns the i'th command-line argument. Arg(0) is the first remaining
// argument after flags have been processed. Arg returns an empty string if
// the requested element does not exist.
func (p *Parser) Arg(i int) string {
	if i < len(p.remainingArguments) {
		return p.remainingArguments[i]
	}
	return ""
}

// Args returns the non-flag command-line arguments.
func (p *Parser) Args() []string {
	return p.remainingArguments
}

// ensureNoRedefinition ensures the specified single and double hyphen prefix
// options do not redefine any existing option definitions.
func (p *Parser) ensureNoRedefinition(short string, long string) error {
	for _, opt := range p.options {
		if long != "" && opt.Long() == long {
			return fmt.Errorf("cannot add option that duplicates long flag: %q", long)
		}
		if short != "" && opt.Short() == short {
			return fmt.Errorf("cannot add option that duplicates short flag: %q", short)
		}
	}
	return nil
}

// Err returns the error state of a parser.
func (p *Parser) Err() error {
	return p.err
}

// optionFromDoubleHyphenPrefix performs linear search for the option with a
// matching double-hyphen prefix in the list of options. It returns the option
// found, or nil if the requested name was not found.
func (p *Parser) optionFromDoubleHyphenPrefix(long string) option {
	if long == "" {
		return nil
	}
	for _, option := range p.options {
		if option.Long() == long {
			return option
		}
	}
	return nil
}

// optionFromSingleHyphenPrefix performs linear search for the option with a
// matching single-hyphen prefix in the list of options. It returns the option
// found, or nil if the requested short flag name was not found.
func (p *Parser) optionFromSingleHyphenPrefix(short rune) option {
	if utf8.RuneError == short {
		return nil
	}
	s := fmt.Sprintf("%c", short)
	for _, option := range p.options {
		if option.Short() == s {
			return option
		}
	}
	return nil
}

// NArg returns the number of arguments remaining after flags have been
// processed.
func (p *Parser) NArg() int {
	return len(p.remainingArguments)
}

// NFlag returns the number of command-line flags that have been processed.
func (p *Parser) NFlag() int {
	return p.argsProcessed
}

// Parse parses args.
func (p *Parser) Parse(args []string) error {
	if p.err != nil {
		return p.err // cannot parse when in state of error
	}

	// Reset parser.
	p.argsProcessed = 0
	p.remainingArguments = p.remainingArguments[:0]
	p.parsed = false

	var flagType slurpType
	var flagName, flagText string
	var f option

	for ai, arg := range args { // ai (arg index)
		debug("arg %d: %q; start argParserState: %v\n", ai, arg, flagType)

		if flagType != nothingToSlurp {
			p.err = f.Callback(arg)
			if p.err != nil {
				p.remainingArguments = append(p.remainingArguments, args[ai:]...)
				return p.err
			}
			flagType = nothingToSlurp
			p.argsProcessed++
			continue // with next argument
		}

		// nothing to slurp, so need to read runes one by one
		runeParserState := beginArgument
		flagText = ""

		for bi, r := range arg { // bi (byte index)
			debug("  runeParserState: %v; rune: %q\n", runeParserState, r)

			// NOTE: Cannot put this in switch statement because this needs
			// need to be able to break out of enclosing for statement that
			// loops over arg runes, and continue processing not the next
			// argument, but with the switch statement following this inner
			// loop.
			if runeParserState == wantText {
				flagText = arg[bi:]
				debug("  FLAG TEXT: %q\n", flagText)
				break // out of parsing this arg
			} else if runeParserState == wantLongName {
				flagName = arg[bi:]
				debug("  LONG FLAG NAME: %q\n", flagName)
				break // out of parsing this arg
			} else if runeParserState == beginArgument {
				if r != '-' {
					if false {
						debug("index: %d; this rune ends processing: %q\n", ai, r)
						p.parsed = true
						p.remainingArguments = append(p.remainingArguments, args[ai:]...)
						return nil
					}
					runeParserState = wantArgument
					break // out of parsing this arg
				}
				runeParserState = consumedHyphen
			} else if runeParserState == consumedHyphen {
				switch r {
				case '-':
					runeParserState = wantLongName
				default:
					if f = p.optionFromSingleHyphenPrefix(r); f == nil {
						p.remainingArguments = append(p.remainingArguments, args[ai:]...)
						p.err = fmt.Errorf("unknown flag: %q", r)
						return p.err
					}
					switch flagType = f.NextSlurp(); flagType {
					case nothingToSlurp:
						f.Callback("1")
						runeParserState = wantShortFlagsOnly
					default:
						runeParserState = wantText
					}
				}
			} else if runeParserState == wantShortFlagsOnly {
				if f = p.optionFromSingleHyphenPrefix(r); f == nil {
					p.remainingArguments = append(p.remainingArguments, args[ai:]...)
					p.err = fmt.Errorf("unknown flag: %q", r)
					return p.err
				}
				switch flagType = f.NextSlurp(); flagType {
				case nothingToSlurp:
					f.Callback("1")
				default:
					runeParserState = wantText
				}
			} else {
				panic(fmt.Errorf("TODO: runeParserState: %v; rune: %q", runeParserState, r))
			}
		}

		debug("  POST: runeParserState: %v\n", runeParserState)

		switch runeParserState {
		case consumedHyphen:
			p.remainingArguments = append(p.remainingArguments, args[ai:]...)
			p.err = errors.New("hyphen without flags")
			return p.err
		case wantArgument:
			p.remainingArguments = append(p.remainingArguments, arg)
		case wantText:
			if flagText == "" {
				debug("  finished arg and got no text: %q\n", flagText)
				p.argsProcessed++
				continue // with next arg, where we will slurp in the value
			}
			debug("  finished arg and got text: %q\n", flagText)
			if flagType == nothingToSlurp {
				panic(fmt.Errorf("got text %q but invalid nextSlurp: %v", flagText, flagType))
			}
			p.err = f.Callback(flagText)
			if p.err != nil {
				p.remainingArguments = append(p.remainingArguments, args[ai:]...)
				return p.err
			}
			flagType = nothingToSlurp
			p.argsProcessed++
		case wantShortFlagsOnly:
			debug("  finished arg while looking for short flags\n")
			p.argsProcessed++
		case wantLongName:
			if flagName == "" {
				// double-hyphen: remaining arguments everything after this
				p.argsProcessed++
				p.parsed = true
				p.remainingArguments = args[ai+1:]
				return nil
			}
			if f = p.optionFromDoubleHyphenPrefix(flagName); f == nil {
				p.err = fmt.Errorf("unknown flag: %q", flagName)
				p.remainingArguments = append(p.remainingArguments, args[ai:]...)
				return p.err
			}
			flagName = "" // reset
			if flagType = f.NextSlurp(); flagType == nothingToSlurp {
				f.Callback("1")
			}
			p.argsProcessed++
		default:
			p.err = fmt.Errorf("TODO: handle runeParserState: %v", runeParserState)
			p.remainingArguments = append(p.remainingArguments, args[ai:]...)
			return p.err
		}

		// POST: end of an indexed argument
		// debug("arg %d: %q end argParserState: %v\n", ai, arg, flagType)
		debug("  POST: %q end argParserState: %v\n", arg, flagType)
	}

	debug("after all args: %v\n", flagType)

	// we might have hit the end of the command line while doing stuff
	if flagType != nothingToSlurp {
		if long := f.Long(); long != "" {
			p.err = fmt.Errorf("flag requires argument: %q", long)
			return p.err
		}
		p.err = fmt.Errorf("flag requires argument: %q", f.Short())
		return p.err
	}

	p.parsed = true
	return nil
}

// parseShortAndLongFlag is called when there is both a short and a long flag to
// validate and ensure there are no duplicates.
func (p *Parser) parseShortAndLongFlag(short rune, long string) error {
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

	return p.ensureNoRedefinition(fmt.Sprintf("%c", short), long)
}

// parseSingleFlag is called when there is a single flag, and it is not known
// whether the flag is short or long. If validates the flag and ensures it is
// not a duplicate.
func (p *Parser) parseSingleFlag(flag string) (string, string, error) {
	if flag == "" {
		return "", "", errors.New("cannot use empty flag string")
	}

	var firstRune string
	var bufIndex, runeCount int

	// Ensure all bytes are valid runes
	buf := []byte(flag)

	for bufIndex < len(buf) {
		thisRune, runeSize := utf8.DecodeRune(buf[bufIndex:])
		thisRuneString := fmt.Sprintf("%c", thisRune)
		if thisRune == utf8.RuneError {
			return thisRuneString, "", fmt.Errorf("cannot use flag with invalid rune: %q", flag)
		}
		if runeCount == 0 {
			if thisRune == '-' {
				return thisRuneString, "", fmt.Errorf("cannot use flag that starts with a hyphen: %q", flag)
			}
			firstRune = thisRuneString
		}
		runeCount++
		bufIndex += runeSize
	}

	if runeCount == 1 {
		return firstRune, "", p.ensureNoRedefinition(firstRune, "")
	}
	return "", flag, p.ensureNoRedefinition("", flag)
}

// Parsed reports whether the command-line flags have been parsed.
func (p Parser) Parsed() bool {
	return p.parsed
}

// PrintDefaults prints to standard error, a usage message showing the default
// settings of all defined command-line flags.
func (p *Parser) PrintDefaults() {
	p.PrintDefaultsTo(os.Stderr)
}

// PrintDefaultsTo prints to w, a usage message showing the default settings of
// all defined command-line flags.
func (p *Parser) PrintDefaultsTo(w io.Writer) {
	for _, opt := range p.options {
		var def, typeName string
		description := opt.Description()
		value := opt.Default()

		switch value.(type) {
		case bool:
			typeName = "" // do not want to add a default blob when boolean
		case string, rune:
			def = fmt.Sprintf(" (default: %q)", value)
			typeName = fmt.Sprintf(" %T", value)
		default:
			def = fmt.Sprintf(" (default: %v)", value)
			typeName = fmt.Sprintf(" %T", value)
		}

		short := opt.Short()
		long := opt.Long()

		if short != "" {
			if long != "" {
				fmt.Fprintf(w, "  -%s, --%s%s%s\n", short, long, typeName, def)
			} else {
				fmt.Fprintf(w, "  -%s%s%s\n", short, typeName, def)
			}
		} else {
			fmt.Fprintf(w, "  --%s%s%s\n", long, typeName, def)
		}

		if description != "" {
			fmt.Fprint(w, dw.Wrap(fmt.Sprintf("%s", description)))
		}
	}
}
