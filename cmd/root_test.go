package cmd

import (
	"path/filepath"
	"testing"

	"github.com/DanDobrick/stagehand/models"
)

func TestFileName(t *testing.T) {
	expectedPath := "/stagehand/workspaces/main.yml"
	result := FileName()
	resultPath := result[len(result)-len(expectedPath):]
	if resultPath != expectedPath {
		t.Errorf("Expected filename to equal %s, got %s", expectedPath, result)
	}
}

//TestParseYaml uses yaml in /testdata
func TestParseYaml(t *testing.T) {
	filepath := filepath.Join("..", "testdata", "test_yaml.yaml")
	result, _ := parseYaml(filepath)
	app1 := models.Application{
		Name:   "ApplicationName",
		File:   "File",
		Bounds: "1 2 3 4 5",
	}

	app2 := models.Application{
		Name:   "noFile",
		Bounds: "1 2 3 4",
	}
	expectedApps := []models.Application{app1, app2}

	for i := range result {
		if result[i] != expectedApps[i] {
			t.Errorf("Expected %v got %v", expectedApps[i], result[i])
		}
	}
}

func TestOpenAndPosition(t *testing.T) {
	// Pending
}
