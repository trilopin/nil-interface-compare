package main

import (
	"github.com/trilopin/nilinterfacecompare"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(nilinterfacecompare.Analyzer)
}
