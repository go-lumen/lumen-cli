package commands

import (
	"bytes"
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/adrien3d/lumen/utils"
	"github.com/gedex/inflector"
	"github.com/spf13/cobra"
	"go/format"
	"gopkg.in/godo.v2/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type SelectedMethods struct {
	ModelName string
	Methods   []*Method
}

type Method struct {
	Name string
}

func generateControllerFile(filename string, selected []string) {
	fmt.Println("Filename for controller:", filename)

	methods := []*Method{}
	selectedMethods := SelectedMethods{strings.ToLower(filename), methods}
	for _, methodName := range selected {
		methodName = methodName[0 : len(methodName)-len(filename)]
		selectedMethods.Methods = append(selectedMethods.Methods, &Method{methodName})
	}

	path := filepath.Join("templates", "controller.tmpl")
	body, _ := ioutil.ReadFile(path)
	tmpl, err := template.New("model").Funcs(funcMap).Parse(string(body))
	utils.Check(err)

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, selectedMethods)
	utils.Check(err)

	src, _ := format.Source(buf.Bytes())
	dstPath := filepath.Join("generated/controllers/", strings.ToLower(filename)+".go")

	if !util.FileExists(filepath.Dir(dstPath)) {
		if err := os.Mkdir(filepath.Dir(dstPath), 0644); err != nil {
			fmt.Println(err)
		}
	}
	if err := ioutil.WriteFile(dstPath, src, 0644); err != nil {
		fmt.Println(err)
	}
}

func ControllerCmd(cmd *cobra.Command, args []string) {
	shell := ishell.New()

	shell.Println("Generating Controller")

	files, err := ioutil.ReadDir("generated/models")
	if err != nil {
		fmt.Println(err)
	}

	var fileNames []string
	for _, file := range files {
		fmt.Println(file.Name())
		fileNames = append(fileNames, file.Name())
	}

	// Step 1: Ask user which model to use
	choice := shell.MultiChoice(fileNames, "Please select the model you want to generate the matching controller:")
	fileNames[choice] = strings.Title(inflector.Singularize(strings.TrimRight(fileNames[choice], ".go")))

	//Step 2
	methods := []string{"Create" + fileNames[choice], "Get" + fileNames[choice], "GetAll" + fileNames[choice] + "", "Update" + fileNames[choice], "Delete" + fileNames[choice]}
	choices := shell.Checklist(methods,
		"What method do you want to implement ? (space to select/deselect)",
		nil)
	out := func() (c []string) {
		for _, v := range choices {
			c = append(c, methods[v])
		}
		return
	}
	selectedMethods := out()

	generateControllerFile(fileNames[choice], selectedMethods)
	os.Exit(1)

	shell.Run()
}
