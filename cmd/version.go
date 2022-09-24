package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Unknown15082/httpstatus/pkg/general"
)

var versionCommand = &cobra.Command{
	Use: "version",
	Short: "Display current version",
	Run: general.CurrentVersion,
}

func init() {
	rootCommand.AddCommand(versionCommand)
}