/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/aemengo/hdi/config"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
	"strings"
)

var filter string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks that 'hdi' has stored",
	Run: func(cmd *cobra.Command, args []string) {
		err := performList()
		expectNoError(err)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.
	listCmd.Flags().StringVarP(&filter, "filter", "f", "", "Filter commands by a query")
}

func performList() error {
	commands, err := parseCommands()
	if err != nil {
		return err
	}

	commands = filterCommandsByFilter(commands, filter)

	if len(commands) == 0 {
		fmt.Println("No results found")
		return nil
	}

	table := []string{
		"Name | Description | Snippet",
	}

	for _, command := range commands {
		if len(command.Steps) == 0 {
			table = append(table, fmt.Sprintf("%s | %s | -", command.Name, command.Description))
		} else {
			table = append(table, fmt.Sprintf("%s | %s | %s", command.Name, command.Description, truncateString(command.Steps[0].Script, 28)))
		}
	}

	result := strings.Split(columnize.SimpleFormat(table), "\n")
	boldWhite.Println(result[0])
	fmt.Println(strings.Join(result[1:], "\n"))
	return nil
}

func filterCommandsByFilter(commands []config.Command, query string) []config.Command {
	q := strings.TrimSpace(query)

	if q == "" {
		return commands
	}

	var cmds []config.Command
	for _, c := range commands {
		if strings.Contains(c.Name, q) {
			cmds = append(cmds, c)
			continue
		}

		if strings.Contains(c.Description, q) {
			cmds = append(cmds, c)
			continue
		}
	}

	return cmds
}

func truncateString(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num] + "..."
	}
	return bnoden
}