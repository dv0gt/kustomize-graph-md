package kustomizationgraph

import (
	"github.com/dv0gt/kustomize-graph-md/pkg/kustomizationfile"
	"github.com/pkg/errors"
)

type kustomizationFileContext interface {
	GetFromDirectory(directoryPath string) (*kustomizationfile.KustomizationFile, error)
}

type kustomizationGraph struct {
	KustomizationFileContext kustomizationFileContext
}

func New(kustomizeFileContext kustomizationFileContext) *kustomizationGraph {
	return &kustomizationGraph{
		KustomizationFileContext: kustomizeFileContext,
	}
}

// Generate returns a DOT graph based on the dependencies
// from the kustomization.yaml file located in the current working directory
func (g *kustomizationGraph) BuildGraph(entryPath string) (string, error) {

	_, err := g.KustomizationFileContext.GetFromDirectory(entryPath)

	if err != nil {
		return "", errors.Wrapf(err, "Unable to get kustomization file from given directory %v", entryPath)
	}

	return "", nil
}
