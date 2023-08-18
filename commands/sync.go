package commands

import (
	"fmt"

	"github.com/pfltdv/izmir/util"
	"github.com/spf13/cobra"
)

var SyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync platform configuration to the cloud",
	RunE:  syncCmdF,
}

func syncCmdF(command *cobra.Command, args []string) error {
	if !util.DirExists(ConfigRoot) {
		return fmt.Errorf("Configuration file root does not exists")
	}
	fmt.Fprintln(command.OutOrStdout(), "Platform Definition Folder is", ConfigRoot)

	return nil
}
