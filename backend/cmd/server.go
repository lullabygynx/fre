package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "starts the fre backend server",
	RunE: func(cmd *cobra.Command, agrs []string) error {
		fmt.Println("You have hit the server cmd")
		return nil
	},
}
