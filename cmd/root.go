package main

import (
	"errors"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "paket",
	Short: "paket",
	Long:  `paket https://github.com/ModernCircuits/Paket`,
	RunE:  runRootCommand,
}

func runRootCommand(cmd *cobra.Command, args []string) error {
	return errors.New("no subcommand used")
}

func execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolP("json", "j", false, "json output")
}
