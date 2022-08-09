package main

import (
	linters "github.com/k3forx/linter-plugin-example"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(linters.TodoAnalyzer)
}
