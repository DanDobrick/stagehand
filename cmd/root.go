// Copyright © 2018 Daniel Dobrick <dandobrick@gmail.com>
//
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at the root directory of this project
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var rootCmd = &cobra.Command{
	Use:   "stagehand",
	Short: "A command-line tool to open applications and position windows based on recorded preferences",
	Long:  `Will open the applications specified in the yaml file`,
	Run:   openAllApplications,
}

// Application is a struct to hold information for each application to be opened
type Application struct {
	Name string `yaml:"name"`
	File string `yaml:"file"`
	Pos  []int  `yaml:"pos"`
}

// Execute is the main function for Cobra
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func openAllApplications(cmd *cobra.Command, args []string) {
	f, err := ioutil.ReadFile(FileName())
	if err != nil {
		fmt.Println("File read error:", err)
		os.Exit(1)
	}
	as := parseYaml(f)
	for _, app := range as {
		openApp(app)
	}
}

// FileName currently returns the default filename but will eventually determine the filename based on input
func FileName() string {
	u, err := user.Current()
	if err != nil {
		fmt.Println("User error:", err)
		os.Exit(1)
	}
	return u.HomeDir + "/stagehand/workspaces/main.yml"
}

func parseYaml(f []byte) []Application {
	var apps []Application
	err := yaml.Unmarshal(f, &apps)
	if err != nil {
		fmt.Println("Yaml Error:", err)
		os.Exit(1)
	}
	return apps
}

func openApp(app Application) {
	var osCmd *exec.Cmd

	if app.File != "" {
		osCmd = exec.Command("open", "-a", app.Name, app.File)
	} else {
		osCmd = exec.Command("open", "-a", app.Name)
	}
	output, err := osCmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error opening application:", err, string(output))
		os.Exit(1)
	}
}
