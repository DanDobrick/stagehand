package models

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Application is a struct to hold information for each application to be opened
type Application struct {
	Name   string `yaml:"name"`
	File   string `yaml:"file"`
	Bounds string `yaml:"pos"`
}

// Open opens the application using system calls
func (app Application) Open() ([]byte, error) {
	var osCmd *exec.Cmd

	if app.File != "" {
		osCmd = exec.Command("open", "-a", app.Name, app.File)
	} else {
		osCmd = exec.Command("open", "-a", app.Name)
	}

	return osCmd.CombinedOutput()
}

// Position positions the application using osascript
func (app Application) Position() string {
	// TODO: allow for default size to be changed
	if app.Bounds == "" {
		app.Bounds = "0, 0, 1280, 800"
	}

	command := "tell application \"" + app.Name + "\" to tell window 1 to set bounds to {" + app.Bounds + "}"
	op, err := exec.Command("osascript", "-e", command).CombinedOutput()
	if err != nil {
		fmt.Println("Positioning error:", "Command output:", string(op), "See 'README.md#Positioning Error' for more info")
		os.Exit(1)
	}
	return string(op)
}

// RequestBounds requests current bounds of the application window using osacript
func (app Application) RequestBounds() string {
	command := "tell application \"" + app.Name + "\" to tell window 1 to get bounds"
	op, err := exec.Command("osascript", "-e", command).CombinedOutput()
	if err != nil {
		fmt.Println("Error requesting bounds", err, "Command output:", string(op))
		os.Exit(1)
	}
	return string(op)
}

func (app *Application) RecordBounds() {
	app.Bounds = app.RequestBounds()
}

// Yamlize turns the properties in the application into YAML for writing to workspace file
func (app Application) Yamlize() string {
	var t []string
	t = append(t, "- name: "+app.Name)

	if app.File != "" {
		t = append(t, "  file: "+app.File)
	}

	if app.Bounds != "" {
		t = append(t, "  pos: "+app.Bounds)
	}
	return strings.Join(t, "\n") + "\n"
}
