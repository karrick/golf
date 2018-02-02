package golf

import (
	"testing"
	"unicode/utf8"
)

func TestUint64Invalid(t *testing.T) {
	ensurePanic(t, "cannot use empty flag string", func() {
		_ = Uint64("", 0, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"-e\"", func() {
		_ = Uint64("-e", 0, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"--example\"", func() {
		_ = Uint64("--example", 0, "some example flag")
	})
}

func TestParseUint64MissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := Uint64("t", 0, "little")
		b := Uint64("T", 0, "big")

		if got, want := parseArgs([]string{"-t"}), "flag requires argument: 't'"; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		resetParser()
		a := Uint64("little", 0, "little")
		b := Uint64("big", 0, "big")

		if got, want := parseArgs([]string{"--little"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseUint64ShortOption(t *testing.T) {
	t.Run("single option with space", func(t *testing.T) {
		resetParser()
		a := Uint64("t", 0, "little")
		b := Uint64("T", 0, "big")

		if got, want := parseArgs([]string{"-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		resetParser()
		a := Uint64("t", 0, "little")
		b := Uint64("T", 0, "big")

		if got, want := parseArgs([]string{"-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		resetParser()
		a := Uint64("t", 0, "little")
		b := Uint64("T", 0, "big")

		if got, want := parseArgs([]string{"-t13", "-T42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		resetParser()
		a := Uint64("t", 0, "little")
		b := Uint64("T", 0, "big")

		if got, want := parseArgs([]string{"-T", "42", "-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		resetParser()
		a := Uint64("t", 0, "little")
		b := Uint64("T", 0, "big")

		if got, want := parseArgs([]string{"-T42", "-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseUint64LongOption(t *testing.T) {
	t.Run("both options", func(t *testing.T) {
		resetParser()
		a := Uint64("little", 0, "little")
		b := Uint64("big", 0, "big")

		if got, want := parseArgs([]string{"--little", "13", "--big", "42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		resetParser()
		a := Uint64("little", 0, "little")
		b := Uint64("big", 0, "big")

		if got, want := parseArgs([]string{"--big", "42", "--little", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestUint64PInvalid(t *testing.T) {
	ensurePanic(t, "cannot use flag with invalid rune", func() {
		_ = Uint64P(utf8.RuneError, "", 13, "some example flag")
	})
	ensurePanic(t, "cannot use hyphen as a flag", func() {
		_ = Uint64P('-', "example", 13, "some example flag")
	})
	ensurePanic(t, "cannot use empty flag", func() {
		_ = Uint64P('b', "", 13, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen", func() {
		_ = Uint64P('e', "--example", 13, "some example flag")
	})
}

func TestParseUint64PMissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := Uint64P('t', "little", 0, "little")
		b := Uint64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		resetParser()
		a := Uint64P('t', "little", 0, "little")
		b := Uint64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--little"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseUint64PShortOption(t *testing.T) {
	t.Run("single option with space", func(t *testing.T) {
		resetParser()
		a := Uint64P('t', "little", 0, "little")
		b := Uint64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		resetParser()
		a := Uint64P('t', "little", 0, "little")
		b := Uint64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		resetParser()
		a := Uint64P('t', "little", 0, "little")
		b := Uint64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t13", "-T42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		resetParser()
		a := Uint64P('t', "little", 0, "little")
		b := Uint64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-T", "42", "-t", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		resetParser()
		a := Uint64P('t', "little", 0, "little")
		b := Uint64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-T42", "-t13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseUint64PLongOption(t *testing.T) {
	t.Run("both options", func(t *testing.T) {
		resetParser()
		a := Uint64P('t', "little", 0, "little")
		b := Uint64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--little", "13", "--big", "42"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		resetParser()
		a := Uint64P('t', "little", 0, "little")
		b := Uint64P('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--big", "42", "--little", "13"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, uint64(13); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, uint64(42); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}
