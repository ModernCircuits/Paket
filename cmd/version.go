package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/moderncircuits/paket/runtime"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// versionCmd represents the version sub command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print current application version",
	RunE:  runVersionCommand,
}

// runVersionCommand ...
func runVersionCommand(cmd *cobra.Command, args []string) error {
	fmt.Fprintf(os.Stdout, "Version: %s\n", runtime.Version)
	return nil
}
