package kustomizationgraph

import (
	"fmt"
	"hash/fnv"
	"path/filepath"
	"strconv"

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

	markdown := "```mermaid"
	markdown += "\nflowchart TD"

	subgraph, err := g.addSubGraph(entryPath, filepath.Base(entryPath))
	if err != nil {
		return "", err
	}
	markdown += subgraph
	markdown += "\n```"

	return markdown, nil
}

func (g *KustomizationGraph) addSubGraph(directory string, subGraphName string) (string, error) {
	start := "K" + hash(directory)
	markdown := "\nsubgraph " + subGraphName
	markdown += "\ndirection TB"
	markdown += "\n" + start + "{{kustomization.yaml}}"

	file, err := g.kustomizationContext.GetFromDirectory(directory)
	if err != nil {
		return "", errors.Wrapf(err, "Unable to get kustomization file from given directory %v", directory)
	}

	for i, p := range file.PatchesStrategicMerge {
		markdown += "\n" + start + " --> |patchesStrategicMerge| " + start + "P" + fmt.Sprint(i) + "(" + p + ")"
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
			markdown += "\n" + start + " --> |resources| " + r
			continue
		}

		// regular resource file (no kustomization.yaml)
		markdown += "\n" + start + " --> " + start + "R" + fmt.Sprint(i) + "(" + r + ")"
	}

	markdown += "\nend"

	return markdown, err
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.FormatUint(uint64(h.Sum32()), 10)
}
