package main

import (
	"github.com/pfltdv/izmir/config"
	"github.com/pfltdv/izmir/model"
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

var DiffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Calculates the difference between local configuration with remote cloud",
	RunE:  izmirCmdF,
}

var DropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop everything at cloud created by local platform configuration",
	RunE:  izmirCmdF,
}

var SyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync platform configuration to the cloud",
	RunE:  izmirCmdF,
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	RunE:  versionCmdF,
}

var DryRun bool
var MainConfigRoot string
var ConfigRoot string

func init() {
	RootCmd.PersistentFlags().StringVarP(&MainConfigRoot, "config", "c", "", "Izmir configuration file location")

	SyncCmd.Flags().StringVarP(&ConfigRoot, "root", "r", "", "Root path of the platform configuration files")
	SyncCmd.Flags().BoolVarP(&DryRun, "plan", "p", false, "Execute in dry-run mode and display list of resources")

	DiffCmd.Flags().StringVarP(&ConfigRoot, "root", "r", "", "Root path of the platform configuration files")
	DropCmd.Flags().StringVarP(&ConfigRoot, "root", "r", "", "Root path of the platform configuration files")

	DropCmd.Flags().BoolVarP(&DryRun, "plan", "p", false, "Execute in dry-run mode and display list of resources")

	RootCmd.AddCommand(DiffCmd)
	RootCmd.AddCommand(DropCmd)
	RootCmd.AddCommand(SyncCmd)
	RootCmd.AddCommand(VersionCmd)
}

func versionCmdF(command *cobra.Command, args []string) error {
	return model.PrintVersion(command.OutOrStdout())
}

func izmirCmdF(command *cobra.Command, args []string) error {
	_, err := config.Load(MainConfigRoot)
	/*	if err != nil {

		}*/
	return err
}
