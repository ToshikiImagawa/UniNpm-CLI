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
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

type unityConfig struct {
	authUrl    string
	authToken  string
	userEmail  string
	alwaysAuth string
}

var userEmail string

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		callNpmConfigSet(args)
		UpdateUpmConfig(cmd)
	},
}

func init() {
	configCmd.AddCommand(setCmd)
}

func callNpmConfigSet(args []string) {
	var cm *exec.Cmd

	if args == nil || len(args) == 0 {
		cm = exec.Command("npm", "config", "set")
	} else if len(args) == 1 {
		cm = exec.Command("npm", "config", "set", args[0])
	} else if len(args) == 2 {
		cm = exec.Command("npm", "config", "set", args[0], args[1])
	} else if len(args) == 3 {
		cm = exec.Command("npm", "config", "set", args[0], args[1], args[2])
	} else {
		fmt.Println("Not Support.")
		os.Exit(1)
		return
	}
	_ = cm.Run()
}
