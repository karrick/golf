package golf

import (
	"testing"
	"unicode/utf8"
)

func TestUintInvalid(t *testing.T) {
	var b uint

	ensureParserError(t, "cannot use empty flag string", func(t *testing.T, p *Parser) {
		p.WithUintVar(&b, "", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen: \"-e\"", func(t *testing.T, p *Parser) {
		p.WithUintVar(&b, "-e", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen: \"--example\"", func(t *testing.T, p *Parser) {
		p.WithUintVar(&b, "--example", "some example flag")
	})
}

func TestParseUintMissingArgument(t *testing.T) {
	var p Parser
	var a uint
	var b uint
	p.WithUintVarP(&a, 't', "little", "little")
	p.WithUintVarP(&b, 'T', "big", "big")

	t.Run("short", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t"}), "flag requires argument")
		if got, want := a, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little"}), "flag requires argument")
		if got, want := a, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseUintShortOption(t *testing.T) {
	var p Parser
	var a uint
	var b uint
	p.WithUintVar(&a, "t", "little")
	p.WithUintVar(&b, "T", "big")

	t.Run("single option with space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t", "13"}))
		if got, want := a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13"}))
		if got, want := a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13", "-T42"}))
		if got, want := a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T", "42", "-t", "13"}))
		if got, want := a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T42", "-t13"}))
		if got, want := a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseUintLongOption(t *testing.T) {
	var p Parser
	var a uint
	var b uint
	p.WithUintVarP(&a, 't', "little", "little")
	p.WithUintVarP(&b, 'T', "big", "big")

	t.Run("both options", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little", "13", "--big", "42"}))
		if got, want := a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--big", "42", "--little", "13"}))
		if got, want := a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestUintPInvalid(t *testing.T) {
	var a uint

	ensureParserError(t, "cannot use flag with invalid rune", func(t *testing.T, p *Parser) {
		p.WithUintVarP(&a, utf8.RuneError, "", "some example flag")
	})
	ensureParserError(t, "cannot use empty flag", func(t *testing.T, p *Parser) {
		p.WithUintVarP(&a, 'b', "", "some example flag")
	})
	ensureParserError(t, "cannot use hyphen as a flag", func(t *testing.T, p *Parser) {
		p.WithUintVarP(&a, '-', "example", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen", func(t *testing.T, p *Parser) {
		p.WithUintVarP(&a, 'e', "--example", "some example flag")
	})
}

func TestParseUintPMissingArgument(t *testing.T) {
	var p Parser
	var a uint
	var b uint
	p.WithUintVarP(&a, 't', "little", "little")
	p.WithUintVarP(&b, 'T', "big", "big")

	t.Run("short", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t"}), "flag requires argument")
		if got, want := a, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little"}), "flag requires argument")
		if got, want := a, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseUintPShortOption(t *testing.T) {
	var p Parser
	var a uint
	var b uint
	p.WithUintVar(&a, "t", "little")
	p.WithUintVar(&b, "T", "big")

	t.Run("single option with space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t", "13"}))
		if got, want := a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13"}))
		if got, want := a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13", "-T42"}))
		if got, want := a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T", "42", "-t", "13"}))
		if got, want := a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T42", "-t13"}))
		if got, want := a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseUintPLongOption(t *testing.T) {
	var p Parser
	var a uint
	var b uint
	p.WithUintVarP(&a, 't', "little", "little")
	p.WithUintVarP(&b, 'T', "big", "big")

	t.Run("both options", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little", "13", "--big", "42"}))
		if got, want := a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--big", "42", "--little", "13"}))
		if got, want := a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}
