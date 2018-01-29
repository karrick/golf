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

    example -t3.14
    example -t 3.14

golf does allow boolean options to be grouped together when using
their single letter equivalents, such as common in many UNIX
programs. Assuming "-l" and "--limit" point to the same option, all of
the following are equivalent:

    example -va -t 4 -shost.example.com
    example -va -t4 -s host.example.com
    example -va --threshold 4 --server host.example.com
    example -v -a --threshold 4 --server host.example.com

To prevent ambiguities, however, golf does not allow mixing boolean
short flags and short flags that require an argument in the same
parameters. The following would result in error messages.

    $ example -t4v
    ERROR: cannot parse argument: "-t4v"

    $ example -vt4
    ERROR: cannot parse argument: "-vt4"

## Usage Example

Documentation is available via
[![GoDoc](https://godoc.org/github.com/karrick/golf?status.svg)](https://godoc.org/github.com/karrick/golf).

Use is nearly identical to the standard library flag package. The main
difference is the ability to use both short and long option names. You
may use either short, long, or both command line option flags for each
option you define. To omit either the short or the long flags, simply
use the empty string as its value.


```Go
optAbsolute := golf.Bool("a", "", true, "Display absolute values")
optHelp := golf.Bool("h", "help", false, "Display command line help and exit")
optLimit := golf.Int("l", "limit", 0, "Limit output to specified number of lines")
optServer := golf.String("s", "server", "", "Send query to specified server")
optTheshold := golf.Float("t", "threshold", 0, "Set minimum threshold")
optVerbose := golf.Bool("v", "verbose", false, "Print verbose output to stderr and exit")
optVersion := golf.Bool("V", "version", false, "Print version output to stderr and exit")

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

In an attempt to be largely compatible with the flag library,
specifying an option flag has no error return value, so attempting to
create a flag with illegal arguments will panic. While this behavior
is not necessarily acceptable for libraries, if command line options
are not correctly defined by the program the case will be caught early
by running the program.

## Help Example

Invoking `golf.Usage()` will display the program name, followed by a
list of command line option flags. The short and long flag names are
displayed if they are both defined, otherwise, just the short or the
long is displayed. After the flag name will be a token representing
the expected value data type, so the user knows what type of parsing
will be invoked on any value provided for that flag.

On a separate line after each flag is printed a tab character,
followed by the description and the default value for that flag.

```
Usage of example:
  -h, --help
	Display command line help and exit (default: false)
  -l, --limit int
	Limit output to specified number of lines (default: 0)
  -q, --quiet
	Do not print intermediate errors to stderr (default: false)
  -v, --verbose
	Print verbose output to stderr and exit (default: false)
  -V, --version
	Print version output to stderr and exit (default: false)
  -s, --servers string
	string with both (default: host1,host2)
  -t string
	string with short (default: host3,host4)
  --flubbers string
	string with long (default: host5)
```

## TODO

* Support remaining functions from flag package in the standard
  library.
