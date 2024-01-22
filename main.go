package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func main() {

	moduleName := flag.String("name", "", "module name")
	flag.Parse()

	if *moduleName == "" {
		log.Fatal("flag name is required")
	}

	moduleDir := "./modules"

	if err := os.Mkdir(moduleDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	fullDir := path.Join(moduleDir, *moduleName)

	if err := os.Mkdir(fullDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	type templateFile struct {
		Kind string
		Path string
	}

	templateFiles := []templateFile{
		{
			Kind: "module",
			Path: "template/module/module.go.template",
		},
		{
			Kind: "dto",
			Path: "template/module/dto.go.template",
		},
		{
			Kind: "service",
			Path: "template/module/service.go.template",
		},
		{
			Kind: "repository",
			Path: "template/module/repository.go.template",
		},
		{
			Kind: "entity",
			Path: "template/module/entity.go.template",
		},
	}

	for _, file := range templateFiles {
		if err := generate(*moduleName, file.Kind, file.Path); err != nil {
			panic(err)
		}
	}

}

func generate(moduleName, kind, templatePath string) error {
	data, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	newData := strings.ReplaceAll(string(data), "{{module_name}}", strings.Title(moduleName))
	newData = strings.ReplaceAll(newData, "{{module_name_lower}}", strings.ToLower(moduleName))

	filePath := fmt.Sprintf("./modules/%s/%s.%s.go", moduleName, moduleName, kind)

	err = os.WriteFile(filePath, []byte(newData), 0644)
	if err != nil {
		return err
	}

	return nil

}
