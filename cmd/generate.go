package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/moderncircuits/paket"
	"github.com/moderncircuits/paket/macos"
	"github.com/moderncircuits/paket/windows/innosetup"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Print current application generate",
	RunE:  runGenerateCommand,
}

func runGenerateCommand(cmd *cobra.Command, args []string) error {
	runner := paket.NewRunner()
	generators := []paket.Generator{
		paket.NullGenerator{},
		&macos.Native{},
		&innosetup.Compiler{},
	}

	if err := registerGenerators(runner, generators); err != nil {
		return err
	}

	config, err := paket.ReadProjectFile("testdata/full.hcl")
	if err != nil {
		return err
	}

	if err := runner.RunTag(*config, "Windows"); err != nil {
		return err
	}

	js, err := json.MarshalIndent(runner, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("test.json", js, 0644)
	if err != nil {
		return err
	}

	return nil
}

func registerGenerators(runner *paket.Runner, generators []paket.Generator) error {
	for _, generator := range generators {
		if err := runner.RegisterGenerator(generator); err != nil {
			return err
		}
	}
	return nil
}
