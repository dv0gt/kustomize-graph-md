package kustomizationgraph

import (
	"fmt"
	"os"
	"testing"

	"github.com/dv0gt/kustomize-graph-md/pkg/kustomizationcontext"
	"github.com/dv0gt/kustomize-graph-md/pkg/models"
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
K2648915859 --> K3198009923
K3198009923 --> K117322154
K117322154[[./moduleA<br/><br/>deploymentModuleA.yaml]]
K3198009923[[../../base<br/><br/>deployment.yaml<br/>namespace.yaml]]
K2648915859[[./production]]
` + "```"

	fmt.Println(markdown)

	assert.Equal(t, expected, markdown)
}
