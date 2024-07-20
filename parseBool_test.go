package golf

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestBoolInvalid(t *testing.T) {
	ensurePanic(t, "cannot use empty flag string", func() {
		_ = Bool("", false, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"-e\"", func() {
		_ = Bool("-e", false, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"--example\"", func() {
		_ = Bool("--example", false, "some example flag")
	})
}

func TestParseBoolShortOption(t *testing.T) {
	t.Run("single option", func(t *testing.T) {
		resetParser()
		a := Bool("v", false, "print verbose info")
		b := Bool("V", false, "print version info")

		if got, want := parseArgs([]string{"-v"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options seperate", func(t *testing.T) {
		resetParser()
		a := Bool("v", false, "print verbose info")
		b := Bool("V", false, "print version info")

		if got, want := parseArgs([]string{"-v", "-V"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options together 1", func(t *testing.T) {
		resetParser()
		a := Bool("v", false, "print verbose info")
		b := Bool("V", false, "print version info")

		if got, want := parseArgs([]string{"-vV"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options together 2", func(t *testing.T) {
		resetParser()
		a := Bool("v", false, "print verbose info")
		b := Bool("V", false, "print version info")

		if got, want := parseArgs([]string{"-Vv"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseBoolLongOption(t *testing.T) {
	t.Run("single option", func(t *testing.T) {
		resetParser()
		a := Bool("verbose", false, "print verbose info")
		b := Bool("version", false, "print version info")

		if got, want := parseArgs([]string{"--verbose"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options order 1", func(t *testing.T) {
		resetParser()
		a := Bool("verbose", false, "print verbose info")
		b := Bool("version", false, "print version info")

		if got, want := parseArgs([]string{"--verbose", "--version"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options order 2", func(t *testing.T) {
		resetParser()
		a := Bool("verbose", false, "print verbose info")
		b := Bool("version", false, "print version info")

		if got, want := parseArgs([]string{"--version", "--verbose"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestBoolPInvalid(t *testing.T) {
	ensurePanic(t, "cannot use flag with invalid rune", func() {
		_ = BoolP(utf8.RuneError, "", false, "some example flag")
	})
	ensurePanic(t, "cannot use empty flag", func() {
		_ = BoolP('b', "", false, "some example flag")
	})
	ensurePanic(t, "cannot use hyphen as a flag", func() {
		_ = BoolP('-', "example", false, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen", func() {
		_ = BoolP('e', "--example", false, "some example flag")
	})
}

func TestParseBoolPShortOption(t *testing.T) {
	t.Run("single option", func(t *testing.T) {
		resetParser()
		a := BoolP('v', "verbose", false, "print verbose info")
		b := BoolP('V', "version", false, "print version info")

		if got, want := parseArgs([]string{"-v"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options seperate", func(t *testing.T) {
		resetParser()
		a := BoolP('v', "verbose", false, "print verbose info")
		b := BoolP('V', "version", false, "print version info")

		if got, want := parseArgs([]string{"-v", "-V"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options together 1", func(t *testing.T) {
		resetParser()
		a := BoolP('v', "verbose", false, "print verbose info")
		b := BoolP('V', "version", false, "print version info")

		if got, want := parseArgs([]string{"-vV"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options together 2", func(t *testing.T) {
		resetParser()
		a := BoolP('v', "verbose", false, "print verbose info")
		b := BoolP('V', "version", false, "print version info")

		if got, want := parseArgs([]string{"-Vv"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseBoolPLongOption(t *testing.T) {
	t.Run("single option", func(t *testing.T) {
		resetParser()
		a := BoolP('v', "verbose", false, "print verbose info")
		b := BoolP('V', "version", false, "print version info")

		if got, want := parseArgs([]string{"--verbose"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options order 1", func(t *testing.T) {
		resetParser()
		a := BoolP('v', "verbose", false, "print verbose info")
		b := BoolP('V', "version", false, "print version info")

		if got, want := parseArgs([]string{"--verbose", "--version"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options order 2", func(t *testing.T) {
		resetParser()
		a := BoolP('v', "verbose", false, "print verbose info")
		b := BoolP('V', "version", false, "print version info")

		if got, want := parseArgs([]string{"--version", "--verbose"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseBoolDefault(t *testing.T) {
	t.Run("default to true", func(t *testing.T) {
		resetParser()
		a := BoolP('v', "verbose", true, "print verbose info")

		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := parseArgs([]string{"-v"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := *a, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
	t.Run("default to false", func(t *testing.T) {
		resetParser()
		a := BoolP('v', "verbose", false, "print verbose info")

		if got, want := *a, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := parseArgs([]string{"-v"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := *a, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseBoolFunc(t *testing.T) {
	t.Run("callback called", func(t *testing.T) {
		resetParser()
		var cbArg *bool
		opt := BoolFunc("o", false, "some option", func(v bool) error {
			cbArg = &v
			return nil
		})

		if got, want := *opt, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := parseArgs([]string{"-o"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := *opt, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := (cbArg == nil), false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := *cbArg, true; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
	t.Run("callback error", func(t *testing.T) {
		resetParser()
		cbErr := fmt.Errorf("failure is the only option")
		opt := BoolFunc("o", false, "some option", func(v bool) error {
			return cbErr
		})

		if got, want := *opt, false; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := parseArgs([]string{"-o"}), cbErr; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseBoolFuncP(t *testing.T) {
	type Test struct {
		title  string
		parsed string
	}
	tests := []Test{{title: "short", parsed: "-o"}, {title: "long", parsed: "--option"}}
	for _, test := range tests {
		t.Run("callback called with "+test.title+" option", func(t *testing.T) {
			resetParser()
			var cbArg *bool
			opt := BoolFuncP('o', "option", false, "some option", func(v bool) error {
				cbArg = &v
				return nil
			})

			if got, want := *opt, false; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := parseArgs([]string{test.parsed}), error(nil); got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := *opt, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := (cbArg == nil), false; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := *cbArg, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
		t.Run("callback error with "+test.title+" option", func(t *testing.T) {
			resetParser()
			cbErr := fmt.Errorf("failure is the only option")
			opt := BoolFuncP('o', "option", false, "some option", func(v bool) error {
				return cbErr
			})

			if got, want := *opt, false; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := parseArgs([]string{test.parsed}), cbErr; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	}
}
