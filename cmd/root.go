package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use: "httpstatus",
	Short: "Display HTTP status codes",
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}