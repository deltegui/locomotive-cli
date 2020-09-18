package store

//store THIS FILE IS GENERATED WITH GO:GENERATE.

var store map[string]string = map[string]string{
	"apierrorcontroller": `package controllers

import (
	"net/http"

	"github.com/deltegui/phoenix"
)

func NotFoundError() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		phoenix.NewJSONPresenter(w).Present(struct{ Code string }{Code: "404"})
	}
}
`,
	"apimain": `package main

import (
	"{{.}}/src/configuration"
	"{{.}}/src/controllers"

	"github.com/deltegui/phoenix"
)

func main() {
	app := phoenix.NewApp()
	app.Configure().
		SetProjectInfo("{{.}}", "0.1.0").
		EnableLogoFile()
	config := configuration.Load()
	controllers.Register(app)
	app.Run(config.ListenURL)
}

`,
	"config": `package configuration

import "github.com/deltegui/configloader"

//Configuration representation of json config file
type Configuration struct {
	ListenURL string ` + "`paramName:\"url\"`" + `
}

//Load configuration from config.json file and overwrite
//default values if console params are provided
func Load() Configuration {
	return *configloader.NewConfigLoaderFor(&Configuration{}).
		AddHook(configloader.CreateFileHook("./config.json")).
		AddHook(configloader.CreateParamsHook()).
		Retrieve().(*Configuration)
}`,
	"configjson": `{
    "ListenURL": "localhost:8080"
}`,
	"error": `package domain

import "fmt"

// UseCaseError is an error that can return a UseCase
type UseCaseError struct {
	Code   uint16
	Reason string
	Fix    string
}

func (caseErr UseCaseError) Error() string {
	return fmt.Sprintf("UseCaseError -> [%d] %s: %s", caseErr.Code, caseErr.Reason, caseErr.Fix)
}

var (
	MalformedRequestErr = UseCaseError{Code: 000, Reason: "Bad request", Fix: "See documentation and try again"}
	InternalError       = UseCaseError{Code: 001, Reason: "Internal Error", Fix: ""}
)
`,
	"errorcontroller": `package controllers

import (
	"net/http"

	"github.com/deltegui/phoenix"
)

func NotFoundError() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		phoenix.NewHTMLPresenter(w, "notfound.html").Present(nil)
	}
}
`,
	"gateways": `package domain

import (
	"net/http"

	"github.com/deltegui/phoenix"
)

type UseCaseRequest interface{}

type UseCaseResponse interface{}

var EmptyRequest UseCaseRequest = struct{}{}

type UseCase func(UseCaseRequest) UseCaseResponse
`,
	"gitignore": `.DS_Store
node_modules
build
/static/bundle.js`,
	"injector": `package controllers

import (
	"github.com/deltegui/phoenix"
)

func Register(app phoenix.App) {
	app.MapRoot(NotFoundError)
}
`,
	"mpamain": `package main

import (
	"{{.}}/src/configuration"
	"{{.}}/src/controllers"

	"github.com/deltegui/phoenix"
)

func main() {
	app := phoenix.NewApp()
	app.Configure().
		SetProjectInfo("{{.}}", "0.1.0").
		EnableLogoFile().
		EnableStaticServer().
		EnableSessions()
	config := configuration.Load()
	controllers.Register(app)
	app.Run(config.ListenURL)
}
`,
	"mpamakefile": `build:
	mkdir ./build
	go build -o ./build/{{.}} ./main.go

clean:
	rm -rf ./build

watch:
	reflex -r '(.go|.html)' -s -- sh -c 'go run ./main.go'`,
	"notfound.html": `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{{.}}</title>
</head>
<body>
    404
</body>
</html>`,
}
