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

	ctx := kustomizationcontext.NewContext()
	graphBuilder := kustomizationgraph.NewGraphWithDisplayMode(ctx, displayMode)
	output, err := graphBuilder.BuildGraph(workingDir)

	if err != nil {
		panic(err)
	}

	fmt.Print(output)
}
