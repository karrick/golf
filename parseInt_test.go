package golf

import (
	"testing"
	"unicode/utf8"
)

func TestIntInvalid(t *testing.T) {
	var b int

	ensureParserError(t, "cannot use empty flag string", func(t *testing.T, p *Parser) {
		p.IntVar(&b, "", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen: \"-e\"", func(t *testing.T, p *Parser) {
		p.IntVar(&b, "-e", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen: \"--example\"", func(t *testing.T, p *Parser) {
		p.IntVar(&b, "--example", "some example flag")
	})
}

func TestParseIntMissingArgument(t *testing.T) {
	var p Parser
	var a int
	var b int
	p.IntVarP(&a, 't', "little", "little")
	p.IntVarP(&b, 'T', "big", "big")

	t.Run("short", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t"}), "flag requires argument")
		if got, want := a, int(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little"}), "flag requires argument")
		if got, want := a, int(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseIntShortOption(t *testing.T) {
	var p Parser
	var a int
	var b int
	p.IntVar(&a, "t", "little")
	p.IntVar(&b, "T", "big")

	t.Run("single option with space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t", "13"}))
		if got, want := a, int(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13"}))
		if got, want := a, int(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13", "-T42"}))
		if got, want := a, int(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T", "42", "-t", "13"}))
		if got, want := a, int(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T42", "-t13"}))
		if got, want := a, int(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseIntLongOption(t *testing.T) {
	var p Parser
	var a int
	var b int
	p.IntVarP(&a, 't', "little", "little")
	p.IntVarP(&b, 'T', "big", "big")

	t.Run("both options", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little", "13", "--big", "42"}))
		if got, want := a, int(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--big", "42", "--little", "13"}))
		if got, want := a, int(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestIntPInvalid(t *testing.T) {
	var a int

	ensureParserError(t, "cannot use flag with invalid rune", func(t *testing.T, p *Parser) {
		p.IntVarP(&a, utf8.RuneError, "", "some example flag")
	})
	ensureParserError(t, "cannot use empty flag", func(t *testing.T, p *Parser) {
		p.IntVarP(&a, 'b', "", "some example flag")
	})
	ensureParserError(t, "cannot use hyphen as a flag", func(t *testing.T, p *Parser) {
		p.IntVarP(&a, '-', "example", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen", func(t *testing.T, p *Parser) {
		p.IntVarP(&a, 'e', "--example", "some example flag")
	})
}

func TestParseIntPMissingArgument(t *testing.T) {
	var p Parser
	var a int
	var b int
	p.IntVarP(&a, 't', "little", "little")
	p.IntVarP(&b, 'T', "big", "big")

	t.Run("short", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t"}), "flag requires argument")
		if got, want := a, int(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little"}), "flag requires argument")
		if got, want := a, int(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseIntPShortOption(t *testing.T) {
	var p Parser
	var a int
	var b int
	p.IntVar(&a, "t", "little")
	p.IntVar(&b, "T", "big")

	t.Run("single option with space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t", "13"}))
		if got, want := a, int(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13"}))
		if got, want := a, int(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13", "-T42"}))
		if got, want := a, int(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T", "42", "-t", "13"}))
		if got, want := a, int(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T42", "-t13"}))
		if got, want := a, int(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseIntPLongOption(t *testing.T) {
	var p Parser
	var a int
	var b int
	p.IntVarP(&a, 't', "little", "little")
	p.IntVarP(&b, 'T', "big", "big")

	t.Run("both options", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little", "13", "--big", "42"}))
		if got, want := a, int(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--big", "42", "--little", "13"}))
		if got, want := a, int(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}
