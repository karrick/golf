package golf

import (
	"os"
)

func ExampleTextFlow_10_Short() {
	err := Print(os.Stdout, 10, "", "Hello.")
	if err != nil {
		panic(err) // example only
	}
	// Output: Hello.
}

func ExampleTextFlow_10_Last() {
	err := Print(os.Stdout, 7, "", "Hello. A")
	if err != nil {
		panic(err) // example only
	}
	// Output:
	// Hello.
	// A
}

func ExampleTextFlow_10() {
	err := Print(os.Stdout, 10, "", "The  quick  brown  fox  jumped  over  the  lazy  red  dog.")
	if err != nil {
		panic(err) // example only
	}
	// Output:
	// The quick
	// brown fox
	// jumped
	// over the
	// lazy red
	// dog.
}

func ExampleTextFlow_15() {
	err := Print(os.Stdout, 15, "", "The  quick  brown  fox  jumped  over  the  lazy  red  dog.")
	if err != nil {
		panic(err) // example only
	}
	// Output:
	// The quick brown
	// fox jumped over
	// the lazy red
	// dog.
}

func ExampleTextFlowPrefix_20() {
	err := Print(os.Stdout, 20, "|", "The  quick  brown  fox  jumped  over  the  lazy  red  dog.")
	if err != nil {
		panic(err) // example only
	}
	// Output:
	// |The quick brown fox
	// |jumped over the
	// |lazy red dog.
}
