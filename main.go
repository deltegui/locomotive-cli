package main

import (
	"flag"
	"fmt"
	"github.com/gobuffalo/packr"
	"html/template"
	"log"
	"os"
)

var (
	files       *template.Template
	projectName string
)

func main() {
	name := flag.String("name", "", "Your project name. Needed")
	projectType := flag.String("type", "mpa", "Project type. Can be 'mpa' 'webpack' or 'api'")
	flag.Parse()
	loadTemplates()
	if len(*name) == 0 {
		log.Fatalln("Invalid project name")
	}
	projectName = *name
	os.Mkdir(projectName, os.ModePerm)
	createDefaultProject()
	switch *projectType {
	case "webpack":
		createWebpackProject()
		break
	case "api":
		createAPIProject()
		break
	default:
		createMpaProject()
	}
}

func loadTemplates() {
	box := packr.NewBox("./files")
	boxFiles := box.List()
	files = template.New("files")
	for _, boxFile := range boxFiles {
		data, err := box.FindString(boxFile)
		if err != nil {
			log.Fatalln("Error loading virtual file: %s\n", err)
		}
		files.Parse(data)
	}
}

func createDefaultProject() {
	createDir("/src")
	createDir("/src/configuration")
	createDir("/src/controllers")
	createDir("/src/domain")

	writeFile("/src/configuration/config.go", "config")

	writeFile("/src/controllers/injector.go", "injector")

	writeFile("/src/domain/error.go", "error")
	writeFile("/src/domain/gateways.go", "gateways")

	writeFile("/config.json", "configjson")
	writeFile("/logo", "logo")

	writeFile("/.gitignore", "gitignore")
}

func createMpaProject() {
	createDir("/static")
	createDir("/templates")
	createDir("/templates/errors")
	writeFile("/templates/errors/notfound.html", "notfound.html")
	writeFile("/makefile", "mpamakefile")
	writeFile("/main.go", "mpamain")
	writeFile("/src/controllers/error.controller.go", "errorcontroller")
}

func createWebpackProject() {
	createDir("/static")
	createDir("/templates")
	createDir("/templates/errors")
	writeFile("/templates/errors/notfound.html", "notfound.html")
	writeFile("/static/index.js", "webpackindexjs")
	writeFile("/makefile", "webpackmakefile")
	writeFile("/main.go", "mpamain")
	writeFile("/package.json", "packagejson")
	writeFile("/webpack.config.js", "webpackconf")
	writeFile("/src/controllers/error.controller.go", "errorcontroller")
}

func createAPIProject() {
	writeFile("/makefile", "mpamakefile")
	writeFile("/main.go", "apimain")
	writeFile("/src/controllers/error.controller.go", "apierrorcontroller")
}

func writeFile(path, template string) {
	output, err := os.Create(fmt.Sprintf("%s%s", projectName, path))
	if err != nil {
		log.Fatalf("Cannot create file: %s\n", err)
	}
	defer output.Close()
	files.ExecuteTemplate(output, template, projectName)
}

func createDir(path string) {
	os.Mkdir(fmt.Sprintf("%s%s", projectName, path), os.ModePerm)
}
