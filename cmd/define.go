/*
Copyright © 2021 iminfinity <this.is.me.infinity@gmai.com>

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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// defineCmd represents the define command
var defineCmd = &cobra.Command{
	Use:   "define",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getMeaning(args)
	},
}

func init() {
	rootCmd.AddCommand(defineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// defineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// defineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Dictionary struct
type Dictionary struct {
	Word       string `json:"word,omitempty"`
	Definition string `json:"definition,omitempty"`
}

var dictionary []Dictionary

func getMeaning(args []string) {
	jsonFile, err := os.Open("dictionary.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonData, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(jsonData, &dictionary)

	found := false
	for _, dict := range dictionary {
		if dict.Word == args[0] {
			fmt.Printf("\n\t\t %s\n", dict.Definition)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("\n\t Word not in dictionary \n")
	}
}
