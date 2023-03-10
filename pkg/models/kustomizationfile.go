package models

type KustomizationFile struct {
	Bases                 []string                     `yaml:"bases"`
	Resources             []string                     `yaml:"resources"`
	Patches               []string                     `yaml:"patches"`
	PatchesStrategicMerge []string                     `yaml:"patchesStrategicMerge"`
	Images                []KustomizationFileImageSpec `yaml:"images"`
	Kind                  string                       `yaml:"kind"`
	Namespace             string                       `yaml:"namespace"`
	ApiVersion            string                       `yaml:"apiVersion"`
}

type KustomizationFileImageSpec struct {
	Name    string `yaml:"name"`
	NewName string `yaml:"newName"`
}

var KustomizationFileNames = []string{
	"kustomization.yaml",
	"kustomization.yml",
}
