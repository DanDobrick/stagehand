package models

import (
	"testing"
)

var testApp = Application{
	Name:   "App name",
	File:   "File Name",
	Bounds: "1, 2, 3, 4",
}

func TestOpen(t *testing.T) {
	// Figure out testing exec.command
}

func TestPosition(t *testing.T) {
	// Figure out testing exec.command
}

func TestRequestBounds(t *testing.T) {
	// Figure out testing exec.command
}

func TestPos(t *testing.T) {
	// TODO: test app not having bounds set and mock response from RequestBounds()
	expected := testApp.Bounds[:4]
	result := testApp.Pos()

	if result != expected {
		t.Errorf("Expected: %s, got %s", expected, result)
	}
}

func TestSize(t *testing.T) {
	// TODO: test app not having bounds set and mock response from RequestBounds()
	result := testApp.Size()
	expected := testApp.Bounds[6:]

	if result != expected {
		t.Errorf("Expected: \"%s\", got \"%s\"", expected, result)
	}
}

func TestYamlize(t *testing.T) {
	expectedString := `- name: App name
  file: File Name
  pos: 1, 2, 3, 4
`
	result := testApp.Yamlize()

	if result != expectedString {
		t.Errorf("\nExpected: \n %s \ngot: \n %s", expectedString, result)
	}
}
