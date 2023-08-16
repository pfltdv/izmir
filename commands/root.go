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
	Short: "Open source Cloud IaC Automation Tool",
	Long:  "Izmir offers cloud provisioning with configuration files.",
}
