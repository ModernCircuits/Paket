package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"

	"github.com/moderncircuits/paket"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "paket",
	Short: "paket",
	Long:  `paket https://github.com/ModernCircuits/Paket`,
	RunE:  runRootCommand,
}

// runRootCommand runs if no subcommand was selected.
func runRootCommand(cmd *cobra.Command, args []string) error {
	fmt.Println("Root Command")
	project, err := paket.NewProject("testdata/full.hcl")
	if err != nil {
		return err
	}

	err = project.Run(runtime.GOOS)
	if err != nil {
		return err
	}

	// err = project.Run("darwin")
	// if err != nil {
	// 	return err
	// }

	jsonFile, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("test.json", jsonFile, 0644)
	if err != nil {
		return err
	}

	return nil
}

// execute is the main entry point for the cli
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
