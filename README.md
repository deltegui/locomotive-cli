# Phoenix CLI
In spite of you can create a new phoenix project from scratch, can be a repeating task. You can automate it with phoenix cli

## Install
just run
```
go install github.com/deltegui/phoenix-cli
```

## Usage
To create a new API project run:
```
phoenix-cli -new <<your project name>> -type api
```
or omit -type:
```
phoenix-cli -new <<your project name>>
```

To create a new MPA project do:
```
phoenix-cli -new <<your project name>> -type mpa
```

## Project structure


| Folder/File         | Explaination                                                                                                                               |
|--------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| /src         | Your code lives here
| /src/configuration | Normally you will have only config.go file. Here you can change configuration options that can be passed by command line arguments or config.json |
| /src/controllers   | Here you store controllers and injectors                                                                                                          |
| /src/domain        | Here you store application services and domain entities                                                                                           |
| /static     | Here you can serve your static files. Only appears if you create a MPA project. You need to call EnableStaticServer                                                                         |
| /templates  | Here you can store templates to render. Only appears if you create a MPA project.                                                                                                           |
| /config.json        | This file stores values that will be loaded to configuration struct defined in /src/configuration/config.json.                                    |
| /logo               | An ASCII logo that appears when you run the project. You can disable it deleting EnableLogoFile in main.go                                        |
| /main.go            | Application starting point. Here you can configure phoenix.                                                                                       |
| /Makefile           | Simple makefile to automate tasks                                                                                                                 |