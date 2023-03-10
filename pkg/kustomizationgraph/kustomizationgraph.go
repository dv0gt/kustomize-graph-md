package kustomizationgraph

import (
	"github.com/dv0gt/kustomize-graph-md/pkg/models"
	"github.com/pkg/errors"
)

type KustomizationContext interface {
	GetFromDirectory(directoryPath string) (*models.KustomizationFile, error)
}

type KustomizationGraph struct {
	KustomizationContext KustomizationContext
}

func NewGraph(kustomizeContext KustomizationContext) *KustomizationGraph {
	return &KustomizationGraph{
		KustomizationContext: kustomizeContext,
	}
}

// Generate returns a DOT graph based on the dependencies
// from the kustomization.yaml file located in the current working directory
func (g *KustomizationGraph) BuildGraph(entryPath string) (string, error) {

	_, err := g.KustomizationContext.GetFromDirectory(entryPath)

	if err != nil {
		return "", errors.Wrapf(err, "Unable to get kustomization file from given directory %v", entryPath)
	}

	return "", nil
}
