package nilinterfacecompare

import (
	"fmt"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	fmt.Println("Registering nilinterfacecompare")
	register.Plugin("nilinterfacecompare", NewLinter)
	fmt.Println("Registered nilinterfacecompare")
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
