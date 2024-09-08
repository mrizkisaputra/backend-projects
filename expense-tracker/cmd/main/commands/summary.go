package commands

import (
	"expense-tracker/cmd/main/repositories"
	"expense-tracker/cmd/main/services"
	"expense-tracker/internal/log"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	summaryCmd.Flags().StringVarP(&Month, "month", "m", "", "Month for the expense")

	/* rootCmd --> summaryCmd */
	rootCmd.AddCommand(summaryCmd)
}

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Use for view of all summary expenses",
	RunE:  actionSummaryCmd,
}

func actionSummaryCmd(cmd *cobra.Command, args []string) error {
	expenseRepo := repositories.NewExpenseRepository(log.Log)
	expenseService := services.NewExpenseService(expenseRepo)

	/* summary by specific monthly */
	flagMonth, err := cmd.Flags().GetString("month")
	if err != nil {
		return err
	}

	if flagMonth != "" {
		summaryByMonth, err := expenseService.GetSummaryByMonthCurrentYear(Month)
		if err != nil {
			return err
		}
		fmt.Println(summaryByMonth)
		return nil
	}

	/* summary */
	summary, errSummary := expenseService.GetSummary()
	if errSummary != nil {
		return errSummary
	}
	fmt.Println(summary)
	return nil
}
