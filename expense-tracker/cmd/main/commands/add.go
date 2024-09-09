package commands

import (
	"expense-tracker/cmd/main/repositories"
	"expense-tracker/cmd/main/services"
	"expense-tracker/internal/log"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.Flags().
		StringVarP(&Description, "description", "d", "", "Description for the expense (mandatory)")
	addCmd.Flags().
		Float64VarP(&Amount, "amount", "a", 0, "Amount for the expense (mandatory)")
	addCmd.Flags().
		StringVarP(&Category, "category", "c", "default", "Category for the expense (optional)")

	/* mandatory flags */
	_ = addCmd.MarkFlagRequired("description")
	_ = addCmd.MarkFlagRequired("amount")

	/* rootCmd --> addCmd */
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Use for insert a new expense",
	RunE:  actionAddCmd,
}

func actionAddCmd(cmd *cobra.Command, args []string) error {
	expenseRepo := repositories.NewExpenseRepository(log.Log)
	expenseService := services.NewExpenseService(expenseRepo)

	result, err := expenseService.Add(Description, Amount, Category)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
