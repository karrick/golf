# golf

Go long flag: a light-weight long and short command line option
parser.

## Description

golf is modest options parsing library for Go command line interface
programs. Meant to be small, like flag, included in Go's standard
library, does not re-architect how you make command line programs, nor
request you use a DSL for describing your command line program. It
merely allows you to specify which options your program accepts, and
provides the values to your program based on the user's arguments.

golf is 

## Usage Example

Use is nearly identical to the standard library flag package. The main
difference is the ability to use both short and long option names.

```Go
    optHelp := golf.Bool("h", "help", false, "Display command line help and exit")
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

## TODO

* Support remaining functions from flag package in the standard
  library.
