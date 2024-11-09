package golf

import "time"

// WithBool returns a pointer to a bool command line option, allowing for
// either a short or a long flag. If both are desired, use the BoolP function.
func (p *Parser) WithBool(flag string, value bool, description string) *bool {
	v := value
	p.WithBoolVar(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithBoolP returns a pointer to a bool command line option, allowing for
// both a short and a long flag.
func (p *Parser) WithBoolP(short rune, long string, value bool, description string) *bool {
	v := value
	p.WithBoolVarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithBoolVar updates the Parser to recognize flag as a bool with the default
// value and description.
func (p *Parser) WithBoolVar(pv *bool, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short rune
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionBool{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}

// WithBoolVar updates the Parser to recognize short and long flag as a bool
// with the default value and description.
func (p *Parser) WithBoolVarP(pv *bool, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionBool{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}

// WithDuration returns a pointer to a time.Duration command line option,
// allowing for either a short or a long flag. If both are desired, use the
// DurationP function.
func (p *Parser) WithDuration(flag string, value time.Duration, description string) *time.Duration {
	v := value
	p.WithDurationVar(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithDurationP returns a pointer to a time.Duration command line option,
// allowing for both a short and a long flag.
func (p *Parser) WithDurationP(short rune, long string, value time.Duration, description string) *time.Duration {
	v := value
	p.WithDurationVarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithDurationVar updates the Parser to recognize flag as a duration with the
// default value and description.
func (p *Parser) WithDurationVar(pv *time.Duration, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short rune
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionDuration{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}

// WithDurationVar updates the Parser to recognize short and long flag as a
// duration with the default value and description.
func (p *Parser) WithDurationVarP(pv *time.Duration, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionDuration{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}

// WithFloat returns a pointer to a float64 command line option, allowing for
// either a short or a long flag. If both are desired, use the FloatP
// function.
func (p *Parser) WithFloat(flag string, value float64, description string) *float64 {
	v := value
	p.WithFloatVar(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithFloatP returns a pointer to a float64 command line option, allowing for
// both a short and a long flag.
func (p *Parser) WithFloatP(short rune, long string, value float64, description string) *float64 {
	v := value
	p.WithFloatVarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithFloatVar updates the Parser to recognize flag as a float with the
// default value and description.
func (p *Parser) WithFloatVar(pv *float64, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short rune
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionFloat{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}

// WithFloatVar updates the Parser to recognize short and long flag as a float
// with the default value and description.
func (p *Parser) WithFloatVarP(pv *float64, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionFloat{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}

// WithInt returns a pointer to a int command line option, allowing for either
// a short or a long flag. If both are desired, use the IntP function.
func (p *Parser) WithInt(flag string, value int, description string) *int {
	v := value
	p.WithIntVar(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithIntP returns a pointer to a int command line option, allowing for both
// a short and a long flag.
func (p *Parser) WithIntP(short rune, long string, value int, description string) *int {
	v := value
	p.WithIntVarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithIntVar updates the Parser to recognize flag as a int with the default
// value and description.
func (p *Parser) WithIntVar(pv *int, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short rune
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionInt{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}

// WithIntVar updates the Parser to recognize short and long flag as a int
// with the default value and description.
func (p *Parser) WithIntVarP(pv *int, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionInt{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}

// WithInt64 returns a pointer to a int64 command line option, allowing for
// either a short or a long flag. If both are desired, use the Int64P
// function.
func (p *Parser) WithInt64(flag string, value int64, description string) *int64 {
	v := value
	p.WithInt64Var(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithInt64P returns a pointer to a int64 command line option, allowing for
// both a short and a long flag.
func (p *Parser) WithInt64P(short rune, long string, value int64, description string) *int64 {
	v := value
	p.WithInt64VarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithInt64Var updates the Parser to recognize flag as a int64 with the
// default value and description.
func (p *Parser) WithInt64Var(pv *int64, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short rune
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionInt64{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}

// WithInt64Var updates the Parser to recognize short and long flag as a int64
// with the default value and description.
func (p *Parser) WithInt64VarP(pv *int64, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionInt64{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}

// WithString returns a postringer to a string command line option, allowing
// for either a short or a long flag. If both are desired, use the StringP
// function.
func (p *Parser) WithString(flag string, value string, description string) *string {
	v := value
	p.WithStringVar(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithStringP returns a postringer to a string command line option, allowing
// for both a short and a long flag.
func (p *Parser) WithStringP(short rune, long string, value string, description string) *string {
	v := value
	p.WithStringVarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithStringVar updates the Parser to recognize flag as a string with the
// default value and description.
func (p *Parser) WithStringVar(pv *string, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short rune
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionString{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}

// WithStringVar updates the Parser to recognize short and long flag as a
// string with the default value and description.
func (p *Parser) WithStringVarP(pv *string, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionString{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}

// WithUint returns a pouinter to a uint command line option, allowing for
// either a short or a long flag. If both are desired, use the UintP function.
func (p *Parser) WithUint(flag string, value uint, description string) *uint {
	v := value
	p.WithUintVar(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithUintP returns a pouinter to a uint command line option, allowing for
// both a short and a long flag.
func (p *Parser) WithUintP(short rune, long string, value uint, description string) *uint {
	v := value
	p.WithUintVarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithUintVar updates the Parser to recognize flag as a uint with the default
// value and description.
func (p *Parser) WithUintVar(pv *uint, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short rune
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionUint{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}

// WithUintVar updates the Parser to recognize short and long flag as a uint
// with the default value and description.
func (p *Parser) WithUintVarP(pv *uint, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionUint{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}

// WithUint64 returns a pouinter to a uint64 command line option, allowing for
// either a short or a long flag. If both are desired, use the Uint64P function.
func (p *Parser) WithUint64(flag string, value uint64, description string) *uint64 {
	v := value
	p.WithUint64Var(&v, flag, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithUint64P returns a pouinter to a uint64 command line option, allowing for
// both a short and a long flag.
func (p *Parser) WithUint64P(short rune, long string, value uint64, description string) *uint64 {
	v := value
	p.WithUint64VarP(&v, short, long, description)
	if err := p.Err(); err != nil {
		panic(err)
	}
	return &v
}

// WithUint64Var updates the Parser to recognize flag as a uint64 with the
// default value and description.
func (p *Parser) WithUint64Var(pv *uint64, flag string, description string) *Parser {
	if p.err != nil {
		return p
	}
	var short rune
	var long string
	short, long, p.err = p.parseSingleFlag(flag)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionUint64{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}

// WithUint64Var updates the Parser to recognize short and long flag as a
// uint64 with the default value and description.
func (p *Parser) WithUint64VarP(pv *uint64, short rune, long string, description string) *Parser {
	if p.err != nil {
		return p
	}
	p.err = p.parseShortAndLongFlag(short, long)
	if p.err != nil {
		return p
	}
	p.flags = append(p.flags, &optionUint64{
		description: description,
		long:        long,
		short:       short,
		pv:          pv,
		def:         *pv,
	})
	return p
}
