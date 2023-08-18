package commands

import (
	"fmt"
	"os"

	"github.com/pfltdv/izmir/util"
	"github.com/spf13/cobra"
)

var SyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Display version information",
	RunE:  syncCmdF,
}

var DryRun bool
var ConfigRoot string

func init() {
	SyncCmd.Flags().StringVarP(&ConfigRoot, "root-path", "r", "", "Root path of the platform configuration files")
	SyncCmd.Flags().BoolVarP(&DryRun, "plan", "p", false, "Execute in dry-run mode and display list of resources")
	RootCmd.AddCommand(SyncCmd)
}

func validate() error {
	if !util.DirExists(ConfigRoot) {
		return fmt.Errorf("Configuration file root does not exists")
	}
	return nil
}

func syncCmdF(command *cobra.Command, args []string) error {
	if err := validate(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(command.OutOrStdout(), "Platform Definition Folder is", ConfigRoot)

	return nil
}
