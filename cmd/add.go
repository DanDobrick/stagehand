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
	"io"
	"os"

	"github.com/DanDobrick/stagehand/models"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [application]",
	Short: "Adds a new application to the list of applications to open",
	Long: `Appends application information to ~/stagehand/workspaces/main.yaml.
	
If you want to specify a file for the application to open, you will need to use the "-f" flag.
To record the current window position use the "-p" flag.

EXAMPLE:
stagehand add APP NAME -f Filename`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile(FileName(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 644)
		defer file.Close()

		if err != nil {
			fmt.Println("Error opening file", err)
		}

		addApplication(args, file)
	},
}

var file string
var pos bool

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&file, "file", "f", file, "Adds specified file to application")
	addCmd.Flags().BoolVarP(&pos, "pos", "p", false, "Records application bounds")
}

func addApplication(args []string, writer io.Writer) {
	app := models.Application{
		Name: args[0],
		File: file,
	}
	if pos == true {
		app.RecordBounds()
	}
	_, err := writer.Write([]byte(app.Yamlize()))
	if err != nil {
		fmt.Println("ERROR", err)
	}
}
