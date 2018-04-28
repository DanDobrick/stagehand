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

	"github.com/DanDobrick/stagehand/models"
)

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "returns the current position information for specified application",
	Long:  `eventually this will contain a flag that lets you save this information to the yaml file`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		a := models.Application{Name: args[0]}
		fmt.Println(a.RequestBounds())
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)
	recordCmd.Flags().BoolP("save", "s", false, "Saves application bounds to YAML file (does not overwrite)")
}
