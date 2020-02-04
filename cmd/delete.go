/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"errors"
	"fmt"
	"github.com/aemengo/hdi/config"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from the library",
	Run: func(cmd *cobra.Command, args []string) {
		err := performDelete(args)
		expectNoError(err)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func performDelete(args []string) error {
	if len(args) == 0 {
		return errors.New("missing the action name. Please consult the 'list' command for all action names")
	}

	cfg, err := parseConfig()
	if err != nil {
		return err
	}

	command, ok := filterCommandsByName(cfg.Commands, args[0])
	if !ok {
		return fmt.Errorf("an action by the name of '%s' could not be found", args[0])
	}

	var cmds []config.Command

	for _, cfgCommand := range cfg.Commands {
		if cfgCommand.Name != command.Name {
			cmds = append(cmds, cfgCommand)
		}
	}

	cfg.Commands = cmds

	return writeConfig(cfg)
}