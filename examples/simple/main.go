package main

import (
	"github.com/karrick/golf"
)

func main() {
	_ = golf.Bool("v", "verbose", false, "bool with both")
	_ = golf.Bool("V", "", false, "bool with short")
	_ = golf.Bool("", "quiet", false, "bool with long")

	_ = golf.String("s", "servers", "host1,host2", "string with both")
	_ = golf.String("t", "", "host3,host4", "string with short")
	_ = golf.String("", "flubbers", "host5", "string with long")

	golf.Parse()
}
