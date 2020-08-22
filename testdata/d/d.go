package main

import (
	"flag"

	dflag "github.com/wingyplus/flagparse/testdata/d/flag"
)

func init() {
	flag := dflag.New()
	flag.Parse()
}

func main() {
	flag.Parse()
}
