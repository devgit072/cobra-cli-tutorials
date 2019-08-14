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
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

// sayCmd represents the say command
var sayCmd = &cobra.Command{
	Use:   "math-fun",
	Short: "Short: It will if math is fun or not",
	Long: `Longer description: It will say if math is fun or not. 
    It will say in random order: 50% of time it is fun and 50% time it is not fun.
	It will use rand package to calculate that random value.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(isMathFun())
	},
}

func init() {
	rootCmd.AddCommand(sayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func isMathFun() string{
	rand.Seed(time.Now().UnixNano())
	val := rand.Intn(10000);
	if val % 2 == 0 {
		return "Math is fun"
	} else {
		return "Math is NOT fun"
	}
}
