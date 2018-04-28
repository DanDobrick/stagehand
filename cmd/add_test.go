package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestAddApplication(t *testing.T) {
	args := []string{"Test Name"}
	appName := "- name: " + args[0]

	//Test without flags
	nameOnly := appName + "\n"

	buffer := bytes.Buffer{}
	addApplication(args, &buffer)
	result := buffer.String()

	if result != nameOnly {
		t.Errorf("Expected '%s', got '%s'", nameOnly, result)
	}

	//Test with flags
	buffer.Reset()
	// Sets filename via the "file" flag
	addCmd.Flags().Set("file", "File Name")
	// Sets bounds via the "pos" flag
	addCmd.Flags().Set("pos", "1 2 3 4")

	ss := []string{appName, "  file: File Name", "  pos: 1 2 3 4"}
	expected := strings.Join(ss, "\n") + "\n"

	addApplication(args, &buffer)
	result = buffer.String()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
