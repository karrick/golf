package golf

import (
	"testing"
	"unicode/utf8"
)

func TestBoolInvalid(t *testing.T) {
	var a bool

	ensureParserError(t, "cannot use empty flag string", func(t *testing.T, p *Parser) {
		p.WithBoolVar(&a, "", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen: \"-e\"", func(t *testing.T, p *Parser) {
		p.WithBoolVar(&a, "-e", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen: \"--example\"", func(t *testing.T, p *Parser) {
		p.WithBoolVar(&a, "--example", "some example flag")
	})
}

func TestParseBoolShortOption(t *testing.T) {
	var a bool
	var b bool
	var p Parser
	p.WithBoolVar(&a, "v", "print verbose info")
	p.WithBoolVar(&b, "V", "print version info")

	t.Run("single option", func(t *testing.T) {
		if got, want := p.Parse([]string{"-v"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options seperate", func(t *testing.T) {
		if got, want := p.Parse([]string{"-v", "-V"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options together 1", func(t *testing.T) {
		if got, want := p.Parse([]string{"-vV"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options together 2", func(t *testing.T) {
		if got, want := p.Parse([]string{"-Vv"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseBoolLongOption(t *testing.T) {
	var p Parser
	var a bool
	var b bool
	p.WithBoolVarP(&a, 'v', "verbose", "print verbose info")
	p.WithBoolVarP(&b, 'V', "version", "print version info")

	t.Run("single option", func(t *testing.T) {
		if got, want := p.Parse([]string{"--verbose"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options order 1", func(t *testing.T) {
		if got, want := p.Parse([]string{"--verbose", "--version"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options order 2", func(t *testing.T) {
		if got, want := p.Parse([]string{"--version", "--verbose"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestBoolPInvalid(t *testing.T) {
	var a bool

	ensureParserError(t, "cannot use flag with invalid rune", func(t *testing.T, p *Parser) {
		p.WithBoolVarP(&a, utf8.RuneError, "", "some example flag")
	})
	ensureParserError(t, "cannot use empty flag", func(t *testing.T, p *Parser) {
		p.WithBoolVarP(&a, 'b', "", "some example flag")
	})
	ensureParserError(t, "cannot use hyphen as a flag", func(t *testing.T, p *Parser) {
		p.WithBoolVarP(&a, '-', "example", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen", func(t *testing.T, p *Parser) {
		p.WithBoolVarP(&a, 'e', "--example", "some example flag")
	})
}

func TestParseBoolPShortOption(t *testing.T) {
	var p Parser
	var a bool
	var b bool
	p.WithBoolVarP(&a, 'v', "verbose", "print verbose info")
	p.WithBoolVarP(&b, 'V', "version", "print version info")

	t.Run("single option", func(t *testing.T) {
		if got, want := p.Parse([]string{"-v"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options seperate", func(t *testing.T) {
		if got, want := p.Parse([]string{"-v", "-V"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options together 1", func(t *testing.T) {
		if got, want := p.Parse([]string{"-vV"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options together 2", func(t *testing.T) {
		if got, want := p.Parse([]string{"-Vv"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseBoolPLongOption(t *testing.T) {
	var p Parser
	var a bool
	var b bool
	p.WithBoolVarP(&a, 'v', "verbose", "print verbose info")
	p.WithBoolVarP(&b, 'V', "version", "print version info")

	t.Run("single option", func(t *testing.T) {
		if got, want := p.Parse([]string{"--verbose"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options order 1", func(t *testing.T) {
		if got, want := p.Parse([]string{"--verbose", "--version"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options order 2", func(t *testing.T) {
		if got, want := p.Parse([]string{"--version", "--verbose"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}
