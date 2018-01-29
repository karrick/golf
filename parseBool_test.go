package golf

import "testing"

func TestBoolInvalid(t *testing.T) {
	ensurePanic(t, "cannot add option without either short, long, or both flags", func() {
		_ = Bool("", "", false, "some example flag")
	})
	ensurePanic(t, "cannot set short flag to a hyphen: \"-e\"", func() {
		_ = Bool("-e", "example", false, "some example flag")
	})
	ensurePanic(t, "cannot start long flag with a hyphen: \"--example\"", func() {
		_ = Bool("e", "--example", false, "some example flag")
	})
}

func TestParseBoolShortOption(t *testing.T) {
	t.Run("single option", func(t *testing.T) {
		resetParser()
		a := Bool("v", "verbose", false, "print verbose info")
		b := Bool("V", "version", false, "print version info")

		if got, want := parse("-v"), error(nil); got != want {
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
		a := Bool("v", "verbose", false, "print verbose info")
		b := Bool("V", "version", false, "print version info")

		if got, want := parse("-v -V"), error(nil); got != want {
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
		a := Bool("v", "verbose", false, "print verbose info")
		b := Bool("V", "version", false, "print version info")

		if got, want := parse("-vV"), error(nil); got != want {
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
		a := Bool("v", "verbose", false, "print verbose info")
		b := Bool("V", "version", false, "print version info")

		if got, want := parse("-Vv"), error(nil); got != want {
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
		a := Bool("", "verbose", false, "print verbose info")
		b := Bool("", "version", false, "print version info")

		if got, want := parse("--verbose"), error(nil); got != want {
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
		a := Bool("", "verbose", false, "print verbose info")
		b := Bool("", "version", false, "print version info")

		if got, want := parse("--verbose --version"), error(nil); got != want {
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
		a := Bool("", "verbose", false, "print verbose info")
		b := Bool("", "version", false, "print version info")

		if got, want := parse("--version --verbose"), error(nil); got != want {
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
