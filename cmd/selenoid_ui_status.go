package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var selenoidUIStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Shows Selenoid UI status",
	Run: func(cmd *cobra.Command, args []string) {
		lifecycle, err := createLifecycle(uiConfigDir)
		if err != nil {
			stderr("Failed to initialize: %v\n", err)
			os.Exit(1)
		}
		lifecycle.UIStatus()
	},
}
