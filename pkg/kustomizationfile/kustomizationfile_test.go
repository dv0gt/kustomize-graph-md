package kustomizationfile

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
	fakeFileSystem.Mkdir("app", 0755)

	_, err := NewFromFileSystem(fakeFileSystem).GetFromDirectory("app")

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
	fakeFileSystem.Mkdir("app", 0755)
	emptyFileContents := ""

	afero.WriteFile(fakeFileSystem, "app/kustomization.yaml", []byte(emptyFileContents), 0644)
	afero.WriteFile(fakeFileSystem, "app/kustomization.yml", []byte(emptyFileContents), 0644)
	_, err := NewFromFileSystem(fakeFileSystem).GetFromDirectory("app")

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
	overlayDir := workingDir + "./../../demo/overlays/"
	fmt.Printf("overlay directory: %v\n", overlayDir)

	kustomizationFileContext := New()
	kustomizationFile, err := kustomizationFileContext.GetFromDirectory(overlayDir)

	if err != nil {
		t.Errorf("Could  not load kustomization file: %v\n", err.Error())
	}

	assert.Equal(t, "kustomize.config.k8s.io/v1beta1", kustomizationFile.ApiVersion)
	assert.Equal(t, "Kustomization", kustomizationFile.Kind)
	assert.Equal(t, "./production", kustomizationFile.Resources[0])
	assert.Equal(t, "./staging", kustomizationFile.Resources[1])
	assert.Equal(t, "my-app", kustomizationFile.Images[0].Name)
	assert.Equal(t, "gcr.io/my-platform/my-app", kustomizationFile.Images[0].NewName)
}
