package commands

import (
	"github.com/pfltdv/izmir/model"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	RunE:  versionCmdF,
}

func versionCmdF(command *cobra.Command, args []string) error {
	return model.PrintVersion(command.OutOrStdout())
}
