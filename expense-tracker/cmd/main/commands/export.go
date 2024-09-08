package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	exportCmd.Flags().StringVarP(&Export, "ext", "e", "csv", "export for the expense")
	_ = exportCmd.MarkFlagRequired("ext")

	/* rootCmd --> exportCmd */
	rootCmd.AddCommand(exportCmd)
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "for export expense",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("coming soon!")
		return nil
	},
}
