package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/FrameParadorn/firefin-cli/template"
)

func main() {

	moduleName := flag.String("name", "", "module name")
	flag.Parse()

	if *moduleName == "" {
		log.Fatal("flag name is required")
	}

	moduleDir := "./modules"

	if err := os.Mkdir(moduleDir, os.ModePerm); err != nil {
		if !strings.Contains(err.Error(), "file exists") {
			log.Fatal(err)
		}
	}

	fullDir := path.Join(moduleDir, *moduleName)

	if err := os.Mkdir(fullDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	type templateFile struct {
		Kind string
		Data string
	}

	templateFiles := []templateFile{
		{
			Kind: "module",
			Data: template.Module,
		},
		{
			Kind: "controller",
			Data: template.Controller,
		},
		{
			Kind: "dto",
			Data: template.Dto,
		},
		{
			Kind: "service",
			Data: template.Service,
		},
		{
			Kind: "repository",
			Data: template.Repository,
		},
		{
			Kind: "entity",
			Data: template.Etity,
		},
	}

	for _, file := range templateFiles {
		if err := generate(*moduleName, file.Kind, file.Data); err != nil {
			panic(err)
		}
	}

}

func generate(moduleName, kind, data string) error {

	newData := strings.ReplaceAll(string(data), "{{module_name}}", strings.Title(moduleName))
	newData = strings.ReplaceAll(newData, "{{module_name_lower}}", strings.ToLower(moduleName))

	filePath := fmt.Sprintf("./modules/%s/%s.%s.go", moduleName, moduleName, kind)

	err := os.WriteFile(filePath, []byte(newData), 0644)
	if err != nil {
		return err
	}

	return nil

}
