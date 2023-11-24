package kustomizationcontext

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

// TestNoKustomizationFiles tests to validate that when no kustomization files
// are found, an error is returned
func TestNoKustomizationFiles(t *testing.T) {

	// Folder structure for this test
	//
	//   /app

	fakeFileSystem := afero.NewMemMapFs()

	err := fakeFileSystem.Mkdir("app", 0755)
	if err != nil {
		t.Errorf("Couldn't create directory 'app'")
	}

	_, err = NewContextFromFileSystem(fakeFileSystem).GetFromDirectory("app")
	if err == nil {
		t.Errorf("Expected error when reading directory that contains no kustomization files")
	}
}

// TestMultipleKustomizationFiles tests to validate that when multiple kustomization files
// are found, an error is returned
func TestMultipleKustomizationFiles(t *testing.T) {
	// Folder structure for this test
	//
	//   /app
	//   ├── kustomization.yaml
	//   └── kustomization.yml

	fakeFileSystem := afero.NewMemMapFs()
	err := fakeFileSystem.Mkdir("app", 0755)
	if err != nil {
		t.Errorf("Couldn't create directory 'app'")
	}

	emptyFileContents := ""

	err = afero.WriteFile(fakeFileSystem, "app/kustomization.yaml", []byte(emptyFileContents), 0644)
	if err != nil {
		t.Errorf("Couldn't write file")
	}

	err = afero.WriteFile(fakeFileSystem, "app/kustomization.yml", []byte(emptyFileContents), 0644)
	if err != nil {
		t.Errorf("Couldn't write file")
	}

	_, err = NewContextFromFileSystem(fakeFileSystem).GetFromDirectory("app")
	if err == nil {
		t.Errorf("Expected error when reading directory that contains multiple kustomization files")
	}
}

// TestGetFromDirectory tests the GetFromDirectory method to validate that the kustomization
// yaml file was marshaled correctly from the provided path
func TestGetFromDirectory(t *testing.T) {

	workingDir, err := os.Getwd()

	if err != nil {
		t.Error("Working directory could not be retrieved.")
	}
	overlayDir := workingDir + "./../../sample/overlays/"
	fmt.Printf("overlay directory: %v\n", overlayDir)

	kustomizationFileContext := NewContext()
	kustomizationFile, err := kustomizationFileContext.GetFromDirectory(overlayDir)

	if err != nil {
		t.Errorf("Could  not load kustomization file: %v\n", err.Error())
	}

	assert.Equal(t, "kustomize.config.k8s.io/v1beta1", kustomizationFile.ApiVersion)
	assert.Equal(t, "Kustomization", kustomizationFile.Kind)
	assert.Equal(t, "./production", kustomizationFile.Resources[0])
	assert.Equal(t, "./staging", kustomizationFile.Resources[1])
}
