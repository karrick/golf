package golf

import (
	"testing"
)

func TestWrap(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		if got, want := Wrap(""), "\n"; got != want {
			t.Errorf("\nGOT:\n%q\nWANT:\n%q", got, want)
		}
	})

	t.Run("newline", func(t *testing.T) {
		if got, want := Wrap("\n"), "\n"; got != want {
			t.Errorf("\nGOT:\n%q\nWANT:\n%q", got, want)
		}
	})

	t.Run("final character space", func(t *testing.T) {
		if got, want := Wrap("one two three  "), "one two three\n"; got != want {
			t.Errorf("\nGOT:\n%q\nWANT:\n%q", got, want)
		}
	})

	t.Run("long", func(t *testing.T) {
		got := Wrap("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod velit nec sollicitudin euismod. Lorem ipsum dolor sit amet, consectetur adipiscing elit. In molestie quam ut faucibus lobortis. Mauris sit amet felis dapibus, condimentum metus quis, volutpat nulla. Morbi magna felis, pellentesque vel pellentesque vitae, suscipit quis felis. Donec porta tincidunt nisl id tempus. Cras eros mi, dapibus in laoreet quis, hendrerit et nisl. Quisque dapibus lectus sem, a laoreet turpis accumsan at.")
		want := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod velit\nnec sollicitudin euismod. Lorem ipsum dolor sit amet, consectetur adipiscing\nelit. In molestie quam ut faucibus lobortis. Mauris sit amet felis dapibus,\ncondimentum metus quis, volutpat nulla. Morbi magna felis, pellentesque vel\npellentesque vitae, suscipit quis felis. Donec porta tincidunt nisl id tempus.\nCras eros mi, dapibus in laoreet quis, hendrerit et nisl. Quisque dapibus\nlectus sem, a laoreet turpis accumsan at.\n"
		if got != want {
			t.Errorf("\nGOT:\n%v\nWANT:\n%v", got, want)
		}
	})

	t.Run("final word on new line", func(t *testing.T) {
		got := Wrap("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod velit nec")
		want := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod velit\nnec\n"
		if got != want {
			t.Errorf("\nGOT:\n%v\nWANT:\n%v", got, want)
		}
	})

	t.Run("prefix", func(t *testing.T) {
		got := LineWrapper{Prefix: "|", Max: 40}.Wrap("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod velit nec")
		want := "|Lorem ipsum dolor sit amet, consectetur\n|adipiscing elit. Donec euismod velit\n|nec\n"
		if got != want {
			t.Errorf("\nGOT:\n%v\nWANT:\n%v", got, want)
		}
	})

	t.Run("non-breaking space", func(t *testing.T) {
		got := Wrap("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod velit\u00A0nec")
		want := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod\nvelit\u00A0nec\n"
		if got != want {
			t.Errorf("\nGOT:\n%v\nWANT:\n%v", got, want)
		}
	})
}
