package golf

import (
	"math"
	"testing"
	"unicode/utf8"
)

func TestFloatInvalid(t *testing.T) {
	ensurePanic(t, "cannot use empty flag string", func() {
		_ = Float("", 0, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"-e\"", func() {
		_ = Float("-e", 0, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"--example\"", func() {
		_ = Float("--example", 0, "some example flag")
	})
}

func TestParseFloatMissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := Float("t", 0, "little")
		b := Float("T", 0, "big")

		if got, want := parseArgs([]string{"-t"}), "flag requires argument: 't'"; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		resetParser()
		a := Float("little", 0, "little")
		b := Float("big", 0, "big")

		if got, want := parseArgs([]string{"--little"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseFloatShortOption(t *testing.T) {
	t.Run("single option with space", func(t *testing.T) {
		resetParser()
		a := Float("t", math.NaN(), "little")
		b := Float("T", 0, "big")

		if got, want := parseArgs([]string{"-t", "3.14"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		resetParser()
		a := Float("t", 0, "little")
		b := Float("T", 0, "big")

		if got, want := parseArgs([]string{"-t3.14"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		resetParser()
		a := Float("t", 0, "little")
		b := Float("T", 0, "big")

		if got, want := parseArgs([]string{"-t3.14", "-T2.78"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		resetParser()
		a := Float("t", 0, "little")
		b := Float("T", 0, "big")

		if got, want := parseArgs([]string{"-T", "2.78", "-t", "3.14"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		resetParser()
		a := Float("t", 0, "little")
		b := Float("T", 0, "big")

		if got, want := parseArgs([]string{"-T2.78", "-t3.14"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseFloatLongOption(t *testing.T) {
	t.Run("both options", func(t *testing.T) {
		resetParser()
		a := Float("little", 0, "little")
		b := Float("big", 0, "big")

		if got, want := parseArgs([]string{"--little", "3.14", "--big", "2.78"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		resetParser()
		a := Float("little", 0, "little")
		b := Float("big", 0, "big")

		if got, want := parseArgs([]string{"--big", "2.78", "--little", "3.14"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestFloatPInvalid(t *testing.T) {
	ensurePanic(t, "cannot use flag with invalid rune", func() {
		_ = FloatP(utf8.RuneError, "", 13.42, "some example flag")
	})
	ensurePanic(t, "cannot use hyphen as a flag", func() {
		_ = FloatP('-', "example", 13.42, "some example flag")
	})
	ensurePanic(t, "cannot use empty flag", func() {
		_ = FloatP('b', "", 13.42, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen", func() {
		_ = FloatP('e', "--example", 13.42, "some example flag")
	})
}

func TestParseFloatPMissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := FloatP('t', "little", 0, "little")
		b := FloatP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		resetParser()
		a := FloatP('t', "little", 0, "little")
		b := FloatP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--little"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseFloatPShortOption(t *testing.T) {
	t.Run("single option with space", func(t *testing.T) {
		resetParser()
		a := FloatP('t', "little", math.NaN(), "little")
		b := FloatP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t", "3.14"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		resetParser()
		a := FloatP('t', "little", 0, "little")
		b := FloatP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t3.14"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 0.0; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		resetParser()
		a := FloatP('t', "little", 0, "little")
		b := FloatP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t3.14", "-T2.78"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		resetParser()
		a := FloatP('t', "little", 0, "little")
		b := FloatP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-T", "2.78", "-t", "3.14"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		resetParser()
		a := FloatP('t', "little", 0, "little")
		b := FloatP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-T2.78", "-t3.14"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseFloatPLongOption(t *testing.T) {
	t.Run("both options", func(t *testing.T) {
		resetParser()
		a := FloatP('t', "little", 0, "little")
		b := FloatP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--little", "3.14", "--big", "2.78"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		resetParser()
		a := FloatP('t', "little", 0, "little")
		b := FloatP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--big", "2.78", "--little", "3.14"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 3.14; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 2.78; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}
