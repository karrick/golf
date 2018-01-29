# golf

Go long flag: a light-weight long and short command line option
parser.

## Description

golf is a modest options parsing library for Go command line interface
programs. Meant to be small, like flag included in Go's standard
library, golf does not re-architect how you make command line
programs, nor request you use a DSL for describing your command line
program. It merely allows you to specify which options your program
accepts, and provides the values to your program based on the user's
arguments.

golf does not require a space between the short letter flag equivalent
and its argument. For instance, both of the following are equivalent:

    example -f3.14
    example -f 3.14

golf does allow boolean options to be grouped together when using
their single letter equivalents, such as common in many UNIX
programs. Assuming "-l" and "--limit" point to the same option, all of
the following are equivalent:

    example -vab -l 4
    example -vab -l4
    example -vab --limit 4
    example -v -a -b --limit 4

To prevent ambiguities, however, golf does not allow mixing boolean
short flags and short flags that require an argument. The following
would result in an error message.

    $ example -vl4
    ERROR: cannot parse argument: "-vl4"

## Usage Example

Documentation is available via
[![GoDoc](https://godoc.org/github.com/karrick/golf?status.svg)](https://godoc.org/github.com/karrick/golf).

Use is nearly identical to the standard library flag package. The main
difference is the ability to use both short and long option names. You
may use either short, long, or both command line option flags for each
option you define. To omit either the short or the long flags, simply
use the empty string as its value.

```Go
    optHelp := golf.Bool("h", "help", false, "Display command line help and exit")
    optVerbose := golf.Bool("v", "verbose", false, "Print verbose output to stderr and exit")
    optVersion := golf.Bool("V", "version", false, "Print version output to stderr and exit")
    optLimit := golf.Int("l", "limit", 0, "Limit output to specified number of lines")
    golf.Parse()

    if *optHelp || *optVersion {
        fmt.Fprintf(os.Stderr, "%s version %s\n", filepath.Base(os.Args[0]), versionString)
        if *optHelp {
            fmt.Fprintf(os.Stderr, "\tdemonstration program\n\n")
            golf.Usage()
        }
        os.Exit(0)
    }
```

## TODO

* Support remaining functions from flag package in the standard
  library.
