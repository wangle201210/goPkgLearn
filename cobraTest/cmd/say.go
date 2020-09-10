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
)

var say string
// sayCmd represents the say command
var sayCmd = &cobra.Command{
	Use:   "say",
	Short: "say something",
	Long: `I can print what you say.`,
	Run: doSay,
}

func init() {
	rootCmd.AddCommand(sayCmd)
	// go run main.go say -s="I can output"
	// go run main.go say -s=icanoutput
	// go run main.go say --str=icanoutput
	sayCmd.Flags().StringVarP(&say,"str","s","hello word","enter wath you want to say")
}

func doSay(cmd *cobra.Command, args []string) {
	fmt.Printf("the string you have input is %s\n",say)
}
