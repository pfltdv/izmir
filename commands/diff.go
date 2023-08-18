package commands

import (
	"fmt"

	"github.com/pfltdv/izmir/util"
	"github.com/spf13/cobra"
)

var DiffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Calculates the difference between local configuration with remote cloud",
	RunE:  diffCmdF,
}

func diffCmdF(command *cobra.Command, args []string) error {
	if !util.DirExists(ConfigRoot) {
		return fmt.Errorf("Configuration file root does not exists")
	}
	fmt.Fprintln(command.OutOrStdout(), "Platform Definition Folder is", ConfigRoot)

	return nil
}
