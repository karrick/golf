package golf

import (
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
