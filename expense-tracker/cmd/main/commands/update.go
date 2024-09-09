package commands

import (
	"expense-tracker/cmd/main/repositories"
	"expense-tracker/cmd/main/services"
	"expense-tracker/internal/exceptions"
	"expense-tracker/internal/log"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	updateCmd.Flags().
		Float64VarP(&Amount, "amount", "a", 0, "New amount for the expense")
	updateCmd.Flags().
		StringVarP(&Description, "description", "d", "", "New description for the expense")
	_ = updateCmd.MarkFlagRequired("amount")

	/* rootCmd --> updateCmd */
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Args:  cobra.ExactArgs(1),
	Short: "Use for update a expense",
	RunE:  actionUpdateCmd,
}

func actionUpdateCmd(cmd *cobra.Command, args []string) error {
	expenseRepo := repositories.NewExpenseRepository(log.Log)
	expenseService := services.NewExpenseService(expenseRepo)

	parsedToInt, err := strconv.Atoi(args[0])
	if err != nil {
		return exceptions.NewErrValidation("Please input numeric id")
	}

	result, errUpdate := expenseService.Update(parsedToInt, Description, Amount)
	if errUpdate != nil {
		return errUpdate
	}

	fmt.Println(result)
	return nil
}
