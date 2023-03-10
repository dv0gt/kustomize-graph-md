package kustomizationgraph

import (
	"fmt"
	"path/filepath"

	"github.com/dv0gt/kustomize-graph-md/pkg/models"
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

	file, err := g.kustomizationContext.GetFromDirectory(entryPath)

	if err != nil {
		return "", errors.Wrapf(err, "Unable to get kustomization file from given directory %v", entryPath)
	}

	markdown := "```mermaid"
	markdown += "\nflowchart TD"
	markdown += "\nsubgraph " + filepath.Base(entryPath)
	markdown += "\nP0{{kustomization.yaml}}"

	for i, p := range file.PatchesStrategicMerge {
		markdown += "\nP0 --> |patchesStrategicMerge| PATCH" + fmt.Sprint(i) + "(" + p + ")"
	}

	for i, r := range file.Resources {
		file, err = g.kustomizationContext.GetFromDirectory(entryPath + "/" + r)

		if file != nil { // directory with kustomization.yaml
			markdown += "\nsubgraph " + r
			markdown += "\ndirection TB"
			markdown += "\nRESOURCE" + fmt.Sprint(i) + "{{kustomization.yaml}}"

			// TODO: recursive ...

			markdown += "\nend"

			markdown += "\nP0 --> |resources| " + r
		}

	}

	markdown += "\nend"
	markdown += "\n```"

	return markdown, nil
}
