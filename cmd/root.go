package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)


// rootCmd represents the base command when called without any subcommands
// So here the root command is cobra-cli-tutorials
// In rootCmd, just add the description and mention what to do when users enter bare command i.e cobra-cli-tutorials.
// So cobra-cli-tutorials and cobra-cli-tutorials --help will yield same command.

var rootCmd = &cobra.Command{
  Use:   "cobra-cli-tutorials",
  Short: "cobra-cli-tutorials is basic cli utility create for tutorials purpose",
  Long: `cobra-cli-tutorials has few basic cli options.
         One: Say if math is fun or not
         Add: Add two numbers and display the sum`,

  // Uncomment the following line if your bare application
  // has an action associated with it:
  // Usually you dont want to uncomment it.
  // When user issues command: cobra-cli-tutorials , it will automatically display help.

  /*Run: func(cmd *cobra.Command, args []string) {
  	fmt.Println("Something in bare command")
  },*/
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}