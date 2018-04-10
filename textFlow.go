package golf

import (
	"fmt"
	"io"
	"unicode"
)

func Print(iow io.Writer, width int, prefix, s string) error {
	var inWord bool
	var lineColumns, iWordStart, iWordEnd, wordColumns int

	for i, r := range s {
		if unicode.IsSpace(r) != inWord {
			if inWord {
				wordColumns++ // track word rune count
			}
			continue // slurping up
		}
		inWord = !inWord // toggle state
		if inWord {
			// was space, now word
			wordColumns = 1 // already read first rune
			iWordStart = i	// word starting index
		} else {
			// was word, now space
			iWordEnd = i // word ending index
			// fmt.Fprintf(os.Stderr, "[%q: L %d; W %d; T %d]\n", s[wordStarts:wordEnds], column, wordColumns, newColumn)

			if lineColumns == 0 {
				// very first line
				fmt.Fprintf(iow, "%s%s", prefix, s[iWordStart:iWordEnd])
				lineColumns = len(prefix) + wordColumns
			} else if lineColumns + wordColumns >= width {
				// this word will not fit on current line
				fmt.Fprintf(iow, "\n%s%s", prefix, s[iWordStart:iWordEnd])
				lineColumns = len(prefix) + wordColumns
			} else {
				// this word should fit (TODO: long word)
				fmt.Fprintf(iow, " %s", s[iWordStart:iWordEnd])
				lineColumns += 1 + wordColumns
			}
		}
	}
	if lineColumns == 0 {
		fmt.Fprintf(iow, "%s%s\n", prefix, s[iWordStart:])
	} else if lineColumns+wordColumns >= width {
		fmt.Fprintf(iow, "\n%s%s\n", prefix, s[iWordStart:])
	} else {
		fmt.Fprintf(iow, " %s\n", s[iWordStart:])
	}
	return nil
}
