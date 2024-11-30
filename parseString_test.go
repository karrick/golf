package golf

import (
	"testing"
	"unicode/utf8"
)

func TestStringInvalid(t *testing.T) {
	var b string

	ensureParserError(t, "cannot use empty flag string", func(t *testing.T, p *Parser) {
		p.StringVar(&b, "", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen: \"-e\"", func(t *testing.T, p *Parser) {
		p.StringVar(&b, "-e", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen: \"--example\"", func(t *testing.T, p *Parser) {
		p.StringVar(&b, "--example", "some example flag")
	})
}

func TestParseStringMissingArgument(t *testing.T) {
	var p Parser
	var a string
	var b string
	p.StringVarP(&a, 't', "little", "little")
	p.StringVarP(&b, 'T', "big", "big")

	t.Run("short", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t"}), "flag requires argument")
		if got, want := a, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little"}), "flag requires argument")
		if got, want := a, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseStringShortOption(t *testing.T) {
	var p Parser
	var a string
	var b string
	p.StringVar(&a, "t", "little")
	p.StringVar(&b, "T", "big")

	t.Run("single option with space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t", "13"}))
		if got, want := a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13"}))
		if got, want := a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13", "-T42"}))
		if got, want := a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T", "42", "-t", "13"}))
		if got, want := a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T42", "-t13"}))
		if got, want := a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseStringLongOption(t *testing.T) {
	var p Parser
	var a string
	var b string
	p.StringVarP(&a, 't', "little", "little")
	p.StringVarP(&b, 'T', "big", "big")

	t.Run("both options", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little", "13", "--big", "42"}))
		if got, want := a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--big", "42", "--little", "13"}))
		if got, want := a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestStringPInvalid(t *testing.T) {
	var a string

	ensureParserError(t, "cannot use flag with invalid rune", func(t *testing.T, p *Parser) {
		p.StringVarP(&a, utf8.RuneError, "", "some example flag")
	})
	ensureParserError(t, "cannot use empty flag", func(t *testing.T, p *Parser) {
		p.StringVarP(&a, 'b', "", "some example flag")
	})
	ensureParserError(t, "cannot use hyphen as a flag", func(t *testing.T, p *Parser) {
		p.StringVarP(&a, '-', "example", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen", func(t *testing.T, p *Parser) {
		p.StringVarP(&a, 'e', "--example", "some example flag")
	})
}

func TestParseStringPMissingArgument(t *testing.T) {
	var p Parser
	var a string
	var b string
	p.StringVarP(&a, 't', "little", "little")
	p.StringVarP(&b, 'T', "big", "big")

	t.Run("short", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t"}), "flag requires argument")
		if got, want := a, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little"}), "flag requires argument")
		if got, want := a, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseStringPShortOption(t *testing.T) {
	var p Parser
	var a string
	var b string
	p.StringVar(&a, "t", "little")
	p.StringVar(&b, "T", "big")

	t.Run("single option with space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t", "13"}))
		if got, want := a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13"}))
		if got, want := a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13", "-T42"}))
		if got, want := a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T", "42", "-t", "13"}))
		if got, want := a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T42", "-t13"}))
		if got, want := a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseStringPLongOption(t *testing.T) {
	var p Parser
	var a string
	var b string
	p.StringVarP(&a, 't', "little", "little")
	p.StringVarP(&b, 'T', "big", "big")

	t.Run("both options", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little", "13", "--big", "42"}))
		if got, want := a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--big", "42", "--little", "13"}))
		if got, want := a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}
