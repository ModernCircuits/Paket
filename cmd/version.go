package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/moderncircuits/paket/cmd/runtime"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print current application version",
	RunE:  runVersionCommand,
}

func runVersionCommand(cmd *cobra.Command, args []string) error {
	fmt.Fprintf(os.Stdout, "Version: %s\n", runtime.Version)
	return nil
}
