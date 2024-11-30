package golf

import (
	"math"
	"testing"
	"unicode/utf8"
)

func TestFloatInvalid(t *testing.T) {
	var b float64

	ensureParserError(t, "cannot use empty flag string", func(t *testing.T, p *Parser) {
		p.FloatVar(&b, "", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen: \"-e\"", func(t *testing.T, p *Parser) {
		p.FloatVar(&b, "-e", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen: \"--example\"", func(t *testing.T, p *Parser) {
		p.FloatVar(&b, "--example", "some example flag")
	})
}

func TestParseFloatMissingArgument(t *testing.T) {
	var p Parser
	var a float64
	var b float64
	p.FloatVarP(&a, 't', "little", "little")
	p.FloatVarP(&b, 'T', "big", "big")

	t.Run("short", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t"}), "flag requires argument")
		if got, want := a, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little"}), "flag requires argument")
		if got, want := a, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseFloatShortOption(t *testing.T) {
	var p Parser
	a := math.NaN()
	var b float64
	p.FloatVar(&a, "t", "little")
	p.FloatVar(&b, "T", "big")

	t.Run("single option with space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t", "3.14"}))
		if got, want := a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t3.14"}))
		if got, want := a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t3.14", "-T2.78"}))
		if got, want := a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T", "2.78", "-t", "3.14"}))
		if got, want := a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T2.78", "-t3.14"}))
		if got, want := a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseFloatLongOption(t *testing.T) {
	var p Parser
	var a float64
	var b float64
	p.FloatVarP(&a, 't', "little", "little")
	p.FloatVarP(&b, 'T', "big", "big")

	t.Run("both options", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little", "3.14", "--big", "2.78"}))
		if got, want := a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--big", "2.78", "--little", "3.14"}))
		if got, want := a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestFloatPInvalid(t *testing.T) {
	var a float64

	ensureParserError(t, "cannot use flag with invalid rune", func(t *testing.T, p *Parser) {
		p.FloatVarP(&a, utf8.RuneError, "", "some example flag")
	})
	ensureParserError(t, "cannot use empty flag", func(t *testing.T, p *Parser) {
		p.FloatVarP(&a, 'b', "", "some example flag")
	})
	ensureParserError(t, "cannot use hyphen as a flag", func(t *testing.T, p *Parser) {
		p.FloatVarP(&a, '-', "example", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen", func(t *testing.T, p *Parser) {
		p.FloatVarP(&a, 'e', "--example", "some example flag")
	})
}

func TestParseFloatPMissingArgument(t *testing.T) {
	var p Parser
	var a float64
	var b float64
	p.FloatVarP(&a, 't', "little", "little")
	p.FloatVarP(&b, 'T', "big", "big")

	t.Run("short", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t"}), "flag requires argument")
		if got, want := a, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little"}), "flag requires argument")
		if got, want := a, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseFloatPShortOption(t *testing.T) {
	var p Parser
	a := math.NaN()
	var b float64
	p.FloatVar(&a, "t", "little")
	p.FloatVar(&b, "T", "big")

	t.Run("single option with space", func(t *testing.T) {
		if got, want := p.Parse([]string{"-t", "3.14"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		if got, want := p.Parse([]string{"-t3.14"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		if got, want := p.Parse([]string{"-t3.14", "-T2.78"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T", "2.78", "-t", "3.14"}))
		if got, want := a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T2.78", "-t3.14"}))
		if got, want := a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseFloatPLongOption(t *testing.T) {
	var p Parser
	var a float64
	var b float64
	p.FloatVarP(&a, 't', "little", "little")
	p.FloatVarP(&b, 'T', "big", "big")

	t.Run("both options", func(t *testing.T) {
		if got, want := p.Parse([]string{"--little", "3.14", "--big", "2.78"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		if got, want := p.Parse([]string{"--big", "2.78", "--little", "3.14"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}
