package commands

import (
	"fmt"

	"github.com/pfltdv/izmir/util"
	"github.com/spf13/cobra"
)

var DropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop everything at cloud created by local platform configuration",
	RunE:  dropCmdF,
}

func dropCmdF(command *cobra.Command, args []string) error {
	if !util.DirExists(ConfigRoot) {
		return fmt.Errorf("Configuration file root does not exists")
	}
	fmt.Fprintln(command.OutOrStdout(), "Platform Definition Folder is", ConfigRoot)

	return nil
}
