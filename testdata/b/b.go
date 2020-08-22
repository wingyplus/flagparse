package main

import (
	"flag"
	"os"
)

func init() {
	flag := flag.NewFlagSet("flag", flag.ExitOnError)
	flag.Parse(os.Args[1:])
}

func main() {
	flag.Parse()
}
