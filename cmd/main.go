package main

import (
	"fmt"
	"os"

	"github.com/dv0gt/kustomize-graph-md/pkg/kustomizationcontext"
	"github.com/dv0gt/kustomize-graph-md/pkg/kustomizationgraph"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	kustomizationCtx := kustomizationcontext.NewContext()

	graph, err := kustomizationgraph.NewGraph(kustomizationCtx).BuildGraph(workingDir)

	if err != nil {
		panic(err)
	}

	fmt.Print(graph)
}
