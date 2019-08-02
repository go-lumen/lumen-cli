package commands

import (
	"github.com/go-lumen/lumen-cli/utils"
	"github.com/spf13/cobra"
	"os"
)

// RouterCmd holds functions to generate router
func RouterCmd(cmd *cobra.Command, args []string) {
	selectedModels := utils.SelectMethodsModels("router")

	utils.GenerateFile("router.tmpl", "server/router.go", selectedModels)

	os.Exit(1)
}
