package golf

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestStringInvalid(t *testing.T) {
	ensurePanic(t, "cannot use empty flag string", func() {
		_ = String("", "", "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"-e\"", func() {
		_ = String("-e", "", "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"--example\"", func() {
		_ = String("--example", "", "some example flag")
	})
}

func TestParseStringMissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := String("t", "", "little")
		b := String("T", "", "big")

		if got, want := parseArgs([]string{"-t"}), "flag requires argument: 't'"; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		resetParser()
		a := String("little", "", "little")
		b := String("big", "", "big")

		if got, want := parseArgs([]string{"--little"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseStringShortOption(t *testing.T) {
	t.Run("single option with space", func(t *testing.T) {
		resetParser()
		a := String("t", "", "little")
		b := String("T", "", "big")

		if got, want := parseArgs([]string{"-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		resetParser()
		a := String("t", "", "little")
		b := String("T", "", "big")

		if got, want := parseArgs([]string{"-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		resetParser()
		a := String("t", "", "little")
		b := String("T", "", "big")

		if got, want := parseArgs([]string{"-t13", "-T42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		resetParser()
		a := String("t", "", "little")
		b := String("T", "", "big")

		if got, want := parseArgs([]string{"-T", "42", "-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		resetParser()
		a := String("t", "", "little")
		b := String("T", "", "big")

		if got, want := parseArgs([]string{"-T42", "-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseStringLongOption(t *testing.T) {
	t.Run("both options", func(t *testing.T) {
		resetParser()
		a := String("little", "", "little")
		b := String("big", "", "big")

		if got, want := parseArgs([]string{"--little", "13", "--big", "42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		resetParser()
		a := String("little", "", "little")
		b := String("big", "", "big")

		if got, want := parseArgs([]string{"--big", "42", "--little", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestStringPInvalid(t *testing.T) {
	ensurePanic(t, "cannot use flag with invalid rune", func() {
		_ = StringP(utf8.RuneError, "", "default-value", "some example flag")
	})
	ensurePanic(t, "cannot use hyphen as a flag", func() {
		_ = StringP('-', "example", "default-value", "some example flag")
	})
	ensurePanic(t, "cannot use empty flag", func() {
		_ = StringP('b', "", "default-value", "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen", func() {
		_ = StringP('e', "--example", "default-value", "some example flag")
	})
}

func TestParseStringPMissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := StringP('t', "little", "", "little")
		b := StringP('T', "big", "", "big")

		if got, want := parseArgs([]string{"-t"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		resetParser()
		a := StringP('t', "little", "", "little")
		b := StringP('T', "big", "", "big")

		if got, want := parseArgs([]string{"--little"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseStringPShortOption(t *testing.T) {
	t.Run("single option with space", func(t *testing.T) {
		resetParser()
		a := StringP('t', "little", "", "little")
		b := StringP('T', "big", "", "big")

		if got, want := parseArgs([]string{"-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		resetParser()
		a := StringP('t', "little", "", "little")
		b := StringP('T', "big", "", "big")

		if got, want := parseArgs([]string{"-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		resetParser()
		a := StringP('t', "little", "", "little")
		b := StringP('T', "big", "", "big")

		if got, want := parseArgs([]string{"-t13", "-T42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		resetParser()
		a := StringP('t', "little", "", "little")
		b := StringP('T', "big", "", "big")

		if got, want := parseArgs([]string{"-T", "42", "-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		resetParser()
		a := StringP('t', "little", "", "little")
		b := StringP('T', "big", "", "big")

		if got, want := parseArgs([]string{"-T42", "-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseStringPLongOption(t *testing.T) {
	t.Run("both options", func(t *testing.T) {
		resetParser()
		a := StringP('t', "little", "", "little")
		b := StringP('T', "big", "", "big")

		if got, want := parseArgs([]string{"--little", "13", "--big", "42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		resetParser()
		a := StringP('t', "little", "", "little")
		b := StringP('T', "big", "", "big")

		if got, want := parseArgs([]string{"--big", "42", "--little", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, "13"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, "42"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseStringFunc(t *testing.T) {
	t.Run("callback called", func(t *testing.T) {
		resetParser()
		var cbArg *string
		opt := StringFunc("o", "", "some option", func(v string) error {
			cbArg = &v
			return nil
		})

		if got, want := *opt, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := parseArgs([]string{"-o", "test"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := *opt, "test"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := (cbArg == nil), false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := *cbArg, "test"; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
	t.Run("callback error", func(t *testing.T) {
		resetParser()
		cbErr := fmt.Errorf("failure is the only option")
		opt := StringFunc("o", "", "some option", func(v string) error {
			return cbErr
		})

		if got, want := *opt, ""; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := parseArgs([]string{"-o", "test"}), cbErr; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseStringFuncP(t *testing.T) {
	type Test struct {
		title  string
		parsed string
	}
	tests := []Test{{title: "short", parsed: "-o"}, {title: "long", parsed: "--option"}}
	for _, test := range tests {
		t.Run("callback called with "+test.title+" option", func(t *testing.T) {
			resetParser()
			var cbArg *string
			opt := StringFuncP('o', "option", "", "some option", func(v string) error {
				cbArg = &v
				return nil
			})

			if got, want := *opt, ""; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := parseArgs([]string{test.parsed, "test"}), error(nil); got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := *opt, "test"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := (cbArg == nil), false; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := *cbArg, "test"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
		t.Run("callback error with "+test.title+" option", func(t *testing.T) {
			resetParser()
			cbErr := fmt.Errorf("failure is the only option")
			opt := StringFuncP('o', "option", "", "some option", func(v string) error {
				return cbErr
			})

			if got, want := *opt, ""; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := parseArgs([]string{test.parsed, "test"}), cbErr; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	}
}
