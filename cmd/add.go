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
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [application]",
	Short: "Adds a new application to the list of applications to open",
	Long:  `Desc here`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add command called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
