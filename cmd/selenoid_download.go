package cmd

import (
	"github.com/aerokube/cm/selenoid"
	"github.com/spf13/cobra"
	"os"
)

var selenoidDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download Selenoid latest or specified release",
	Run: func(cmd *cobra.Command, args []string) {
		downloadImpl(configDir, func(lc *selenoid.Lifecycle) error {
			return lc.Download()
		})
	},
}

func downloadImpl(configDir string, downloadAction func(*selenoid.Lifecycle) error) {
	lifecycle, err := createLifecycle(configDir)
	if err != nil {
		stderr("Failed to initialize: %v\n", err)
		os.Exit(1)
	}
	err = downloadAction(lifecycle)
	if err != nil {
		lifecycle.Printf("Failed to download: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
