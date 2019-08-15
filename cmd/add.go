package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add two numbers and display the results",
	Long: `Add two numbers and display the results.
It will take two parameters only numerical value else it will throw error`,
	Run: func(cmd *cobra.Command, args []string) {
		addTwoNumbersAndDisplay(cmd)
	},
	// You can add PreRunE check on flags value or type etc
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkIfValueAreNotLarge(cmd)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().IntP("ValueA", "A", 0, "-A <some value>")
	addCmd.PersistentFlags().IntP("ValueB", "B", 0, "-B <some value>")

	// Tell cli that these are mandatory params.
	addCmd.MarkPersistentFlagRequired("ValueA")
	addCmd.MarkPersistentFlagRequired("ValueB")

	err := addCmd.PreRunE
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func addTwoNumbersAndDisplay(cmd *cobra.Command) {
	a,_ := strconv.Atoi(cmd.Flag("ValueA").Value.String())
	b,_ := strconv.Atoi(cmd.Flag("ValueB").Value.String())
	fmt.Println("Sum: ", a+b)
}

func checkIfValueAreNotLarge(cmd *cobra.Command) error {
	a,_ := strconv.Atoi(cmd.Flag("ValueA").Value.String())
	if a > 1000 {
		return fmt.Errorf("Too large value of flag: %s : %d", "ValueA", a)
	}
	b,_ := strconv.Atoi(cmd.Flag("ValueB").Value.String())
	if b > 1000 {
		return fmt.Errorf("Too large value of flag: %s : %d", "ValueB", b)
	}
	return nil
}