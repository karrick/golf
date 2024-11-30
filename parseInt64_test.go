package golf

import (
	"testing"
	"unicode/utf8"
)

func TestInt64Invalid(t *testing.T) {
	var b int64

	ensureParserError(t, "cannot use empty flag string", func(t *testing.T, p *Parser) {
		p.Int64Var(&b, "", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen: \"-e\"", func(t *testing.T, p *Parser) {
		p.Int64Var(&b, "-e", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen: \"--example\"", func(t *testing.T, p *Parser) {
		p.Int64Var(&b, "--example", "some example flag")
	})
}

func TestParseInt64MissingArgument(t *testing.T) {
	var p Parser
	var a int64
	var b int64
	p.Int64VarP(&a, 't', "little", "little")
	p.Int64VarP(&b, 'T', "big", "big")

	t.Run("short", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t"}), "flag requires argument")
		if got, want := a, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little"}), "flag requires argument")
		if got, want := a, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseInt64ShortOption(t *testing.T) {
	var p Parser
	var a int64
	var b int64
	p.Int64Var(&a, "t", "little")
	p.Int64Var(&b, "T", "big")

	t.Run("single option with space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t", "13"}))
		if got, want := a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13"}))
		if got, want := a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13", "-T42"}))
		if got, want := a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T", "42", "-t", "13"}))
		if got, want := a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T42", "-t13"}))
		if got, want := a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseInt64LongOption(t *testing.T) {
	var p Parser
	var a int64
	var b int64
	p.Int64VarP(&a, 't', "little", "little")
	p.Int64VarP(&b, 'T', "big", "big")

	t.Run("both options", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little", "13", "--big", "42"}))
		if got, want := a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--big", "42", "--little", "13"}))
		if got, want := a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestInt64PInvalid(t *testing.T) {
	var a int64

	ensureParserError(t, "cannot use flag with invalid rune", func(t *testing.T, p *Parser) {
		p.Int64VarP(&a, utf8.RuneError, "", "some example flag")
	})
	ensureParserError(t, "cannot use empty flag", func(t *testing.T, p *Parser) {
		p.Int64VarP(&a, 'b', "", "some example flag")
	})
	ensureParserError(t, "cannot use hyphen as a flag", func(t *testing.T, p *Parser) {
		p.Int64VarP(&a, '-', "example", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen", func(t *testing.T, p *Parser) {
		p.Int64VarP(&a, 'e', "--example", "some example flag")
	})
}

func TestParseInt64PMissingArgument(t *testing.T) {
	var p Parser
	var a int64
	var b int64
	p.Int64VarP(&a, 't', "little", "little")
	p.Int64VarP(&b, 'T', "big", "big")

	t.Run("short", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t"}), "flag requires argument")
		if got, want := a, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little"}), "flag requires argument")
		if got, want := a, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseInt64PShortOption(t *testing.T) {
	var p Parser
	var a int64
	var b int64
	p.Int64Var(&a, "t", "little")
	p.Int64Var(&b, "T", "big")

	t.Run("single option with space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t", "13"}))
		if got, want := a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13"}))
		if got, want := a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t13", "-T42"}))
		if got, want := a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T", "42", "-t", "13"}))
		if got, want := a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-T42", "-t13"}))
		if got, want := a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseInt64PLongOption(t *testing.T) {
	var p Parser
	var a int64
	var b int64
	p.Int64VarP(&a, 't', "little", "little")
	p.Int64VarP(&b, 'T', "big", "big")

	t.Run("both options", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little", "13", "--big", "42"}))
		if got, want := a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--big", "42", "--little", "13"}))
		if got, want := a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}
