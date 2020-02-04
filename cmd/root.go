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
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	cfgFile    string
	boldWhite  = color.New(color.FgWhite, color.Bold)
	boldGreen  = color.New(color.FgGreen, color.Bold)
	boldYellow = color.New(color.FgYellow, color.Bold)
	boldRed    = color.New(color.FgRed, color.Bold)
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hdi",
	Short: "A tool for managing and invoking common command-line tasks",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hdi.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		return
	}

	home, err := homedir.Dir()
	expectNoError(err)
	cfgFile = filepath.Join(home, ".hdi.yaml")

	if notExist(cfgFile) {
		err = ioutil.WriteFile(cfgFile, []byte(`---`), 0600)
		expectNoError(err)
	}
}

func writeConfig(cfg config.Config) error {
	cfgData, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(cfgFile, cfgData, 0600)
}

func parseConfig() (config.Config, error) {
	data, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func parseCommands() ([]config.Command, error) {
	cfg, err := parseConfig()
	if err != nil {
		return nil, err
	}

	return cfg.Commands, nil
}

func notExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true
	}

	return false
}

func expectNoError(err error) {
	if err == nil {
		return
	}

	if err.Error() == "" {
		fmt.Println(boldRed.Sprint("Error"))
	} else {
		fmt.Printf(boldRed.Sprint("Error")+"\n%s.\n", err)
	}

	os.Exit(1)
}
