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
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"tcreport/threatcrowd"
	"time"

	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var (
	ipCmd = &cobra.Command{
		Use:   "ip",
		Short: "Query ThreatCrowd IP report",
		Long:  `Query ThreatCrowd IP report`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			client := threatcrowd.NewClient()

			rawIP := net.ParseIP(args[0])
			if rawIP.To4() != nil {
				ipReport, err := json.Marshal(client.IPReport(args[0]))

				if err != nil {
					log.Fatalf("unable to parse result %s", args[0])
				}

				fmt.Println(string(ipReport))
			} else {

				f, err := os.Open(args[0])

				if err != nil {
					log.Fatalf("unable to read file %s", args[0])
				}
				defer f.Close()

				scanner := bufio.NewScanner(f)

				var results []threatcrowd.IP

				for scanner.Scan() {
					currIP := scanner.Text()
					results = append(results, client.IPReport(currIP))

					time.Sleep(time.Duration(client.Delay) * time.Second)
				}

				ipReports, err := json.Marshal(results)

				if err != nil {
					log.Fatalln("unable to parse json")
				}

				fmt.Println(string(ipReports))
			}

		},
	}
)

func init() {
	rootCmd.AddCommand(ipCmd)
}
