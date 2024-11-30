package golf

import (
	"testing"
	"time"
	"unicode/utf8"
)

func TestDurationInvalid(t *testing.T) {
	var a time.Duration

	ensureParserError(t, "cannot use empty flag string", func(t *testing.T, p *Parser) {
		p.DurationVar(&a, "", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen: \"-e\"", func(t *testing.T, p *Parser) {
		p.DurationVar(&a, "-e", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen: \"--example\"", func(t *testing.T, p *Parser) {
		p.DurationVar(&a, "--example", "some example flag")
	})
}

func TestParseDurationMissingArgument(t *testing.T) {
	var p Parser
	var a time.Duration
	var b time.Duration
	p.DurationVarP(&a, 't', "little", "little")
	p.DurationVarP(&b, 'T', "big", "big")

	t.Run("short", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t"}), "flag requires argument")
		if got, want := a, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little"}), "flag requires argument")
		if got, want := a, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseDurationShortOption(t *testing.T) {
	var p Parser
	var a time.Duration
	var b time.Duration
	p.DurationVar(&a, "t", "little")
	p.DurationVar(&b, "T", "big")

	t.Run("single option with space", func(t *testing.T) {
		if got, want := p.Parse([]string{"-t", "2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		if got, want := p.Parse([]string{"-t2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		if got, want := p.Parse([]string{"-t2m", "-T3h"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		if got, want := p.Parse([]string{"-T", "3h", "-t", "2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		if got, want := p.Parse([]string{"-T3h", "-t2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseDurationLongOption(t *testing.T) {
	var p Parser
	var a time.Duration
	var b time.Duration
	p.DurationVarP(&a, 't', "little", "little")
	p.DurationVarP(&b, 'T', "big", "big")

	t.Run("both options", func(t *testing.T) {
		if got, want := p.Parse([]string{"--little", "2m", "--big", "3h"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		if got, want := p.Parse([]string{"--big", "3h", "--little", "2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestDurationPInvalid(t *testing.T) {
	var a time.Duration

	ensureParserError(t, "cannot use flag with invalid rune", func(t *testing.T, p *Parser) {
		p.DurationVarP(&a, utf8.RuneError, "", "some example flag")
	})
	ensureParserError(t, "cannot use empty flag", func(t *testing.T, p *Parser) {
		p.DurationVarP(&a, 'b', "", "some example flag")
	})
	ensureParserError(t, "cannot use hyphen as a flag", func(t *testing.T, p *Parser) {
		p.DurationVarP(&a, '-', "example", "some example flag")
	})
	ensureParserError(t, "cannot use flag that starts with a hyphen", func(t *testing.T, p *Parser) {
		p.DurationVarP(&a, 'e', "--example", "some example flag")
	})
}

func TestParseDurationPMissingArgument(t *testing.T) {
	var p Parser
	var a time.Duration
	var b time.Duration
	p.DurationVarP(&a, 't', "little", "little")
	p.DurationVarP(&b, 'T', "big", "big")

	t.Run("short", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"-t"}), "flag requires argument:")
		if got, want := a, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		ensureError(t, p.Parse([]string{"--little"}), "flag requires argument:")
		if got, want := a, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseDurationPShortOption(t *testing.T) {
	var p Parser
	var a time.Duration
	var b time.Duration
	p.DurationVar(&a, "t", "little")
	p.DurationVar(&b, "T", "big")

	t.Run("single option with space", func(t *testing.T) {
		if got, want := p.Parse([]string{"-t", "2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		if got, want := p.Parse([]string{"-t2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		if got, want := p.Parse([]string{"-t2m", "-T3h"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		if got, want := p.Parse([]string{"-T", "3h", "-t", "2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		if got, want := p.Parse([]string{"-T3h", "-t2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseDurationPLongOption(t *testing.T) {
	var p Parser
	var a time.Duration
	var b time.Duration
	p.DurationVarP(&a, 't', "little", "little")
	p.DurationVarP(&b, 'T', "big", "big")

	t.Run("both options", func(t *testing.T) {
		if got, want := p.Parse([]string{"--little", "2m", "--big", "3h"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		if got, want := p.Parse([]string{"--big", "3h", "--little", "2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}
