package main

import (
	"github.com/k3forx/ctxRuler"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(ctxRuler.Analyzer) }
