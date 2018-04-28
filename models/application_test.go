package models

import (
	"testing"
)

var testApp = Application{
	Name:   "App name",
	File:   "File Name",
	Bounds: "Bounds",
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

func TestYamlize(t *testing.T) {
	expectedString := `- name: App name
  file: File Name
  pos: Bounds
`
	result := testApp.Yamlize()

	if result != expectedString {
		t.Errorf("\nExpected: \n %s \ngot: \n %s", expectedString, result)
	}
}
