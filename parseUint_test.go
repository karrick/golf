package golf

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestUintInvalid(t *testing.T) {
	ensurePanic(t, "cannot use empty flag string", func() {
		_ = Uint("", 0, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"-e\"", func() {
		_ = Uint("-e", 0, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"--example\"", func() {
		_ = Uint("--example", 0, "some example flag")
	})
}

func TestParseUintMissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := Uint("t", 0, "little")
		b := Uint("T", 0, "big")

		if got, want := parseArgs([]string{"-t"}), "flag requires argument: 't'"; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		resetParser()
		a := Uint("little", 0, "little")
		b := Uint("big", 0, "big")

		if got, want := parseArgs([]string{"--little"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseUintShortOption(t *testing.T) {
	t.Run("single option with space", func(t *testing.T) {
		resetParser()
		a := Uint("t", 0, "little")
		b := Uint("T", 0, "big")

		if got, want := parseArgs([]string{"-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		resetParser()
		a := Uint("t", 0, "little")
		b := Uint("T", 0, "big")

		if got, want := parseArgs([]string{"-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		resetParser()
		a := Uint("t", 0, "little")
		b := Uint("T", 0, "big")

		if got, want := parseArgs([]string{"-t13", "-T42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		resetParser()
		a := Uint("t", 0, "little")
		b := Uint("T", 0, "big")

		if got, want := parseArgs([]string{"-T", "42", "-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		resetParser()
		a := Uint("t", 0, "little")
		b := Uint("T", 0, "big")

		if got, want := parseArgs([]string{"-T42", "-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseUintLongOption(t *testing.T) {
	t.Run("both options", func(t *testing.T) {
		resetParser()
		a := Uint("little", 0, "little")
		b := Uint("big", 0, "big")

		if got, want := parseArgs([]string{"--little", "13", "--big", "42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		resetParser()
		a := Uint("little", 0, "little")
		b := Uint("big", 0, "big")

		if got, want := parseArgs([]string{"--big", "42", "--little", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestUintPInvalid(t *testing.T) {
	ensurePanic(t, "cannot use flag with invalid rune", func() {
		_ = UintP(utf8.RuneError, "", 13, "some example flag")
	})
	ensurePanic(t, "cannot use hyphen as a flag", func() {
		_ = UintP('-', "example", 13, "some example flag")
	})
	ensurePanic(t, "cannot use empty flag", func() {
		_ = UintP('b', "", 13, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen", func() {
		_ = UintP('e', "--example", 13, "some example flag")
	})
}

func TestParseUintPMissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := UintP('t', "little", 0, "little")
		b := UintP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		resetParser()
		a := UintP('t', "little", 0, "little")
		b := UintP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--little"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseUintPShortOption(t *testing.T) {
	t.Run("single option with space", func(t *testing.T) {
		resetParser()
		a := UintP('t', "little", 0, "little")
		b := UintP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		resetParser()
		a := UintP('t', "little", 0, "little")
		b := UintP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		resetParser()
		a := UintP('t', "little", 0, "little")
		b := UintP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t13", "-T42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		resetParser()
		a := UintP('t', "little", 0, "little")
		b := UintP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-T", "42", "-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		resetParser()
		a := UintP('t', "little", 0, "little")
		b := UintP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-T42", "-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseUintPLongOption(t *testing.T) {
	t.Run("both options", func(t *testing.T) {
		resetParser()
		a := UintP('t', "little", 0, "little")
		b := UintP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--little", "13", "--big", "42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		resetParser()
		a := UintP('t', "little", 0, "little")
		b := UintP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--big", "42", "--little", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseUintFunc(t *testing.T) {
	t.Run("callback called", func(t *testing.T) {
		resetParser()
		var cbArg *uint
		opt := UintFunc("o", 0, "some option", func(v uint) error {
			cbArg = &v
			return nil
		})

		if got, want := *opt, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := parseArgs([]string{"-o", "12345"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := *opt, uint(12345); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := (cbArg == nil), false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := *cbArg, uint(12345); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
	t.Run("callback error", func(t *testing.T) {
		resetParser()
		cbErr := fmt.Errorf("failure is the only option")
		opt := UintFunc("o", 0, "some option", func(v uint) error {
			return cbErr
		})

		if got, want := *opt, uint(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := parseArgs([]string{"-o", "12345"}), cbErr; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseUintFuncP(t *testing.T) {
	type Test struct {
		title  string
		parsed string
	}
	tests := []Test{{title: "short", parsed: "-o"}, {title: "long", parsed: "--option"}}
	for _, test := range tests {
		t.Run("callback called with "+test.title+" option", func(t *testing.T) {
			resetParser()
			var cbArg *uint
			opt := UintFuncP('o', "option", 0, "some option", func(v uint) error {
				cbArg = &v
				return nil
			})

			if got, want := *opt, uint(0); got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := parseArgs([]string{test.parsed, "12345"}), error(nil); got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := *opt, uint(12345); got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := (cbArg == nil), false; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := *cbArg, uint(12345); got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
		t.Run("callback error with "+test.title+" option", func(t *testing.T) {
			resetParser()
			cbErr := fmt.Errorf("failure is the only option")
			opt := UintFuncP('o', "option", 0, "some option", func(v uint) error {
				return cbErr
			})

			if got, want := *opt, uint(0); got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := parseArgs([]string{test.parsed, "12345"}), cbErr; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	}
}
