package golf

import (
	"testing"
	"unicode/utf8"
)

func TestIntInvalid(t *testing.T) {
	ensurePanic(t, "cannot use empty flag string", func() {
		_ = Int("", 0, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"-e\"", func() {
		_ = Int("-e", 0, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"--example\"", func() {
		_ = Int("--example", 0, "some example flag")
	})
}

func TestParseIntMissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := Int("t", 0, "little")
		b := Int("T", 0, "big")

		if got, want := parseArgs([]string{"-t"}), "flag requires argument: 't'"; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		resetParser()
		a := Int("little", 0, "little")
		b := Int("big", 0, "big")

		if got, want := parseArgs([]string{"--little"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseIntShortOption(t *testing.T) {
	t.Run("single option with space", func(t *testing.T) {
		resetParser()
		a := Int("t", 0, "little")
		b := Int("T", 0, "big")

		if got, want := parseArgs([]string{"-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		resetParser()
		a := Int("t", 0, "little")
		b := Int("T", 0, "big")

		if got, want := parseArgs([]string{"-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		resetParser()
		a := Int("t", 0, "little")
		b := Int("T", 0, "big")

		if got, want := parseArgs([]string{"-t13", "-T42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 42; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		resetParser()
		a := Int("t", 0, "little")
		b := Int("T", 0, "big")

		if got, want := parseArgs([]string{"-T", "42", "-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 42; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		resetParser()
		a := Int("t", 0, "little")
		b := Int("T", 0, "big")

		if got, want := parseArgs([]string{"-T42", "-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 42; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseIntLongOption(t *testing.T) {
	t.Run("both options", func(t *testing.T) {
		resetParser()
		a := Int("little", 0, "little")
		b := Int("big", 0, "big")

		if got, want := parseArgs([]string{"--little", "13", "--big", "42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 42; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		resetParser()
		a := Int("little", 0, "little")
		b := Int("big", 0, "big")

		if got, want := parseArgs([]string{"--big", "42", "--little", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 42; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestIntPInvalid(t *testing.T) {
	ensurePanic(t, "cannot use flag with invalid rune", func() {
		_ = IntP(utf8.RuneError, "", 13, "some example flag")
	})
	ensurePanic(t, "cannot use hyphen as a flag", func() {
		_ = IntP('-', "example", 13, "some example flag")
	})
	ensurePanic(t, "cannot use empty flag", func() {
		_ = IntP('b', "", 13, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen", func() {
		_ = IntP('e', "--example", 13, "some example flag")
	})
}

func TestParseIntPMissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := IntP('t', "little", 0, "little")
		b := IntP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		resetParser()
		a := IntP('t', "little", 0, "little")
		b := IntP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--little"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseIntPShortOption(t *testing.T) {
	t.Run("single option with space", func(t *testing.T) {
		resetParser()
		a := IntP('t', "little", 0, "little")
		b := IntP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		resetParser()
		a := IntP('t', "little", 0, "little")
		b := IntP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		resetParser()
		a := IntP('t', "little", 0, "little")
		b := IntP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t13", "-T42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 42; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		resetParser()
		a := IntP('t', "little", 0, "little")
		b := IntP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-T", "42", "-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 42; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		resetParser()
		a := IntP('t', "little", 0, "little")
		b := IntP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-T42", "-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 42; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseIntPLongOption(t *testing.T) {
	t.Run("both options", func(t *testing.T) {
		resetParser()
		a := IntP('t', "little", 0, "little")
		b := IntP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--little", "13", "--big", "42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 42; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		resetParser()
		a := IntP('t', "little", 0, "little")
		b := IntP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--big", "42", "--little", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 13; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 42; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}
