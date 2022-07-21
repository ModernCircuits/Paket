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
	project, err := paket.NewProject("testdata/full.hcl")
	if err != nil {
		return err
	}

	generators := []paket.Generator{
		paket.NullGenerator{},
		&macos.Native{},
		&innosetup.Compiler{},
	}

	if err := registerGenerators(project, generators); err != nil {
		return err
	}

	if err := project.RunTag("Windows"); err != nil {
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

func registerGenerators(project *paket.Project, generators []paket.Generator) error {
	for _, generator := range generators {
		if err := project.RegisterGenerator(generator); err != nil {
			return err
		}
	}
	return nil
}
