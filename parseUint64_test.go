package golf

import (
	"testing"
)

func TestUint64Invalid(t *testing.T) {
	ensurePanic(t, "cannot add option without either short, long, or both flags", func() {
		_ = Uint64("", "", 0, "some example flag")
	})
	ensurePanic(t, "cannot set short flag to a hyphen: \"-e\"", func() {
		_ = Uint64("-e", "example", 0, "some example flag")
	})
	ensurePanic(t, "cannot start long flag with a hyphen: \"--example\"", func() {
		_ = Uint64("e", "--example", 0, "some example flag")
	})
}

func TestParseUint64MissingArgument(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		resetParser()
		a := Uint64("t", "little", 0, "little")
		b := Uint64("T", "big", 0, "big")

		if got, want := parse("-t"), "option requires argument: \"little\""; got.Error() != want {
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
		a := Uint64("t", "little", 0, "little")
		b := Uint64("T", "big", 0, "big")

		if got, want := parse("--little"), "option requires argument: \"little\""; got.Error() != want {
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
		a := Uint64("t", "little", 0, "little")
		b := Uint64("T", "big", 0, "big")

		if got, want := parse("-t 13"), error(nil); got != want {
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
		a := Uint64("t", "little", 0, "little")
		b := Uint64("T", "big", 0, "big")

		if got, want := parse("-t13"), error(nil); got != want {
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
		a := Uint64("t", "little", 0, "little")
		b := Uint64("T", "big", 0, "big")

		if got, want := parse("-t13 -T42"), error(nil); got != want {
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
		a := Uint64("t", "little", 0, "little")
		b := Uint64("T", "big", 0, "big")

		if got, want := parse("-T 42 -t 13"), error(nil); got != want {
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
		a := Uint64("t", "little", 0, "little")
		b := Uint64("T", "big", 0, "big")

		if got, want := parse("-T42 -t13"), error(nil); got != want {
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
		a := Uint64("t", "little", 0, "little")
		b := Uint64("T", "big", 0, "big")

		if got, want := parse("--little 13 --big 42"), error(nil); got != want {
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
		a := Uint64("t", "little", 0, "little")
		b := Uint64("T", "big", 0, "big")

		if got, want := parse("--big 42 --little 13"), error(nil); got != want {
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
