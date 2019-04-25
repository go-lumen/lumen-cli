package commands

import (
	"fmt"
	"github.com/go-lumen/lumen/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"

	//"path/filepath"
	"strings"
)

// ReplaceNamespace simply replace oldString by newString in file provided by path
func ReplaceNamespace(path string, oldString string, newString string) error {
	read, err := ioutil.ReadFile(path)
	utils.CheckErr(err)

	newContents := strings.Replace(string(read), oldString, newString, -1)

	err = ioutil.WriteFile(path, []byte(newContents), 0)
	utils.CheckErr(err)

	return nil
}

// BoilerplateCmd holds functions to generate router
func BoilerplateCmd(cmd *cobra.Command, args []string) {
	// Step 1: Checking that argument is in the correct format
	if !(strings.Contains(args[0], `git`) && strings.Contains(args[0], `/`)) {
		fmt.Println(`Name must be like github.com/user/project`)
		os.Exit(1)
	}

	// Step 2: Get lumen directory path from GOPATH environment variable
	searchDir := os.Getenv("GOPATH") + `/src/github.com/go-lumen/lumen-api`
	fmt.Println(`lumen-api dir is: `, searchDir)

	// Step 3: Copying files in a new directory
	currentDir, err := os.Getwd()
	utils.CheckErr(err)

	arguments := strings.Split(args[0], `/`)
	projectDir := filepath.Join(currentDir, arguments[2])
	fmt.Println("Project ", arguments[2], ` to be generated in `, projectDir)

	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		utils.CheckErr(os.Mkdir(projectDir, 755))
	}

	utils.CheckErr(utils.Copy(searchDir, projectDir))

	// Step 4: Select only Go files
	filesList := []string{}
	err = filepath.Walk(projectDir, func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, ".go") { // Selecting only go files
			filesList = append(filesList, path)
		}
		return nil
	})
	utils.CheckErr(err)

	// Step 5: Changing namespace in files
	for _, file := range filesList {
		//fmt.Println(`Changed namespace in: `, file)
		utils.CheckErr(ReplaceNamespace(file, `github.com/go-lumen/lumen-api`, args[0]))
	}

	os.Exit(1)
}
