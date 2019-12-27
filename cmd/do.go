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
	"errors"
	"fmt"
	"github.com/aemengo/hdi/config"
	"github.com/spf13/cobra"
)

var inspect bool

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := performDo(args)
		expectNoError(err)
	},
}

func performDo(args []string) error {
	commands, err := parseCommands()
	if err != nil {
		return err
	}

	if len(args) == 0 {
		return errors.New("missing the action name. Please consult the 'list' command for all action names")
	}

	command, ok := filterCommandsByName(commands, args[0])
	if !ok {
		return fmt.Errorf("an action by the name of '%s' could not be found", args[0])
	}

	if inspect {
		return inspectCommand(command)
	}

	return runCommand(command)
}

func inspectCommand(command config.Command) error {
	for index, step := range command.Steps {
		boldWhite.Printf("Step %d/%d : %s\n", index+1, len(command.Steps), step.Script)
	}

	return nil
}

func runCommand(command config.Command) error {
	fmt.Println("running command...")
	
	boldGreen.Println("Success")
	return nil
}

func filterCommandsByName(commands []config.Command, name string) (config.Command, bool) {
	for _, c := range commands {
		if c.Name == name {
			return c, true
		}
	}

	return config.Command{}, false
}

func init() {
	rootCmd.AddCommand(doCmd)

	doCmd.Flags().BoolVarP(&inspect, "inspect", "i", false, "Display action steps without executing")
}
