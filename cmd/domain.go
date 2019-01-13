// Copyright © 2018 Arthur Kao <aka0@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/aka0/threatcrowd"

	"github.com/spf13/cobra"
)

// domainCmd represents the domain command
var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Query ThreatCrowd Domain report",
	Long:  `Query ThreatCrowd Domain report`,
	Run: func(cmd *cobra.Command, args []string) {
		client := threatcrowd.NewClient()

		fmt.Println(client.DomainReport(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(domainCmd)
}
