package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/karrick/golf"
)

// VersionString can be overridden during the build with command line parameters.
var VersionString = "1.2.3"

func main() {
	optHelp := golf.Bool("h", "help", false, "Display command line help and exit")
	optLimit := golf.Int("l", "limit", 0, "Limit output to specified number of lines")
	optQuiet := golf.Bool("q", "quiet", false, "Do not print intermediate errors to stderr")
	optVerbose := golf.Bool("v", "verbose", false, "Print verbose output to stderr and exit")
	optVersion := golf.Bool("V", "version", false, "Print version output to stderr and exit")

	_ = golf.String("s", "servers", "host1,host2", "string with both")
	_ = golf.String("t", "", "host3,host4", "string with short")
	_ = golf.String("", "flubbers", "host5", "string with long")

	golf.Parse()

	if *optHelp || *optVersion {
		fmt.Fprintf(os.Stderr, "%s version %s\n", filepath.Base(os.Args[0]), VersionString)
		if *optHelp {
			fmt.Fprintf(os.Stderr, "\texample program\n\n")
			golf.Usage()
		}
		os.Exit(0)
	}

	fmt.Fprintf(os.Stderr, "# os.Args: %v\n", os.Args)
	fmt.Fprintf(os.Stderr, "# golf.Args(): %v\n", golf.Args())
	fmt.Fprintf(os.Stderr, "# golf.NArg(): %v\n", golf.NArg())
	fmt.Fprintf(os.Stderr, "# golf.Arg(0): %v\n", golf.Arg(0))

	fmt.Fprintf(os.Stderr, "# limit: %v\n", *optLimit)
	fmt.Fprintf(os.Stderr, "# quiet: %t\n", *optQuiet)
	fmt.Fprintf(os.Stderr, "# verbose: %t\n", *optVerbose)
}
