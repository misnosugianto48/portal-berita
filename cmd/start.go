package cmd

import (
	"portal-berita/internal/apps"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  "start",
	Run:   func(cmd *cobra.Command, args []string) { apps.RunServer() },
}

func init() {
	rootCmd.AddCommand(startCmd)
}
