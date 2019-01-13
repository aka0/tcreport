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

// emailCmd represents the email command
var emailCmd = &cobra.Command{
	Use:   "email",
	Short: "Query ThreatCrowd Email report",
	Long:  `Query ThreatCrowd Email report`,
	Run: func(cmd *cobra.Command, args []string) {
		client := threatcrowd.NewClient()

		fmt.Println(client.EmailReport(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(emailCmd)
}
