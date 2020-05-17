package store

//store THIS FILE IS GENERATED WITH GO:GENERATE.

var store map[string]string = map[string]string{
	"apierrorcontroller": `package controllers

import (
	"{{.}}/src/domain"
	"net/http"
)

func NotFoundError() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		domain.NewJSONPresenter(w).Present(struct{ Code string }{Code: "404"})
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
	phoenix.Configure().
		SetProjectInfo("{{.}}", "0.1.0").
		EnableLogoFile()
	config := configuration.Load()
	controllers.Register()
	phoenix.Run(config.ListenURL)
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
		renderer := phoenix.HTMLRenderer{w}
		renderer.RenderWithMeta(nil, phoenix.RenderMetadata{"notfound.html"})
	}
}
`,
	"gateways": `package domain

import (
	"net/http"

	"github.com/deltegui/phoenix"
)

type Presenter interface {
	Present(data interface{})
	PresentError(data error)
}

type BridgePresenter struct {
	Renderer phoenix.Renderer
}

func (b BridgePresenter) Present(data interface{}) {
	b.Renderer.Render(data)
}

func (b BridgePresenter) PresentError(err error) {
	b.Renderer.RenderError(err)
}

type JSONPresenter struct {
	BridgePresenter
}

func NewJSONPresenter(w http.ResponseWriter) JSONPresenter {
	return JSONPresenter{
		BridgePresenter: BridgePresenter{phoenix.JSONRenderer{w}},
	}
}

type UseCaseRequest interface{}

var EmptyRequest UseCaseRequest = struct{}{}

type UseCase interface {
	Exec(Presenter, UseCaseRequest)
}
`,
	"gitignore": `.DS_Store
node_modules
build
/static/bundle.js`,
	"injector": `package controllers

import (
	"github.com/deltegui/phoenix"
)

func Register() {
	phoenix.MapRoot(NotFoundError)
}
`,
	"mpamain": `package main

import (
	"{{.}}/src/configuration"
	"{{.}}/src/controllers"

	"github.com/deltegui/phoenix"
)

func main() {
	phoenix.Configure().
		SetProjectInfo("{{.}}", "0.1.0").
		EnableLogoFile().
		EnableStaticServer().
		EnableTemplates()
	config := configuration.Load()
	controllers.Register()
	phoenix.Run(config.ListenURL)
}
`,
	"mpamakefile": `build:
	mkdir ./build
	go build -o ./build/fynd ./main.go

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
	"packagejson": `{
  "name": "{{.}}",
  "version": "0.1.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "author": "",
  "license": "ISC",
  "devDependencies": {
    "@babel/core": "^7.7.5",
    "@babel/preset-env": "^7.7.5",
    "babel-loader": "^8.0.6",
    "babel-minify-webpack-plugin": "^0.3.1",
    "css-loader": "^3.2.1",
    "style-loader": "^1.0.1",
    "webpack": "^4.41.2",
    "webpack-cli": "^3.3.10"
  }
}`,
	"webpackconf": `const path = require('path');
const MinifyPlugin = require('babel-minify-webpack-plugin')

console.log((process.env.pro) ? 'production' : 'development');

const outputFolder = (process.env.pro) ? './build/static/' : './static'

module.exports = {
  mode: (process.env.pro) ? 'production' : 'development',
  entry: './static/index.js',
  output: {
    filename: 'bundle.js',
    path: path.join(__dirname, outputFolder),
  },
  module: {
    rules: [
      {
        test: /\.css$/,
        use: [
          'style-loader',
          'css-loader',
        ],
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: [
              '@babel/preset-env',
            ],
          },
        },
      }
    ],
  },
  plugins: [new MinifyPlugin()],
};`,
	"webpackindexjs": ``,
	"webpackmakefile": `build:
	mkdir ./build
	go build -o ./build/fynd ./main.go
	pro=pro node ./node_modules/webpack/bin/webpack.js --config ./webpack.config.js

clean: remove-dev-assets
	rm -rf ./build

dev-assets:
	node ./node_modules/webpack/bin/webpack.js --config ./webpack.config.js

remove-dev-assets:
	rm -rf ./static/bundle.js

watch: dev-assets
	reflex -r '(.go|.html)' -s -R 'node_modules' -- sh -c 'go run ./main.go'`,
}
