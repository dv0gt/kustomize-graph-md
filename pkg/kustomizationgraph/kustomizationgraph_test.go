package kustomizationgraph

import (
	"fmt"
	"os"
	"testing"

	"github.com/dv0gt/kustomize-graph-md/pkg/kustomizationcontext"
	"github.com/dv0gt/kustomize-graph-md/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestSampleMarkDownGraph(t *testing.T) {

	kustomizationContext := kustomizationcontext.NewContext()
	graph := NewGraph(kustomizationContext)

	workingDir, _ := os.Getwd()
	entryPath := workingDir + "./../../sample/overlays/production"
	markdown, err := graph.BuildGraph(entryPath)

	if err != nil {
		t.Errorf("Graph markdown could not be generated: %v", err.Error())
	}

	expected := "```mermaid" + `
flowchart LR
subgraph production
direction LR
K` + util.Hash(entryPath) + `{{kustomization.yaml}}
subgraph ../../base
direction LR
K` + util.Hash(entryPath+"/../../base") + `{{kustomization.yaml}}
K` + util.Hash(entryPath+"/../../base") + ` --> K` + util.Hash(entryPath+"/../../base") + `R0(deployment.yaml)
end
K` + util.Hash(entryPath) + ` --> |resources| ../../base
end
` + "```"

	fmt.Println(markdown)

	assert.Equal(t, expected, markdown)
}
