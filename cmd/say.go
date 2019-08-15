package cmd

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// sayCmd represents the say command
var mathFunCmd = &cobra.Command{
	Use:   "math-fun",
	Short: "It will say if math is fun or not with 50% probability.",
	Long: `Longer description: It will say if math is fun or not. 
    It will say in random order: 50% of time it is fun and 50% time it is not fun.
	It will use rand package to calculate that random value.`,
	// Mention below: Which function you want to call when user issue command: math-fun.
	Run: func(cmd *cobra.Command, args []string) {
		err := isMathFun(cmd)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	// Add mathFunCmd command in root cli tool i.e cobra-cli-tutorials
	rootCmd.AddCommand(mathFunCmd)

	// Add any flags you want to add.
	// cobra-cli-tutorials math-fun -c=[true|false]
	// cobra-cli-tutorials math-fun --caps=[true|false]
	mathFunCmd.PersistentFlags().BoolP("caps", "c", true, "It will control if message printed in small or big caps")

	//cobra-cli-tutorials math-fun -a "this is for tutorials"
	//cobra-cli-tutorials math-fun --append "this is for tutorials"
	// -a="Hello"   -a "Hello" are equivalent command.
	mathFunCmd.Flags().StringP("append", "a", "", "Append some message in last")
}

func isMathFun(cmd *cobra.Command) error {
	rand.Seed(time.Now().UnixNano())
	var mathVal string
	val := rand.Intn(10000)
	if val % 2 == 0 {
		mathVal = "Math is fun"
	} else {
		mathVal = "Math is NOT fun"
	}
	capFlag := cmd.Flag("caps")
	fmt.Println("Debug:", capFlag.Value)
	capValue := capFlag.Value.String()

	if capFlag.Value.String() == "" {
		//do nothing.
	} else {
		boolVal,err := strconv.ParseBool(capValue)
		if err != nil {
			return fmt.Errorf("Invalid boolean value of %s: %s, Usage: %s", capFlag.Name, capFlag.Value, capFlag.Usage)
		}
		if boolVal {
			mathVal = strings.ToUpper(mathVal)
		} else {
			mathVal = strings.ToLower(mathVal)
		}
	}
	mathVal = mathVal + cmd.Flag("append").Value.String()
	fmt.Println(mathVal)
	return nil
}
