package golf

import "time"
import "fmt"

// Bool returns a pointer to a bool command line option, allowing for either a
// short or a long flag. If both are desired, use the BoolP function.
func (p *Parser) Bool(flag string, value bool, description string) *bool {
	v := value
	p.BoolVar(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// BoolP returns a pointer to a bool command line option, allowing for both a
// short and a long flag.
func (p *Parser) BoolP(short rune, long string, value bool, description string) *bool {
	v := value
	p.BoolVarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// BoolVar updates the Parser to recognize flag as a bool with the default
// value and description.
func (p *Parser) BoolVar(pv *bool, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short string
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionBool{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       short,
	})
	return p
}

// BoolVar updates the Parser to recognize short and long flag as a bool with
// the default value and description.
func (p *Parser) BoolVarP(pv *bool, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionBool{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       fmt.Sprintf("%c", short),
	})
	return p
}

// Duration returns a pointer to a time.Duration command line option, allowing
// for either a short or a long flag. If both are desired, use the DurationP
// function.
func (p *Parser) Duration(flag string, value time.Duration, description string) *time.Duration {
	v := value
	p.DurationVar(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// DurationP returns a pointer to a time.Duration command line option,
// allowing for both a short and a long flag.
func (p *Parser) DurationP(short rune, long string, value time.Duration, description string) *time.Duration {
	v := value
	p.DurationVarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// DurationVar updates the Parser to recognize flag as a duration with the
// default value and description.
func (p *Parser) DurationVar(pv *time.Duration, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short string
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionDuration{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       short,
	})
	return p
}

// DurationVarP updates the Parser to recognize short and long flag as a
// duration with the default value and description.
func (p *Parser) DurationVarP(pv *time.Duration, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionDuration{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       fmt.Sprintf("%c", short),
	})
	return p
}

// Float returns a pointer to a float64 command line option, allowing for
// either a short or a long flag. If both are desired, use the FloatP
// function.
func (p *Parser) Float(flag string, value float64, description string) *float64 {
	v := value
	p.FloatVar(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// FloatP returns a pointer to a float64 command line option, allowing for
// both a short and a long flag.
func (p *Parser) FloatP(short rune, long string, value float64, description string) *float64 {
	v := value
	p.FloatVarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// FloatVar updates the Parser to recognize flag as a float with the default
// value and description.
func (p *Parser) FloatVar(pv *float64, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short string
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionFloat{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       short,
	})
	return p
}

// FloatVarP updates the Parser to recognize short and long flag as a float
// with the default value and description.
func (p *Parser) FloatVarP(pv *float64, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionFloat{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       fmt.Sprintf("%c", short),
	})
	return p
}

// Int returns a pointer to a int command line option, allowing for either a
// short or a long flag. If both are desired, use the IntP function.
func (p *Parser) Int(flag string, value int, description string) *int {
	v := value
	p.IntVar(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// IntP returns a pointer to a int command line option, allowing for both a
// short and a long flag.
func (p *Parser) IntP(short rune, long string, value int, description string) *int {
	v := value
	p.IntVarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// IntVar updates the Parser to recognize flag as a int with the default value
// and description.
func (p *Parser) IntVar(pv *int, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short string
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionInt{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       short,
	})
	return p
}

// IntVarP updates the Parser to recognize short and long flag as a int with
// the default value and description.
func (p *Parser) IntVarP(pv *int, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionInt{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       fmt.Sprintf("%c", short),
	})
	return p
}

// Int64 returns a pointer to a int64 command line option, allowing for either
// a short or a long flag. If both are desired, use the Int64P function.
func (p *Parser) Int64(flag string, value int64, description string) *int64 {
	v := value
	p.Int64Var(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// Int64P returns a pointer to a int64 command line option, allowing for both
// a short and a long flag.
func (p *Parser) Int64P(short rune, long string, value int64, description string) *int64 {
	v := value
	p.Int64VarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// Int64Var updates the Parser to recognize flag as a int64 with the default
// value and description.
func (p *Parser) Int64Var(pv *int64, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short string
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionInt64{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       short,
	})
	return p
}

// Int64VarP updates the Parser to recognize short and long flag as a int64
// with the default value and description.
func (p *Parser) Int64VarP(pv *int64, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionInt64{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       fmt.Sprintf("%c", short),
	})
	return p
}

// String returns a pointer to a string command line option, allowing for
// either a short or a long flag. If both are desired, use the StringP
// function.
func (p *Parser) String(flag string, value string, description string) *string {
	v := value
	p.StringVar(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// StringP returns a pointer to a string command line option, allowing for
// both a short and a long flag.
func (p *Parser) StringP(short rune, long string, value string, description string) *string {
	v := value
	p.StringVarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// StringVar updates the Parser to recognize flag as a string with the default
// value and description.
func (p *Parser) StringVar(pv *string, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short string
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionString{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       short,
	})
	return p
}

// StringVarP updates the Parser to recognize short and long flag as a string
// with the default value and description.
func (p *Parser) StringVarP(pv *string, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionString{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       fmt.Sprintf("%c", short),
	})
	return p
}

// Uint returns a pouinter to a uint command line option, allowing for either
// a short or a long flag. If both are desired, use the UintP function.
func (p *Parser) Uint(flag string, value uint, description string) *uint {
	v := value
	p.UintVar(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// UintP returns a pouinter to a uint command line option, allowing for both a
// short and a long flag.
func (p *Parser) UintP(short rune, long string, value uint, description string) *uint {
	v := value
	p.UintVarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// UintVar updates the Parser to recognize flag as a uint with the default
// value and description.
func (p *Parser) UintVar(pv *uint, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short string
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionUint{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       short,
	})
	return p
}

// UintVarP updates the Parser to recognize short and long flag as a uint with
// the default value and description.
func (p *Parser) UintVarP(pv *uint, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionUint{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       fmt.Sprintf("%c", short),
	})
	return p
}

// Uint64 returns a pouinter to a uint64 command line option, allowing for
// either a short or a long flag. If both are desired, use the Uint64P
// function.
func (p *Parser) Uint64(flag string, value uint64, description string) *uint64 {
	v := value
	p.Uint64Var(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// Uint64P returns a pouinter to a uint64 command line option, allowing for
// both a short and a long flag.
func (p *Parser) Uint64P(short rune, long string, value uint64, description string) *uint64 {
	v := value
	p.Uint64VarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// Uint64Var updates the Parser to recognize flag as a uint64 with the default
// value and description.
func (p *Parser) Uint64Var(pv *uint64, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short string
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionUint64{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       short,
	})
	return p
}

// Uint64VarP updates the Parser to recognize short and long flag as a uint64
// with the default value and description.
func (p *Parser) Uint64VarP(pv *uint64, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.options = append(p.options, &optionUint64{
		def:         *pv,
		description: description,
		long:        long,
		pv:          pv,
		short:       fmt.Sprintf("%c", short),
	})
	return p
}
