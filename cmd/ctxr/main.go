package main

import (
	"github.com/k3forx/ctxr"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(ctxr.Analyzer) }
