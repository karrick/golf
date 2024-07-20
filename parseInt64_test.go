package golf

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestInt64Invalid(t *testing.T) {
	ensurePanic(t, "cannot use empty flag string", func() {
		_ = Int64("", 0, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"-e\"", func() {
		_ = Int64("-e", 0, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"--example\"", func() {
		_ = Int64("--example", 0, "some example flag")
	})
}

func TestParseInt64MissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := Int64("t", 0, "little")
		b := Int64("T", 0, "big")

		if got, want := parseArgs([]string{"-t"}), "flag requires argument: 't'"; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		resetParser()
		a := Int64("little", 0, "little")
		b := Int64("big", 0, "big")

		if got, want := parseArgs([]string{"--little"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseInt64ShortOption(t *testing.T) {
	t.Run("single option with space", func(t *testing.T) {
		resetParser()
		a := Int64("t", 0, "little")
		b := Int64("T", 0, "big")

		if got, want := parseArgs([]string{"-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		resetParser()
		a := Int64("t", 0, "little")
		b := Int64("T", 0, "big")

		if got, want := parseArgs([]string{"-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		resetParser()
		a := Int64("t", 0, "little")
		b := Int64("T", 0, "big")

		if got, want := parseArgs([]string{"-t13", "-T42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		resetParser()
		a := Int64("t", 0, "little")
		b := Int64("T", 0, "big")

		if got, want := parseArgs([]string{"-T", "42", "-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		resetParser()
		a := Int64("t", 0, "little")
		b := Int64("T", 0, "big")

		if got, want := parseArgs([]string{"-T42", "-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseInt64LongOption(t *testing.T) {
	t.Run("both options", func(t *testing.T) {
		resetParser()
		a := Int64("little", 0, "little")
		b := Int64("big", 0, "big")

		if got, want := parseArgs([]string{"--little", "13", "--big", "42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		resetParser()
		a := Int64("little", 0, "little")
		b := Int64("big", 0, "big")

		if got, want := parseArgs([]string{"--big", "42", "--little", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestInt64PInvalid(t *testing.T) {
	ensurePanic(t, "cannot use flag with invalid rune", func() {
		_ = Int64P(utf8.RuneError, "", 13, "some example flag")
	})
	ensurePanic(t, "cannot use hyphen as a flag", func() {
		_ = Int64P('-', "example", 13, "some example flag")
	})
	ensurePanic(t, "cannot use empty flag", func() {
		_ = Int64P('b', "", 13, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen", func() {
		_ = Int64P('e', "--example", 13, "some example flag")
	})
}

func TestParseInt64PMissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := Int64P('t', "little", 0, "little")
		b := Int64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		resetParser()
		a := Int64P('t', "little", 0, "little")
		b := Int64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--little"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseInt64PShortOption(t *testing.T) {
	t.Run("single option with space", func(t *testing.T) {
		resetParser()
		a := Int64P('t', "little", 0, "little")
		b := Int64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		resetParser()
		a := Int64P('t', "little", 0, "little")
		b := Int64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		resetParser()
		a := Int64P('t', "little", 0, "little")
		b := Int64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t13", "-T42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		resetParser()
		a := Int64P('t', "little", 0, "little")
		b := Int64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-T", "42", "-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		resetParser()
		a := Int64P('t', "little", 0, "little")
		b := Int64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-T42", "-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseInt64PLongOption(t *testing.T) {
	t.Run("both options", func(t *testing.T) {
		resetParser()
		a := Int64P('t', "little", 0, "little")
		b := Int64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--little", "13", "--big", "42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		resetParser()
		a := Int64P('t', "little", 0, "little")
		b := Int64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--big", "42", "--little", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, int64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, int64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseInt64Func(t *testing.T) {
	t.Run("callback called", func(t *testing.T) {
		resetParser()
		var cbArg *int64
		opt := Int64Func("o", 0, "some option", func(v int64) error {
			cbArg = &v
			return nil
		})

		if got, want := *opt, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := parseArgs([]string{"-o", "12345"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := *opt, int64(12345); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := (cbArg == nil), false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := *cbArg, int64(12345); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
	t.Run("callback error", func(t *testing.T) {
		resetParser()
		cbErr := fmt.Errorf("failure is the only option")
		opt := Int64Func("o", 0, "some option", func(v int64) error {
			return cbErr
		})

		if got, want := *opt, int64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := parseArgs([]string{"-o", "12345"}), cbErr; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseInt64FuncP(t *testing.T) {
	type Test struct {
		title  string
		parsed string
	}
	tests := []Test{{title: "short", parsed: "-o"}, {title: "long", parsed: "--option"}}
	for _, test := range tests {
		t.Run("callback called with "+test.title+" option", func(t *testing.T) {
			resetParser()
			var cbArg *int64
			opt := Int64FuncP('o', "option", 0, "some option", func(v int64) error {
				cbArg = &v
				return nil
			})

			if got, want := *opt, int64(0); got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := parseArgs([]string{test.parsed, "12345"}), error(nil); got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := *opt, int64(12345); got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := (cbArg == nil), false; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := *cbArg, int64(12345); got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
		t.Run("callback error with "+test.title+" option", func(t *testing.T) {
			resetParser()
			cbErr := fmt.Errorf("failure is the only option")
			opt := Int64FuncP('o', "option", 0, "some option", func(v int64) error {
				return cbErr
			})

			if got, want := *opt, int64(0); got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := parseArgs([]string{test.parsed, "12345"}), cbErr; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	}
}
