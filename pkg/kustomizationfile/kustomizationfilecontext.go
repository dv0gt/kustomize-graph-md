package kustomizationfile

import (
	"path"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

type kustomizationFileContext struct {
	fileSystem afero.Fs
}

// New returns a new context to interact with kustomization files
func New() *kustomizationFileContext {
	defaultFileSystem := afero.NewOsFs()

	return NewFromFileSystem(defaultFileSystem)
}

// NewFromFileSystem creates a context to interact with kustomization files from a provided file system
func NewFromFileSystem(fileSystem afero.Fs) *kustomizationFileContext {
	return &kustomizationFileContext{
		fileSystem: fileSystem,
	}
}

// GetFromDirectory attempts to read a kustomization.yaml file from the given directory
func (k *kustomizationFileContext) GetFromDirectory(directoryPath string) (*KustomizationFile, error) {
	var kustomizationFile KustomizationFile

	fileUtility := &afero.Afero{Fs: k.fileSystem}

	fileFoundCount := 0
	kustomizationFilePath := ""
	for _, kustomizationFile := range KustomizationFileNames {
		currentPath := path.Join(directoryPath, kustomizationFile)

		exists, err := fileUtility.Exists(currentPath)
		if err != nil {
			return nil, errors.Wrapf(err, "Could not check if file %v exists", currentPath)
		}

		if exists {
			kustomizationFilePath = currentPath
			fileFoundCount++
		}
	}

	if kustomizationFilePath == "" {
		return nil, errors.Wrapf(errors.New("Missing kustomization file"), "Error in directory %v", directoryPath)
	}

	if fileFoundCount > 1 {
		return nil, errors.Wrapf(errors.New("Too many kustomization files"), "Error in directory %v", directoryPath)
	}

	kustomizationFileBytes, err := fileUtility.ReadFile(kustomizationFilePath)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not read file %s", kustomizationFilePath)
	}

	err = yaml.Unmarshal(kustomizationFileBytes, &kustomizationFile)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not unmarshal yaml file %s", kustomizationFilePath)
	}

	return &kustomizationFile, nil
}
