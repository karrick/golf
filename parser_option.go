package golf

import "time"

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
