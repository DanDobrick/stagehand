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
	"os/user"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/DanDobrick/stagehand/models"
)

var rootCmd = &cobra.Command{
	Use:   "stagehand",
	Short: "A command-line tool to open applications and position windows based on recorded preferences",
	Long:  `Will open and position the applications specified in the yaml file`,
	Run: func(cmd *cobra.Command, args []string) {
		apps, err := parseYaml(FileName())

		if err != nil {
			fmt.Println("Yaml parsing error", err)
			os.Exit(1)
		}

		var wg sync.WaitGroup
		for _, app := range apps {
			wg.Add(1)
			go openAndPosition(app, &wg)
		}
		wg.Wait()
	},
}

// Execute is the main function for Cobra
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
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

func parseYaml(filePath string) ([]models.Application, error) {
	var apps []models.Application

	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return apps, err
	}

	err = yaml.Unmarshal(f, &apps)
	return apps, err
}

func openAndPosition(app models.Application, wg *sync.WaitGroup) {
	output, err := app.Open()
	if err != nil {
		fmt.Println("Error opening application:", app.Name, err, string(output))
		os.Exit(1)
	}
	count := 0
	for count < 7 {
		app.Position()

		currPos := app.RequestBounds()
		if app.Bounds == currPos {
			wg.Done()
			break
		}
		time.Sleep(1 * time.Second)
		count++
	}
	fmt.Println("Positioning Error:", app.Name, "could not position Window. See 'README.md#Positioning Error' for more info")
	os.Exit(1)
}
