package main

import (
	"github.com/wingyplus/flagparse"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(flagparse.Analyzer)
}
