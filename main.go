package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dv0gt/kustomize-graph-md/pkg/kustomizationcontext"
	"github.com/dv0gt/kustomize-graph-md/pkg/kustomizationgraph"
	"github.com/dv0gt/kustomize-graph-md/pkg/models"
)

var isTopDownFlag = flag.Bool("tb", false, "if set, markdown graph will be oriented top-to-bottom.")

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	flag.Parse()
	displayMode := models.LeftRight
	if *isTopDownFlag {
		displayMode = models.TopBottom
	}

	kustomizationCtx := kustomizationcontext.NewContext()
	kustomizationGraph := kustomizationgraph.NewGraphWithDisplayMode(kustomizationCtx, displayMode)
	graph, err := kustomizationGraph.BuildGraph(workingDir)

	if err != nil {
		panic(err)
	}

	fmt.Print(graph)
}
