package main

import (
	"github.com/go-lumen/lumen/commands"
	"github.com/go-lumen/lumen/utils"
	"github.com/spf13/cobra"
)

func main() {
	var cmdNew = &cobra.Command{
		Use:   `new [namespace]`,
		Short: `Generating files`,
		Long:  `Generates files from lumen-api`,
		Args:  cobra.MinimumNArgs(1),
		Run:   commands.BoilerplateCmd,
	}
	var cmdModel = &cobra.Command{
		Use:   `model [name]`,
		Short: `Generating model`,
		Long:  `Generates a model`,
		Args:  cobra.MinimumNArgs(0),
		Run:   commands.ModelCmd,
	}

	var cmdGenerate = &cobra.Command{
		Use:   `generate [name]`,
		Short: `Generating everything`,
		Long:  `Generating everything from models and methods`,
		Args:  cobra.MinimumNArgs(0),
		Run:   commands.GenerateCmd,
	}

	var cmdController = &cobra.Command{
		Use:   `controller [name]`,
		Short: `Generating controller`,
		Long:  `Generates a controller with multiple methods`,
		Args:  cobra.MinimumNArgs(0),
		Run:   commands.ControllerCmd,
	}

	var cmdRouter = &cobra.Command{
		Use:   `router [name]`,
		Short: `Generating router`,
		Long:  `Generates router methods`,
		Args:  cobra.MinimumNArgs(0),
		Run:   commands.RouterCmd,
	}

	var cmdStore = &cobra.Command{
		Use:   `store [name]`,
		Short: `Generating store`,
		Long:  `Generates store methods`,
		Args:  cobra.MinimumNArgs(0),
		Run:   commands.StoreCmd,
	}

	var rootCmd = &cobra.Command{
		Use:   `lumen`,
		Short: `Lumen is a CLI that helps you generating API code for lumen-api`}
	rootCmd.AddCommand(cmdNew, cmdModel, cmdGenerate, cmdController, cmdRouter, cmdStore)

	err := rootCmd.Execute()
	utils.CheckErr(err)
}
