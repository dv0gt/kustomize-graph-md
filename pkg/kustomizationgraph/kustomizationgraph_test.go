package kustomizationgraph

import (
	"os"
	"testing"

	"github.com/dv0gt/kustomize-graph-md/pkg/kustomizationfile"
	"github.com/stretchr/testify/assert"
)

func TestSampleMarkDownGraph(t *testing.T) {

	kustomizationFileContext := kustomizationfile.New()
	graph := New(kustomizationFileContext)

	workingDir, _ := os.Getwd()
	entryPath := workingDir + "./../../sample/overlays/"
	markdown, err := graph.BuildGraph(entryPath)

	if err != nil {
		t.Errorf("Graph markdown could not be generated: %v", err.Error())
	}

	expected := "```mermaid" + `
  flowchart LR
    A[./demo/overlays] --> |include| B(./production)
    A --> |include| C(./staging)
    C --> |include| D(../../base)
	` + "```"

	assert.Equal(t, expected, markdown)
}
