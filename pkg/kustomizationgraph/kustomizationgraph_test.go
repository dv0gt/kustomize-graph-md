package kustomizationgraph

import (
	"os"
	"testing"

	"github.com/dv0gt/kustomize-graph-md/pkg/kustomizationcontext"
	"github.com/dv0gt/kustomize-graph-md/pkg/models"
	"github.com/dv0gt/kustomize-graph-md/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestSampleMarkDownGraph_LR(t *testing.T) {
	assertSampleMarkdownGraph(t, models.LeftRight)
}

func TestSampleMarkDownGraph_TB(t *testing.T) {
	assertSampleMarkdownGraph(t, models.TopBottom)
}

func assertSampleMarkdownGraph(t *testing.T, mode models.DisplayMode) {
	kustomizationContext := kustomizationcontext.NewContext()
	graph := NewGraphWithDisplayMode(kustomizationContext, mode)

	workingDir, _ := os.Getwd()
	entryPath := workingDir + "./../../sample/overlays/production"
	markdown, err := graph.BuildGraph(entryPath)

	if err != nil {
		t.Errorf("Graph markdown could not be generated: %v", err.Error())
	}

	expected := "```mermaid" + `
flowchart ` + mode.ToString() + `
subgraph ./production
direction ` + mode.ToString() + `
K` + util.Hash(entryPath) + `{{kustomization.yaml}}
subgraph ../../base
direction ` + mode.ToString() + `
K` + util.Hash(entryPath+"/../../base") + `{{kustomization.yaml}}
K` + util.Hash(entryPath+"/../../base") + ` --> K` + util.Hash(entryPath+"/../../base") + `R0(deployment.yaml)
K` + util.Hash(entryPath+"/../../base") + ` --> K` + util.Hash(entryPath+"/../../base") + `R1(namespace.yaml)
end
K` + util.Hash(entryPath) + ` --> |resources| ../../base
end
` + "```"

	// fmt.Println(markdown)

	assert.Equal(t, expected, markdown)
}
