package nilinterfacecompare

import (
	"github.com/golangci/plugin-module-register/register"
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
	return []*analysis.Analyzer{Analyzer}, nil
}

func (f *NilInterfaceCompare) GetLoadMode() string {
	return register.LoadModeSyntax
}
