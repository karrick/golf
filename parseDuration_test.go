package golf

import (
	"testing"
	"time"
	"unicode/utf8"
)

func TestDurationInvalid(t *testing.T) {
	ensurePanic(t, "cannot use empty flag string", func() {
		_ = Duration("", 0, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"-e\"", func() {
		_ = Duration("-e", 0, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen: \"--example\"", func() {
		_ = Duration("--example", 0, "some example flag")
	})
}

func TestParseDurationMissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := Duration("t", 0, "little")
		b := Duration("T", 0, "big")

		if got, want := parseArgs([]string{"-t"}), "flag requires argument: 't'"; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		resetParser()
		a := Duration("little", 0, "little")
		b := Duration("big", 0, "big")

		if got, want := parseArgs([]string{"--little"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseDurationShortOption(t *testing.T) {
	t.Run("single option with space", func(t *testing.T) {
		resetParser()
		a := Duration("t", 0, "little")
		b := Duration("T", 0, "big")

		if got, want := parseArgs([]string{"-t", "2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		resetParser()
		a := Duration("t", 0, "little")
		b := Duration("T", 0, "big")

		if got, want := parseArgs([]string{"-t2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		resetParser()
		a := Duration("t", 0, "little")
		b := Duration("T", 0, "big")

		if got, want := parseArgs([]string{"-t2m", "-T3h"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		resetParser()
		a := Duration("t", 0, "little")
		b := Duration("T", 0, "big")

		if got, want := parseArgs([]string{"-T", "3h", "-t", "2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		resetParser()
		a := Duration("t", 0, "little")
		b := Duration("T", 0, "big")

		if got, want := parseArgs([]string{"-T3h", "-t2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseDurationLongOption(t *testing.T) {
	t.Run("both options", func(t *testing.T) {
		resetParser()
		a := Duration("little", 0, "little")
		b := Duration("big", 0, "big")

		if got, want := parseArgs([]string{"--little", "2m", "--big", "3h"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		resetParser()
		a := Duration("little", 0, "little")
		b := Duration("big", 0, "big")

		if got, want := parseArgs([]string{"--big", "3h", "--little", "2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestDurationPInvalid(t *testing.T) {
	ensurePanic(t, "cannot use flag with invalid rune", func() {
		_ = DurationP(utf8.RuneError, "", time.Second, "some example flag")
	})
	ensurePanic(t, "cannot use hyphen as a flag", func() {
		_ = DurationP('-', "example", time.Second, "some example flag")
	})
	ensurePanic(t, "cannot use empty flag", func() {
		_ = DurationP('b', "", time.Second, "some example flag")
	})
	ensurePanic(t, "cannot use flag that starts with a hyphen", func() {
		_ = DurationP('e', "--example", time.Second, "some example flag")
	})
}

func TestParseDurationPMissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := DurationP('t', "little", 0, "little")
		b := DurationP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		resetParser()
		a := DurationP('t', "little", 0, "little")
		b := DurationP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--little"}), "flag requires argument: \"little\""; got.Error() != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseDurationPShortOption(t *testing.T) {
	t.Run("single option with space", func(t *testing.T) {
		resetParser()
		a := DurationP('t', "little", 0, "little")
		b := DurationP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t", "2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("single option without space", func(t *testing.T) {
		resetParser()
		a := DurationP('t', "little", 0, "little")
		b := DurationP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, time.Duration(0); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options without spaces", func(t *testing.T) {
		resetParser()
		a := DurationP('t', "little", 0, "little")
		b := DurationP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-t2m", "-T3h"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with spaces reversed", func(t *testing.T) {
		resetParser()
		a := DurationP('t', "little", 0, "little")
		b := DurationP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-T", "3h", "-t", "2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options with out spaces reversed", func(t *testing.T) {
		resetParser()
		a := DurationP('t', "little", 0, "little")
		b := DurationP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"-T3h", "-t2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestParseDurationPLongOption(t *testing.T) {
	t.Run("both options", func(t *testing.T) {
		resetParser()
		a := DurationP('t', "little", 0, "little")
		b := DurationP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--little", "2m", "--big", "3h"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

	t.Run("both options reversed", func(t *testing.T) {
		resetParser()
		a := DurationP('t', "little", 0, "little")
		b := DurationP('T', "big", 0, "big")

		if got, want := parseArgs([]string{"--big", "3h", "--little", "2m"}), error(nil); got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *a, 2*time.Minute; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

		if got, want := *b, 3*time.Hour; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}
