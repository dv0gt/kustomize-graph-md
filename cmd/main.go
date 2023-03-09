package main

import (
	"fmt"
	"os"

	"github.com/dv0gt/kustomize-graph-md/pkg/kustomizationfile"
	"github.com/dv0gt/kustomize-graph-md/pkg/kustomizationgraph"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	kustomizationfile.New()
	graph, err := kustomizationgraph.New().BuildGraph(workingDir)

	if err != nil {
		panic(err)
	}

	fmt.Print(graph)
}
