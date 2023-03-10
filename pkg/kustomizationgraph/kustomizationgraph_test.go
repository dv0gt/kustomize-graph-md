package kustomizationgraph

import (
	"os"
	"testing"

	"github.com/dv0gt/kustomize-graph-md/pkg/kustomizationcontext"
	"github.com/stretchr/testify/assert"
)

func TestSampleMarkDownGraph(t *testing.T) {

	kustomizationContext := kustomizationcontext.NewContext()
	graph := NewGraph(kustomizationContext)

	workingDir, _ := os.Getwd()
	entryPath := workingDir + "./../../sample/overlays/"
	markdown, err := graph.BuildGraph(entryPath)

	if err != nil {
		t.Errorf("Graph markdown could not be generated: %v", err.Error())
	}

	expected := "```mermaid" + `
flowchart TD

subgraph overlays/production
P0{{kustomization.yaml}}
P0 --> |patchesStrategicMerge| P(patch.yaml)

subgraph base
direction TB
E0{{kustomization.yaml}}
E0 --> |resources| E(deployment.yaml)
end

P0 --> |resources| base
end
	` + "```"

	assert.Equal(t, expected, markdown)
}
