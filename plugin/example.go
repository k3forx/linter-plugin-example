package main

import (
	linters "github.com/k3forx/linter-plugin-example"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		linters.TodoAnalyzer,
	}
}

var AnalyzerPlugin analyzerPlugin
