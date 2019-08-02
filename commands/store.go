package commands

import (
	"github.com/go-lumen/lumen-cli/utils"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// StoreCmd holds functions to generate store
func StoreCmd(cmd *cobra.Command, args []string) {
	selectedModels := utils.SelectMethodsModels("store")

	// Step 1: Generate store.go that indexes methods
	utils.GenerateFile("index.store.tmpl", "store/store.go", selectedModels)

	for _, selectedModel := range selectedModels {
		// Step 2: Generate interfaces in entity.go
		utils.GenerateFile("entity.store.tmpl", "store/"+strings.ToLower(selectedModel.ModelName)+".go", selectedModel)

		// Step 3: Generate mongo methods in mongodb/entity.go
		utils.GenerateFile("mongo.store.tmpl", "store/mongodb/"+strings.ToLower(selectedModel.ModelName)+".go", selectedModel)
	}

	os.Exit(1)
}
