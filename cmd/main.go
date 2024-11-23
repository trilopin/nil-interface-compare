package main

import (
	"github.com/trilopin/nilinterfacecompare/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

// The code is usually called from a metalinter like golangci-lint but sometimes is
// useful to have it as an standalone tool
func main() {
	singlechecker.Main(analyzer.Analyzer)
}
