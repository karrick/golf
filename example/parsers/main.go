package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/karrick/golf"
)

func main() {
	// As a reminder, the first argument provided is the name of the
	// program. Therefore if there is only a single command line argument,
	// then the caller did not provide a sub-command.
	if len(os.Args) == 1 {
		bail(errors.New("missing sub-command"))
	}

	// os.Args[0]:  The string used to invoke the program.
	subCommand := os.Args[1]      // The sub-command.
	subCommandArgs := os.Args[2:] // Arguments to the sub-command.

	switch subCommand {
	case "foo":
		foo(subCommandArgs)
	case "bar":
		bar(subCommandArgs)
	default:
		bail(fmt.Errorf("sub-command not recognized: %q", subCommand))
	}
}

func foo(args []string) {
	// Declare and configure a parser to handle the command line options for
	// the 'foo' sub-command.
	var p golf.Parser

	optBool := p.Bool("b", false, "some bool")
	optDuration := p.DurationP('d', "duration", 0, "some duration")

	// After the parser has been configured, use it to parse the command line
	// arguments provided to this function.
	err := p.Parse(args)
	if err != nil {
		bail(err)
	}

	// Do the sub-command operation with the arguments.
	fmt.Println("optBool:", *optBool)
	fmt.Println("optDuration:", *optDuration)

	// The 'Args' method returns a slice of strings that were not consumed by
	// the Parser variables.
	for i, arg := range p.Args() {
		fmt.Fprintf(os.Stderr, "# %d: %s\n", i, arg)
	}
}

func bar(args []string) {
	// Declare and configure a parser to handle the command line options for
	// the 'bar' sub-command.
	var p golf.Parser
	optInt := p.Int("i", 0, "some int")
	optString := p.String("s", "", "some string")

	// After the parser has been configured, use it to parse the command line
	// arguments provided to this function.
	err := p.Parse(args)
	if err != nil {
		bail(err)
	}

	// Do the sub-command operation with the arguments.
	fmt.Println("optInt:", *optInt)
	fmt.Println("optString:", *optString)

	// The 'Args' method returns a slice of strings that were not consumed by
	// the Parser variables.
	for i, arg := range p.Args() {
		fmt.Fprintf(os.Stderr, "# %d: %s\n", i, arg)
	}
}

func bail(err error) {
	basename := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "%s: %s\n", basename, err)
	fmt.Fprintf(os.Stderr, "USAGE %s foo [-b] [-d DURATION]\n", basename)
	fmt.Fprintf(os.Stderr, "USAGE %s bar [-i INT] [-s STRING]\n", basename)
	os.Exit(2)
}
