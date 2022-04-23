package main

import (
	"github.com/blizzy78/consistent"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(consistent.NewAnalyzer())
}
