package nilinterfacecompare

// Registration code for golangci-lint integration

import (
	"github.com/golangci/plugin-module-register/register"
	nilanalyzer "github.com/trilopin/nilinterfacecompare/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("nilinterfacecompare", NewLinter)
}

type NilInterfaceCompare struct{}

func NewLinter(_ any) (register.LinterPlugin, error) {
	return &NilInterfaceCompare{}, nil
}

func (f *NilInterfaceCompare) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{nilanalyzer.Analyzer}, nil
}

func (f *NilInterfaceCompare) GetLoadMode() string {
	return register.LoadModeSyntax
}
