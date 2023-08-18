package commands

import (
	"github.com/spf13/cobra"
)

type Command = cobra.Command

func Run(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use:   "izmir",
	Short: "Open source Cloud Platform as Code Tool",
	Long:  "Izmir offers platform provisioning from configuration files.",
}

var DryRun bool
var ConfigRoot string

func init() {
	SyncCmd.Flags().StringVarP(&ConfigRoot, "root", "r", "", "Root path of the platform configuration files")
	SyncCmd.Flags().BoolVarP(&DryRun, "plan", "p", false, "Execute in dry-run mode and display list of resources")
	DiffCmd.Flags().StringVarP(&ConfigRoot, "root", "r", "", "Root path of the platform configuration files")
	RootCmd.AddCommand(VersionCmd)
	RootCmd.AddCommand(SyncCmd)
	RootCmd.AddCommand(DiffCmd)
}
