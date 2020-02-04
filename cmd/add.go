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
	"github.com/aemengo/hdi/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the library",
	Run: func(cmd *cobra.Command, args []string) {
		err := performAdd()
		expectNoError(err)
	},
}

var addTemplate = `---
name: openssl-check-cert
description: "Inspect a certificate using openssl"
args:
- CERT_PATH
steps:
- script: |
    openssl x509 -text -noout -in <CERT_PATH>
`

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func performAdd() error {
	f, err := ioutil.TempFile("", "hdi-")
	if err != nil {
		return err
	}
	defer f.Close()
	defer os.RemoveAll(f.Name())

	_, err = f.Write([]byte(addTemplate))
	if err != nil {
		return err
	}

	cmd := exec.Command(os.Getenv("EDITOR"), f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return err
	}

	var command config.Command
	err = yaml.Unmarshal(data, &command)
	if err != nil {
		return err
	}

	cfg, err := parseConfig()
	if err != nil {
		return err
	}

	cfg.Commands = append(cfg.Commands, command)

	return writeConfig(cfg)
}
