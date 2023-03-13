package kustomizationgraph

import (
	"fmt"
	"path/filepath"

	"github.com/dv0gt/kustomize-graph-md/pkg/models"
	"github.com/dv0gt/kustomize-graph-md/pkg/util"
	"github.com/pkg/errors"
)

type KustomizationContext interface {
	GetFromDirectory(directoryPath string) (*models.KustomizationFile, error)
}

type KustomizationGraph struct {
	kustomizationContext KustomizationContext
}

func NewGraph(kustomizeContext KustomizationContext) *KustomizationGraph {
	return &KustomizationGraph{
		kustomizationContext: kustomizeContext,
	}
}

func (g *KustomizationGraph) BuildGraph(entryPath string) (string, error) {

	markdown := "```mermaid"
	markdown += addLine("flowchart LR")

	subgraph, err := g.addSubGraph(entryPath, filepath.Base(entryPath))
	if err != nil {
		return "", err
	}
	markdown += subgraph
	markdown += addLine("```")

	return markdown, nil
}

func (g *KustomizationGraph) addSubGraph(directory string, subGraphName string) (string, error) {
	start := "K" + util.Hash(directory)
	markdown := addLine("subgraph " + subGraphName)
	markdown += addLine("direction LR")
	markdown += addLine(start + "{{kustomization.yaml}}")

	file, err := g.kustomizationContext.GetFromDirectory(directory)
	if err != nil {
		return "", errors.Wrapf(err, "Unable to get kustomization file from given directory %v", directory)
	}

	for i, r := range file.Resources {
		resourcePath := directory + "/" + r
		file, _ = g.kustomizationContext.GetFromDirectory(resourcePath)

		if file != nil { // directory with kustomization.yaml
			subgraph, err := g.addSubGraph(resourcePath, r)
			if err != nil {
				return "", err
			}
			markdown += subgraph
			markdown += addLine(start + " --> |resources| " + r)
			continue
		}

		// regular resource file (no kustomization.yaml)
		markdown += addLine(start + " --> " + start + "R" + fmt.Sprint(i) + "(" + r + ")")
	}

	markdown += addLine("end")

	return markdown, err
}

func addLine(line string) string {
	return "\n" + line
}
