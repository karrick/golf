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
	optFlubber := "host5"
	var optHelp bool
	optLimit := -1
	var optQuiet bool
	optServers := "host1,host2"
	optText := "host3,host4"
	var optVerbose bool
	var optVersion bool
	var p golf.Parser

	p.WithBoolVarP(&optHelp, 'h', "help", "Display command line help and exit")
	p.WithBoolVarP(&optQuiet, 'q', "quiet", "Do not print intermediate errors to stderr")
	p.WithIntVarP(&optLimit, 'l', "limit", "Limit output to specified number of lines")
	p.WithBoolVarP(&optVerbose, 'v', "verbose", "Print verbose output to stderr")
	p.WithBoolVarP(&optVersion, 'V', "version", "Print version to stderr and exit")
	p.WithStringVarP(&optServers, 's', "servers", "Some string")
	p.WithStringVar(&optText, "t", "Another string")
	p.WithStringVar(&optFlubber, "flubbers", "Yet another string")

	if err := p.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", filepath.Base(os.Args[0]), err)
		os.Exit(2)
	}

	if optHelp || optVersion {
		fmt.Fprintf(os.Stderr, "%s version %s\n", filepath.Base(os.Args[0]), VersionString)
		if optHelp {
			fmt.Fprintf(os.Stderr, "\texample program to demonstrate library usage\n\n")
			golf.Usage()
		}
		os.Exit(0)
	}

	fmt.Fprintf(os.Stderr, "# os.Args: %v\n", os.Args)
	fmt.Fprintf(os.Stderr, "# golf.Args(): %v\n", golf.Args())
	fmt.Fprintf(os.Stderr, "# golf.NArg(): %v\n", golf.NArg())
	fmt.Fprintf(os.Stderr, "# golf.Arg(0): %v\n", golf.Arg(0))

	fmt.Fprintf(os.Stderr, "# limit: %v\n", optLimit)
	fmt.Fprintf(os.Stderr, "# quiet: %t\n", optQuiet)
	fmt.Fprintf(os.Stderr, "# verbose: %t\n", optVerbose)
}
