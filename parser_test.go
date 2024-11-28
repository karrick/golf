package golf

import "testing"

func ensureParserError(t *testing.T, description string, callback func(t *testing.T, p *Parser)) {
	t.Helper()
	t.Run(description, func(t *testing.T) {
		t.Helper()
		var p Parser
		callback(t, &p)
		ensureError(t, p.Err(), description)
	})
}

func TestParseEmpty(t *testing.T) {
	var p Parser
	ensureError(t, p.Parse(nil))
}

func TestParseComplex(t *testing.T) {
	var b bool
	var i int
	var s string
	var p Parser

	p.
		WithBoolVarP(&b, 'v', "verbose", "print verbose info").
		WithIntVarP(&i, 'l', "limit", "limit results").
		WithStringVarP(&s, 's', "servers", "ask servers")

	err := p.Parse([]string{"-l", "4", "-v", "-s", "host1,host2", "some", "other", "arguments"})
	ensureError(t, err)

	if got, want := b, true; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
	if got, want := i, 4; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
	if got, want := s, "host1,host2"; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	ensureStringSlicesMatch(t, p.Args(), []string{"some", "other", "arguments"})

	if got, want := p.NArg(), 3; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
	if got, want := p.NFlag(), 5; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
}

func TestParseComplexSomeFlagsAfterArgs(t *testing.T) {
	var b bool
	var i int
	var s string
	var p Parser

	p.
		WithBoolVarP(&b, 'v', "verbose", "print verbose info").
		WithIntVarP(&i, 'l', "limit", "limit results").
		WithStringVarP(&s, 's', "servers", "ask servers")

	err := p.Parse([]string{"-l", "4", "some", "-v", "other", "-s", "host1,host2", "arguments"})
	ensureError(t, err)

	if got, want := b, true; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
	if got, want := i, 4; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
	if got, want := s, "host1,host2"; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	ensureStringSlicesMatch(t, p.Args(), []string{"some", "other", "arguments"})

	if got, want := p.NArg(), 3; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
	if got, want := p.NFlag(), 5; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
}

func TestParseStopsAfterDoubleHyphen(t *testing.T) {
	var b bool
	var i int
	var s string
	var p Parser

	p.
		WithBoolVarP(&b, 'v', "verbose", "print verbose info").
		WithIntVarP(&i, 'l', "limit", "limit results").
		WithStringVarP(&s, 's', "servers", "ask servers")

	err := p.Parse([]string{"--limit", "4", "--verbose", "--", "--servers", "host1,host2", "some", "other", "arguments"})
	ensureError(t, err)

	if got, want := b, true; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
	if got, want := i, 4; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
	if got, want := s, ""; got != want {
		t.Errorf("GOT: %q; WANT: %q", got, want)
	}

	ensureStringSlicesMatch(t, p.Args(), []string{"--servers", "host1,host2", "some", "other", "arguments"})

	if got, want := p.NArg(), 5; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
	if got, want := p.NFlag(), 4; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
}

func TestParseConfused(t *testing.T) {
	var b bool
	var i int
	var s string
	var p Parser

	p.
		WithBoolVarP(&b, 'v', "verbose", "print verbose info").
		WithIntVarP(&i, 'l', "limit", "limit results").
		WithStringVarP(&s, 's', "servers", "ask servers")

	err := p.Parse([]string{"-vs"})
	ensureError(t, err, "flag requires argument")

	if got, want := b, true; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
	if got, want := i, 0; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	ensureStringSlicesMatch(t, p.Args(), nil)

	if got, want := p.NArg(), 0; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
	if got, want := p.NFlag(), 1; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
}

func TestParseHyphenAfterShort(t *testing.T) {
	var b bool
	var i int
	var s string
	var p Parser

	p.
		WithBoolVarP(&b, 'v', "verbose", "print verbose info").
		WithIntVarP(&i, 'l', "limit", "limit results").
		WithStringVarP(&s, 's', "servers", "ask servers")

	err := p.Parse([]string{"-v-l"})
	ensureError(t, err, "unknown flag: '-'")

	if got, want := b, true; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
	if got, want := i, 0; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	ensureStringSlicesMatch(t, p.Args(), []string{"-v-l"})

	if got, want := p.NArg(), 1; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
	if got, want := p.NFlag(), 0; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
}

func TestPanicsWhenAttemptToRedefineFlag(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		var a, b uint
		var p Parser

		p.WithUintVarP(&a, 'f', "flubber", "example one")
		ensureError(t, p.Err())
		p.WithUintVarP(&b, 'f', "blubber", "example two")
		ensureError(t, p.Err(), "cannot add option that duplicates short flag")
	})

	t.Run("long", func(t *testing.T) {
		var a, b uint
		var p Parser

		p.WithUintVarP(&a, 'f', "flubber", "example one")
		ensureError(t, p.Err())
		p.WithUintVarP(&b, 'b', "flubber", "example two")
		ensureError(t, p.Err(), "cannot add option that duplicates long flag")
	})
}

func TestParseShortWithOptionAfterShortWithoutOptions(t *testing.T) {
	t.Run("with intervening space", func(t *testing.T) {
		var a int
		var b bool
		var c string
		var d bool
		var p Parser

		p.
			WithIntVarP(&a, 'a', "alpha", "some integer").
			WithBoolVarP(&b, 'b', "bravo", "some bool").
			WithStringVarP(&c, 'c', "string", "some string").
			WithBoolVarP(&d, 'd', "delta", "some bool")

		err := p.Parse([]string{"-ba", "13", "-c", "foo"})
		ensureError(t, err)

		if got, want := a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := c, "foo"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := d, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		ensureStringSlicesMatch(t, p.Args(), nil)

		if got, want := p.NArg(), 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := p.NFlag(), 4; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("without intervening space", func(t *testing.T) {
		var a int
		var b bool
		var c string
		var d bool
		var p Parser

		p.
			WithIntVarP(&a, 'a', "alpha", "some integer").
			WithBoolVarP(&b, 'b', "bravo", "some bool").
			WithStringVarP(&c, 'c', "string", "some string").
			WithBoolVarP(&d, 'd', "delta", "some bool")

		err := p.Parse([]string{"-ba13", "-c", "foo"})
		ensureError(t, err)

		if got, want := a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := c, "foo"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := d, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		ensureStringSlicesMatch(t, p.Args(), nil)

		if got, want := p.NArg(), 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := p.NFlag(), 3; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseArgs(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var p Parser

		ensureError(t, p.Parse(nil))

		ensureStringSlicesMatch(t, p.Args(), nil)

		if got, want := p.NArg(), 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := p.NFlag(), 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("no flags but args", func(t *testing.T) {
		var p Parser

		ensureError(t, p.Parse([]string{"foo", "bar"}))

		ensureStringSlicesMatch(t, p.Args(), []string{"foo", "bar"})

		if got, want := p.NArg(), 2; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := p.NFlag(), 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("some flags and no args", func(t *testing.T) {
		var b bool
		var i int
		var p Parser

		p.
			WithBoolVar(&b, "b", "").
			WithIntVar(&i, "i", "")

		ensureError(t, p.Parse([]string{"-b", "-i", "13"}))

		ensureStringSlicesMatch(t, p.Args(), nil)

		if got, want := p.NArg(), 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := p.NFlag(), 3; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("some flags and args", func(t *testing.T) {
		var b bool
		var i int
		var p Parser

		p.
			WithBoolVar(&b, "b", "").
			WithIntVar(&i, "i", "")

		ensureError(t, p.Parse([]string{"-b", "-i", "13", "foo", "bar"}))

		if got, want := b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := i, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		ensureStringSlicesMatch(t, p.Args(), []string{"foo", "bar"})

		if got, want := p.NArg(), 2; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := p.NFlag(), 3; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("some flags and args without spaces", func(t *testing.T) {
		var b bool
		var i int
		var p Parser

		p.
			WithBoolVar(&b, "b", "").
			WithIntVar(&i, "i", "")

		ensureError(t, p.Parse([]string{"-b", "-i13", "foo", "bar"}))

		if got, want := b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := i, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		ensureStringSlicesMatch(t, p.Args(), []string{"foo", "bar"})

		if got, want := p.NArg(), 2; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := p.NFlag(), 2; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("some flags with escaped runes and args", func(t *testing.T) {
		var b bool
		var s string
		var p Parser

		p.
			WithBoolVar(&b, "b", "").
			WithStringVar(&s, "s", "")

		ensureError(t, p.Parse([]string{"-b", "-s ", "foo", "bar"}))

		if got, want := s, " "; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		ensureStringSlicesMatch(t, p.Args(), []string{"foo", "bar"})

		if got, want := p.NArg(), 2; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := p.NFlag(), 2; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("some flags with escaped runes and args", func(t *testing.T) {
		var b bool
		var s string
		var p Parser

		p.
			WithBoolVar(&b, "b", "").
			WithStringVar(&s, "s", "")

		ensureError(t, p.Parse([]string{"-b", "-s", " ", "foo", "bar"}))

		if got, want := s, " "; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		ensureStringSlicesMatch(t, p.Args(), []string{"foo", "bar"})

		if got, want := p.NArg(), 2; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := p.NFlag(), 3; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func BenchmarkBool(b *testing.B) {
	var u bool
	d := "\n"
	var p Parser
	p.
		WithBoolVarP(&u, 'u', "unfold", "unfold").
		WithStringVarP(&d, 'd', "delimiter", "delimiter")

	b.ResetTimer()

	var count int

	for i := 0; i < b.N; i++ {
		err := p.Parse([]string{"-u"})
		if err != nil {
			count++
		}
		if u {
			d = ","
		}
	}

	if got, want := count, 0; got != want {
		b.Errorf("GOT: %v; WANT: %v:", got, want)
	}
}

func BenchmarkString(b *testing.B) {
	var u bool
	d := "\n"
	var p Parser

	p.
		WithBoolVarP(&u, 'u', "unfold", "unfold").
		WithStringVarP(&d, 'd', "delimiter", "delimiter")

	b.ResetTimer()

	var count int

	for i := 0; i < b.N; i++ {
		err := p.Parse([]string{"-u"})
		if err != nil {
			count++
		}
		if u {
			d = ","
		}
	}

	if got, want := count, 0; got != want {
		b.Errorf("GOT: %v; WANT: %v:", got, want)
	}
}
