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

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [application]",
	Short: "Adds a new application to the list of applications to open",
	Long:  `Desc here`,
	Run: func(cmd *cobra.Command, args []string) {
		addApplication(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addApplication(n string, f string) {
	d := `- name: ` + n + `
- file: ` + f

	file, err := os.OpenFile(FileName(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 644)
	defer file.Close()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	_, err = file.WriteString(d)
	if err != nil {
		fmt.Println("ERROR", err)
	}
}
