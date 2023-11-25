package kustomizationgraph

import (
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
	displayMode          models.DisplayMode
}

func NewGraph(kustomizeContext KustomizationContext) *KustomizationGraph {
	return NewGraphWithDisplayMode(kustomizeContext, models.LeftRight)
}

func NewGraphWithDisplayMode(kustomizeContext KustomizationContext, displayMode models.DisplayMode) *KustomizationGraph {
	return &KustomizationGraph{
		kustomizationContext: kustomizeContext,
		displayMode:          displayMode,
	}
}

func (g *KustomizationGraph) BuildGraph(entryPath string) (string, error) {

	markdown := "```mermaid"
	markdown += addLine("flowchart " + g.displayMode.ToString())

	_, subgraph, err := g.addSubGraph(entryPath, "./"+filepath.Base(entryPath))
	if err != nil {
		return "", err
	}
	markdown += subgraph
	markdown += addLine("```")

	return markdown, nil
}

func (g *KustomizationGraph) addSubGraph(directory string, subGraphName string) (string, string, error) {
	// first define starting point of the sub graph
	start := "K" + util.Hash(directory)

	file, err := g.kustomizationContext.GetFromDirectory(directory)
	if err != nil {
		return "", "", errors.Wrapf(err, "Unable to get kustomization file from given directory %v", directory)
	}

	resourceFiles := ""
	markdown := ""
	for _, r := range file.Resources {
		resourcePath := directory + "/" + r
		file, _ = g.kustomizationContext.GetFromDirectory(resourcePath)

		// build another sub graph if resource is a directory
		if file != nil {
			subgraphStart, subgraphMarkdown, err := g.addSubGraph(resourcePath, r)
			if err != nil {
				return "", "", err
			}
			markdown += addLine(start + " --> " + subgraphStart)
			markdown += subgraphMarkdown
			continue
		}

		resourceFiles += "<br/>" + r
	}

	if resourceFiles != "" {
		markdown += addLine(start + "[[" + subGraphName + "<br/>#171;#171;#171; #187;#187;#187;" + resourceFiles + "]]")
	} else {
		markdown += addLine(start + "[[" + subGraphName + "]]")
	}

	return start, markdown, err
}

func addLine(line string) string {
	return "\n" + line
}
