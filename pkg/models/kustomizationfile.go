package models

type KustomizationFile struct {
	Resources  []string `yaml:"resources"`
	Kind       string   `yaml:"kind"`
	Namespace  string   `yaml:"namespace"`
	ApiVersion string   `yaml:"apiVersion"`
}

var KustomizationFileNames = []string{
	"kustomization.yaml",
	"kustomization.yml",
}
