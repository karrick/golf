package golf

import (
	"testing"
	"time"
)

func TestDurationInvalid(t *testing.T) {
	ensurePanic(t, "cannot add option without either short, long, or both flags", func() {
		_ = Duration("", "", 0, "some example flag")
	})
	ensurePanic(t, "cannot start short flag with a hyphen: \"-e\"", func() {
		_ = Duration("-e", "example", 0, "some example flag")
	})
	ensurePanic(t, "cannot start long flag with a hyphen: \"--example\"", func() {
		_ = Duration("e", "--example", 0, "some example flag")
	})
}

func TestParseDurationMissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := Duration("t", "little", 0, "little")
		b := Duration("T", "big", 0, "big")

		if got, want := parse("-t"), "option requires argument: \"little\""; got.Error() != want {
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
		a := Duration("t", "little", 0, "little")
		b := Duration("T", "big", 0, "big")

		if got, want := parse("--little"), "option requires argument: \"little\""; got.Error() != want {
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
		a := Duration("t", "little", 0, "little")
		b := Duration("T", "big", 0, "big")

		if got, want := parse("-t 2m"), error(nil); got != want {
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
		a := Duration("t", "little", 0, "little")
		b := Duration("T", "big", 0, "big")

		if got, want := parse("-t2m"), error(nil); got != want {
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
		a := Duration("t", "little", 0, "little")
		b := Duration("T", "big", 0, "big")

		if got, want := parse("-t2m -T3h"), error(nil); got != want {
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
		a := Duration("t", "little", 0, "little")
		b := Duration("T", "big", 0, "big")

		if got, want := parse("-T 3h -t 2m"), error(nil); got != want {
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
		a := Duration("t", "little", 0, "little")
		b := Duration("T", "big", 0, "big")

		if got, want := parse("-T3h -t2m"), error(nil); got != want {
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
		a := Duration("t", "little", 0, "little")
		b := Duration("T", "big", 0, "big")

		if got, want := parse("--little 2m --big 3h"), error(nil); got != want {
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
		a := Duration("t", "little", 0, "little")
		b := Duration("T", "big", 0, "big")

		if got, want := parse("--big 3h --little 2m"), error(nil); got != want {
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
