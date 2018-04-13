package golf

import (
	"errors"
	"strings"
	"testing"
)

func ensurePanic(t *testing.T, errorMessage string, f func()) {
	t.Helper()
	defer func() {
		r := recover()
		if r == nil || !strings.Contains(r.(error).Error(), errorMessage) {
			t.Errorf("GOT: %v; WANT: %v", r, errorMessage)
		}
	}()
	f()
}

func TestParseEmpty(t *testing.T) {
	resetParser()
	if got, want := parseArgs(nil), error(nil); got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
}

func TestParseUnknownOption(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		got, want := parseArgs([]string{"-a"}), errors.New("unknown flag: 'a'")
		if got == nil || got.Error() != want.Error() {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		resetParser()
		got, want := parseArgs([]string{"--version"}), errors.New("unknown flag: \"version\"")
		if got == nil || got.Error() != want.Error() {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseComplex(t *testing.T) {
	resetParser()

	a := IntP('l', "limit", 0, "limit results")
	b := BoolP('v', "verbose", false, "print verbose info")
	c := StringP('s', "servers", "", "ask servers")

	if got, want := parseArgs([]string{"-l", "4", "-v", "-s", "host1,host2", "some", "other", "arguments"}), error(nil); got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	if got, want := *a, 4; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	if got, want := *b, true; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	if got, want := *c, "host1,host2"; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	if got, want := strings.Join(Args(), " "), "some other arguments"; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
}

func TestParseComplexSomeFlagsAfterArgs(t *testing.T) {
	resetParser()

	a := IntP('l', "limit", 0, "limit results")
	b := BoolP('v', "verbose", false, "print verbose info")
	c := StringP('s', "servers", "", "ask servers")

	if got, want := parseArgs([]string{"-l", "4", "some", "-v", "-s", "host1,host2", "other", "arguments"}), error(nil); got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	if got, want := *a, 4; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	if got, want := *b, true; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	if got, want := *c, "host1,host2"; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	if got, want := strings.Join(Args(), " "), "some other arguments"; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
}

func TestParseStopsAfterDoubleHyphen(t *testing.T) {
	resetParser()

	a := IntP('l', "limit", 0, "limit results")
	b := BoolP('v', "verbose", false, "print verbose info")

	if got, want := parseArgs([]string{"-l4", "some", "--", "--verbose", "other", "arguments"}), error(nil); got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	if got, want := *a, 4; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	if got, want := *b, false; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	if got, want := strings.Join(Args(), " "), "some"; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
}

func TestParseConfused(t *testing.T) {
	resetParser()

	a := IntP('l', "limit", 0, "limit results")
	b := BoolP('v', "verbose", false, "print verbose info")

	if got, want := parseArgs([]string{"-vl"}), "flag requires argument"; !strings.HasPrefix(got.Error(), want) {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	if got, want := *a, 0; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	if got, want := *b, true; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
}

func TestParseHyphenAfterShort(t *testing.T) {
	resetParser()

	a := IntP('l', "limit", 0, "limit results")
	b := BoolP('v', "verbose", false, "print verbose info")

	if got, want := parseArgs([]string{"-v-l"}), "unknown flag: '-'"; !strings.HasPrefix(got.Error(), want) {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	if got, want := *a, 0; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}

	if got, want := *b, true; got != want {
		t.Errorf("GOT: %v; WANT: %v", got, want)
	}
}

func TestPanicsWhenAttemptToRedefineFlag(t *testing.T) {
	ensurePanic(t, "cannot add option that duplicates short flag: 'f'", func() {
		resetParser()
		_ = UintP('f', "flubber", 0, "some example flag")
		_ = UintP('f', "blubber", 0, "some example flag")
	})

	ensurePanic(t, "cannot add option that duplicates long flag: \"flubber\"", func() {
		resetParser()
		_ = UintP('f', "flubber", 0, "some example flag")
		_ = UintP('b', "flubber", 0, "some example flag")
	})
}

func TestParseShortWithOptionAfterShortWithoutOptions(t *testing.T) {
	t.Run("with intervening space", func(t *testing.T) {
		resetParser()

		a := IntP('a', "alpha", 0, "some integer")
		b := BoolP('b', "bravo", false, "some bool")
		c := StringP('c', "charlie", "", "some string")
		d := BoolP('d', "delta", false, "some bool")

		if got, want := parseArgs([]string{"-ba", "13", "-c", "foo"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *c, "foo"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *d, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("without intervening space", func(t *testing.T) {
		resetParser()

		a := IntP('a', "alpha", 0, "some integer")
		b := BoolP('b', "bravo", false, "some bool")
		c := StringP('c', "charlie", "", "some string")
		d := BoolP('d', "delta", false, "some bool")

		if got, want := parseArgs([]string{"-ba13", "-c", "foo"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *c, "foo"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *d, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func ensureStringSlices(t *testing.T, a, b []string) {
	t.Helper()
	if a, b := len(a), len(b); a != b {
		t.Errorf("len(a): %v; len(b): %v", a, b)
		return
	}
	for i, s := range a {
		if s != b[i] {
			t.Errorf("a[%d]: %q; b[%d]: %v", i, a, i, b)
		}
	}
}

func TestParseArgs(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		// t.Skip()
		resetParser()

		if err := parseArgs(nil); err != nil {
			t.Error(t)
		}

		if got, want := argsProcessed, 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		ensureStringSlices(t, Args(), []string{})
	})

	t.Run("no flags but args", func(t *testing.T) {
		// t.Skip()
		resetParser()

		if err := parseArgs([]string{"foo", "bar"}); err != nil {
			t.Error(t)
		}

		if got, want := argsProcessed, 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		ensureStringSlices(t, Args(), []string{"foo", "bar"})
	})

	t.Run("some flags and no args", func(t *testing.T) {
		// t.Skip()
		resetParser()

		_ = Bool("b", false, "")
		_ = Int("i", 0, "")

		if err := parseArgs([]string{"-b", "-i", "13"}); err != nil {
			t.Error(t)
		}

		if got, want := argsProcessed, 3; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		ensureStringSlices(t, Args(), []string{})
	})

	t.Run("some flags and args", func(t *testing.T) {
		// t.Skip()
		resetParser()

		b := Bool("b", false, "")
		i := Int("i", -1, "")

		if err := parseArgs([]string{"-b", "-i", "13", "foo", "bar"}); err != nil {
			t.Error(t)
		}

		if got, want := *b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := *i, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := argsProcessed, 3; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		ensureStringSlices(t, Args(), []string{"foo", "bar"})
	})

	t.Run("some flags and args without spaces", func(t *testing.T) {
		// t.Skip()
		resetParser()

		b := Bool("b", false, "")
		i := Int("i", -1, "")

		if err := parseArgs([]string{"-b", "-i13", "foo", "bar"}); err != nil {
			t.Error(t)
		}

		if got, want := *b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := *i, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := argsProcessed, 2; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		ensureStringSlices(t, Args(), []string{"foo", "bar"})
	})

	t.Run("some flags with escaped runes and args", func(t *testing.T) {
		// t.Skip()
		resetParser()

		_ = Bool("b", false, "")
		s := String("s", "", "")

		if err := parseArgs([]string{"-b", "-s ", "foo", "bar"}); err != nil {
			t.Error(t)
		}

		if got, want := *s, " "; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := argsProcessed, 2; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		ensureStringSlices(t, Args(), []string{"foo", "bar"})
	})

	t.Run("some flags with escaped runes and args", func(t *testing.T) {
		resetParser()

		_ = Bool("b", false, "")
		s := String("s", "", "")

		if err := parseArgs([]string{"-b", "-s", " ", "foo", "bar"}); err != nil {
			t.Error(t)
		}

		if got, want := *s, " "; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := argsProcessed, 3; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		ensureStringSlices(t, Args(), []string{"foo", "bar"})
	})
}

func BenchmarkBool(b *testing.B) {
	resetParser()
	optUnfold := BoolP('u', "unfold", false, "unfold")
	optDelimiter := StringP('d', "delimiter", "\n", "delimiter")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseArgs([]string{"-u"})
		if *optUnfold {
			*optDelimiter = ","
		}
	}
}

func BenchmarkString(b *testing.B) {
	resetParser()
	optUnfold := BoolP('u', "unfold", false, "unfold")
	optDelimiter := StringP('d', "delimiter", "\n", "delimiter")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseArgs([]string{"-d,"})
		if *optUnfold {
			*optDelimiter = ","
		}
	}
}
