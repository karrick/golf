package golf

import (
	"fmt"
	"io"
	"unicode"
)

func Print(iow io.Writer, width int, prefix, s string) error {
	var inWord bool
	var lineColumns, iWordStart, wordColumns int

	for i, r := range s {
		if unicode.IsSpace(r) != inWord {
			if inWord {
				wordColumns++ // track word rune count
			}
			continue // slurping
		}
		inWord = !inWord // toggle state
		if inWord {
			// was space, now word
			wordColumns = 1 // already read first rune
			iWordStart = i  // word starting index
			continue
		}
		// was word, now space
		if lineColumns == 0 {
			// very first line
			fmt.Fprintf(iow, "%s%s", prefix, s[iWordStart:i])
			lineColumns = len(prefix) + wordColumns
		} else if lineColumns+wordColumns >= width {
			// this word will not fit on current line
			fmt.Fprintf(iow, "\n%s%s", prefix, s[iWordStart:i])
			lineColumns = len(prefix) + wordColumns
		} else {
			// this word should fit (TODO: long word)
			fmt.Fprintf(iow, " %s", s[iWordStart:i])
			lineColumns += 1 + wordColumns
		}
	}
	// if iWordStart < len(s) {
	if lineColumns == 0 {
		fmt.Fprintf(iow, "%s%s\n", prefix, s[iWordStart:])
	} else if lineColumns+wordColumns >= width {
		fmt.Fprintf(iow, "\n%s%s\n", prefix, s[iWordStart:])
	} else {
		fmt.Fprintf(iow, " %s\n", s[iWordStart:])
	}
	// }
	return nil
}
