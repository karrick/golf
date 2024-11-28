# golf

Go long flag: a light-weight long and short command line option parser.

## Description

golf is a modest options parsing library for Go command line interface
programs. Meant to be small, like flag included in Go's standard library, golf
does not re-architect how you make command line programs, nor request you use
a DSL for describing your command line program. It merely allows you to
specify which options your program accepts, and provides the values to your
program based on the user's arguments.

* Fully POSIX-compliant flags, including short & long flags.
* Helpful functions for printing help and usage information.
* Optional space between short flag and its argument.
* No new concepts to learn beyond typical use of flag library.
* Supports GNU extension allowing flags to appear intermixed with
  other command line arguments.

## Usage

Documentation is available via
[![GoDoc](https://godoc.org/github.com/karrick/golf?status.svg)](https://godoc.org/github.com/karrick/golf).

### Drop-In Example

When the flag package prefix is changed to golf, your program will require a
single-hyphen when the flag name is one rune long, and a double-hyphen prefix
when the flag is more than one rune long.

    $ example -v --limit 3

Drop-in compatibiliy usage is nearly identical to the standard library flag
package, with a few additions for supporting flags with both short and long
style options, albeit with a few features from the standard library not
supported. However, it provides all the functions that most command line
programs would require.

golf is designed such that you can change every flag package prefix to golf,
recompile your program, and be able to use more POSIX friendly command line
options. Both Type and TypeVar style flag declarations are supported, as shown
below:

```Go
package main

import (
    "fmt"

    "github.com/karrick/golf"
)

func main() {
    // 'Type' functions are exactly like their corresponding functions from
    // the 'flag' standard library. However the length of the argument string
    // determines whether a single or double-hyphen prefix is used to change
    // its value. When this argument is a single rune long string then a
    // single-hyphen prefix is used to change its value, for example, "V" will
    // configure the parser to associate '-V' with the option. Otherwise a
    // double-hyphen prefix is used to change its value, for example,
    // "version" will configure the parser to associate '--version' with the
    // option.
    optType1 := golf.Bool("b", false, "can take a one rune flag name and return a variable pointer")
    otpType2 := golf.String("id", "", "can also take a multiple rune flag name and return a variable pointer")

    // 'TypeP' functions take two arguments. The first argument is a rune to
    // be associated with a single-hyphen prefix, for example, '-V'. The
    // second argument is a string to be associated with a double-hyphen
    // prefix, for example, '--version'.
    optTypeP := golf.DurationP('d', "duration", 0, "takes a one rune and a multiple rune flag name and returns a pointer to a variable")

    // The difference between 'Type' and 'TypeVar' is 'Type' returns a pointer
    // to a variable, whereas 'TypeVar' accepts a pointer to a variable.
    var optTypeVar float64
    golf.FloatVar(&optTypeVar, "f", 6.02e-23, "optTypeVar takes a pointer to a variable and a single flag")

    // For completeness, a 'TypeVarP' set of functions is also provided.
    var optTypeVarP int64
    golf.Int64VarP(&optTypeVarP, 'i', "int64", 13, "optTypeVarP takes a pointer to a variable, a rune, and a string")

    // After the options have been declared, the `golf.Parse()` function will
    // parse the command line options from `os.Args`.
    golf.Parse()

    // Notice that the returned 'Type' and 'TypeP' functions return a pointer
    // that must be referenced in order to obtain the value.
    fmt.Println("optType: ", *optType1)
    fmt.Println("optTypeP: ", *optTypeP)

    // In contrast, the 'TypeVar' and 'TypeVarP' functions accept a pointer to
    // the variable for the library to store its value, and accessing the
    // value does not require a pointer dereference.
    fmt.Println("optTypeVar: ", optTypeVar)
    fmt.Println("optTypeVarP: ", optTypeVarP)
}
```

### Parser

golf also provides support for one or more non-global `golf.Parser` instances,
so different `golf.Parser` instances can be created to parse different
combinations of options, for instance, when different sub-commands support
differing sets of options.

```Go
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
    optBool := p.WithBool("b", false, "some bool")
    optDuration := p.WithDurationP('d', "duration", 0, "some duration")

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
    optInt := p.WithInt("i", 0, "some int")
    optString := p.WithString("s", "", "some string")

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
    fmt.Fprintf(os.Stderr, "USAGE %s bar [-i INT] [-s STRING ]\n", basename)
    os.Exit(2)
}
```

## Features

golf allows specifying both a short and a long flag name by calling
the same function stem but with a P suffix, as shown below. When
declaring a short and long flag pair, the short flag type is a rune,
to prevent accidentally initializing a command line flag pair with two
multi-rune strings, which is not allowed. (The reason for this
compromise is because the command line flag declaration pattern used
by `flag` provides no means for signalling an error.)

```Go
optRaw := golf.BoolP('r', "raw", false, "Display raw output")
optServer := golf.StringP('s', "server", "", "Send query to specified server")
optTheshold := golf.FloatP('t', "threshold", 0, "Set minimum threshold")
optVerbose := golf.BoolP('v', "verbose", false, "Display verbose output")
golf.Parse()
```

All of the below examples result in the same flag values being set:

    $ example -sfoo.example.com -t3.14
    $ example -s foo.example.com -t 3.14
    $ example --server foo.example.com --threshold 3.14

golf allows boolean options to be grouped together when using their
single letter equivalents, such as common in many UNIX programs. all
of the following are equivalent:

    $ example -rv -t 4 -shost.example.com
    $ example -rv -t4 -s host.example.com
    $ example -rv --threshold 4 --server host.example.com
    $ example -v -r --threshold 4 --server host.example.com

golf also allow concatenation of one or more boolean short flags with
at most one short flag that requires an argument provided the flag
that requires the argument is the last flag in the parameter. Both of
the following are legal, although the second example happens to parse
oddly in my brain, because v and t appear grouped closer than the t
and the 4. Nevertheless, both are equivalent and unambiguous:

    $ example -vt4
    $ example -vt 4

To prevent ambiguities, however, golf does not allow placing any
additional flags after a flag that requires an argument, even if it
may appear to be legal, in the same argument. For instance, if the i
takes an integer and the s flag takes a string, this will still result
in a parsing error:

    $ example -i4sfoo.example.com
    ERROR: strconv.ParseInt: parsing "4sfoo.example.com": invalid syntax

This however is legal:

    $ example -i4 -sfoo.example.com

In an attempt to be largely compatible with the flag library,
specifying an option flag has no error return value, so attempting to
create a flag with illegal arguments will panic. While causing a panic
is poor practice for a library, if command line options are not
correctly defined by the program the case will be caught early by
running the program.

This library supports a common practice in GNU command line argument
parsing that allows command line flags and command line arguments to
be intermixed. For instance, the following invocations would be
equivalent.

    $ example -t 3.14 arg1 arg2
    $ example arg1 -t3.14 arg2
    $ example arg1 arg2 -t 3.14

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
example version 1.2.3
    example program

Usage of example:
  -h, --help
    Display command line help and exit (default: false)
  -l, --limit int
    Limit output to specified number of lines (default: 0)
  -q, --quiet
    Do not print intermediate errors to stderr (default: false)
  -v, --verbose
    Print verbose output to stderr (default: false)
  -V, --version
    Print version to stderr and exit (default: false)
  -s, --servers string
    Some string (default: host1,host2)
  -t string
    Another string (default: host3,host4)
  --flubbers string
    Yet another string (default: host5)
```

## TODO

* Support remaining functions from flag package in the standard
  library.
