package ctxRuler_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/k3forx/ctxRuler"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, ctxRuler.Analyzer, "a")
}
