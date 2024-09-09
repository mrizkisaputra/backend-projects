package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Description string
	Amount      float64
	Category    string
	Month       string
	Export      string
)

var rootCmd = &cobra.Command{
	Use: "expense-tracker",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Expense tracker APP!. ")
		fmt.Println("Run 'expense-tracker --help' for usage")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
