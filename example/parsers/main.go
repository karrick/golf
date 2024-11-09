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
	var err error

	args := os.Args

	if len(args) == 1 {
		fmt.Fprintf(os.Stderr, "USAGE %s foo [-b] [-d DURATION]\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, "USAGE %s bar [-i INT] [-s STRING ]\n", filepath.Base(os.Args[0]))
		os.Exit(2)
	}

	switch args[1] {
	case "foo":
		foo(os.Args[1:])
	case "bar":
		foo(os.Args[1:])
	default:
		fmt.Fprintf(os.Stderr, "USAGE %s [foo|bar]\n", filepath.Base(os.Args[0]))
		os.Exit(2)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", filepath.Base(os.Args[0]), err)
		os.Exit(2)
	}

}

func foo(args []string) {
	var p golf.Parser
	optBool := p.WithBool("b", false, "some bool")
	optDuration := p.WithDurationP('d', "duration", 0, "some duration")
	err := p.Parse(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", filepath.Base(os.Args[0]), err)
		os.Exit(2)
	}
	fmt.Println("optBool:", *optBool)
	fmt.Println("optDuration:", *optDuration)
	for i, arg := range p.Args() {
		fmt.Fprintf(os.Stderr, "# %d: %s\n", i, arg)
	}
}

func bar(args []string) {
	var p golf.Parser
	optInt := p.WithInt("i", 0, "some int")
	optString := p.WithString("s", "", "some string")
	err := p.Parse(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", filepath.Base(os.Args[0]), err)
		os.Exit(2)
	}
	fmt.Println("optInt:", *optInt)
	fmt.Println("optString:", *optString)
	for i, arg := range p.Args() {
		fmt.Fprintf(os.Stderr, "# %d: %s\n", i, arg)
	}
}
