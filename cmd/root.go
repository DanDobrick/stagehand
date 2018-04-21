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
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stagehand",
	Short: "A command-line tool to open applications and position windows based on recorded preferences",
	Long:  `Creates a boltDb that stores the applications and commands to execute in sequence.`,
	Run:   openAllApplications,
}

// Execute is the main function for Cobra
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func openAllApplications(cmd *cobra.Command, args []string) {
	osCmd := exec.Command("open", "-a", "Terminal", ".")
	a, err := osCmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error:", err, string(a))
	}
}
