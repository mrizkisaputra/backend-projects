package commands

import (
	"expense-tracker/cmd/main/repositories"
	"expense-tracker/cmd/main/services"
	"expense-tracker/internal/log"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.Flags().StringVarP(&Category, "category", "c", "", "Category for the expense")
	/* rootCmd --> listCmd */
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Use for view of all expenses",
	RunE:  actionListCmd,
}

func actionListCmd(cmd *cobra.Command, args []string) error {
	expenseRepo := repositories.NewExpenseRepository(log.Log)
	expenseService := services.NewExpenseService(expenseRepo)

	/* list by specific category */
	flagCategory, err := cmd.Flags().GetString("category")
	if err != nil {
		return err
	}

	if flagCategory != "" {
		listByCategory, err := expenseService.GetByCategory(Category)
		if err != nil {
			return err
		}
		fmt.Println(listByCategory)
		return nil
	}

	result, err := expenseService.GetAll()
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
