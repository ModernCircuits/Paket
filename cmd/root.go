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

	js, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("test.json", js, 0644)
	if err != nil {
		return err
	}

	return nil
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
