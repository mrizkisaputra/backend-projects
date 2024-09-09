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
	/* rootCmd --> deleteCmd */
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.ExactArgs(1),
	Short: "Use for delete a expense",
	RunE:  actionDeleteCmd,
}

func actionDeleteCmd(cmd *cobra.Command, args []string) error {
	expenseRepo := repositories.NewExpenseRepository(log.Log)
	expenseService := services.NewExpenseService(expenseRepo)
	parsedToInt, err := strconv.Atoi(args[0])
	if err != nil {
		return exceptions.NewErrValidation("Please input numeric id")
	}
	result, errDelete := expenseService.Delete(parsedToInt)
	if errDelete != nil {
		return errDelete
	}
	fmt.Println(result)
	return nil
}
