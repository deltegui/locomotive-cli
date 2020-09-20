package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"os"

	"github.com/deltegui/phoenix-cli/store"
)

//go:generate go run ./generators/files.go

const version string = "0.2.1"

var projectName string

func main() {
	versionFlag := flag.Bool("v", false, "Shows phoenix-cli version")
	name := flag.String("new", "", "Creates new project. Usage: phoenix-cli -new [your project name]")
	projectType := flag.String("type", "api", "Project type. Can be 'mpa' 'webpack' or 'api'")
	flag.Parse()
	printLogo()
	if *versionFlag {
		fmt.Printf("phoenix-cli v%s\n", version)
		os.Exit(0)
	}
	if len(*name) == 0 {
		fmt.Println("Invalid project name. Usage:")
		flag.PrintDefaults()
		os.Exit(1)
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
	case "mpa":
		createMpaProject()
		break
	default:
		fmt.Println("Invalid project type. Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println("You are ready to GO!")
}

func createDefaultProject() {
	fmt.Println("Generating project...")
	createDir("/src")
	createDir("/src/configuration")
	createDir("/src/controllers")
	createDir("/src/domain")

	writeFile("/src/configuration/config.go", "config")

	writeFile("/src/controllers/injector.go", "injector")

	writeFile("/src/domain/error.go", "error")
	writeFile("/src/domain/gateways.go", "gateways")

	writeFile("/config.json", "configjson")
	writeLogo()

	writeFile("/.gitignore", "gitignore")
}

func createMpaProject() {
	fmt.Println("Creating MPA project...")
	createDir("/static")
	createDir("/templates")
	createDir("/templates/errors")
	writeFile("/templates/errors/notfound.html", "notfound.html")
	writeFile("/makefile", "mpamakefile")
	writeFile("/main.go", "mpamain")
	writeFile("/src/controllers/error.controller.go", "errorcontroller")
}

func createWebpackProject() {
	fmt.Println("Creating Webpack project...")
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
	fmt.Println("Creating API project...")
	writeFile("/makefile", "mpamakefile")
	writeFile("/main.go", "apimain")
	writeFile("/src/controllers/error.controller.go", "apierrorcontroller")
}

func writeFile(path, templName string) {
	output, err := os.Create(fmt.Sprintf("%s%s", projectName, path))
	if err != nil {
		fmt.Printf("Cannot create file: %s\n", err)
		os.Exit(2)
	}
	defer output.Close()
	tmpl := template.New("a")
	tmpl, err = tmpl.Parse(store.Get(templName))
	if err != nil {
		panic(err)
	}
	tmpl.Execute(output, projectName)
}

func writeLogo() {
	output, err := os.Create(fmt.Sprintf("%s/%s", projectName, "logo"))
	if err != nil {
		fmt.Printf("Cannot create file: %s\n", err)
		os.Exit(2)
	}
	defer output.Close()
	tmpl := template.New("a")
	tmpl, err = tmpl.Parse(genLogo())
	if err != nil {
		panic(err)
	}
	tmpl.Execute(output, projectName)
}

func createDir(path string) {
	os.Mkdir(fmt.Sprintf("%s%s", projectName, path), os.ModePerm)
}

func printLogo() {
	fmt.Println(genLogo())
}

func genLogo() string {
	var b bytes.Buffer
	p := b.WriteString
	p("                (                           )\n")
	p("          ) )( (                           ( ) )( (\n")
	p("       ( ( ( )  ) )                     ( (   (  ) )(\n")
	p("      ) )     ,,\\\\\\                     ///,,       ) (\n")
	p("   (  ((    (\\\\\\\\//                     \\\\////)      )\n")
	p("    ) )    (-(__//                       \\\\__)-)     (\n")
	p("   (((   ((-(__||                         ||__)-))    ) )\n")
	p("  ) )   ((-(-(_||           ```\\__        ||_)-)-))   ((\n")
	p("  ((   ((-(-(/(/\\        ''; 9.- `      //\\)\\)-)-))    )\n")
	p("   )   (-(-(/(/(/\\      '';;;;-\\~      //\\)\\)\\)-)-)   (   )\n")
	p("(  (   ((-(-(/(/(/\\======,:;:;:;:,======/\\)\\)\\)-)-))   )\n")
	p("    )  '(((-(/(/(/(//////:%%%%%%%:\\\\\\\\\\\\)\\)\\)\\)-)))`  ( (\n")
	p("   ((   '((-(/(/(/('uuuu:WWWWWWWWW:uuuu`)\\)\\)\\)-))`    )\n")
	p("     ))  '((-(/(/(/('|||:wwwwwwwww:|||')\\)\\)\\)-))`    ((\n")
	p("  (   ((   '((((/(/('uuu:WWWWWWWWW:uuu`)\\)\\))))`     ))\n")
	p("        ))   '':::UUUUUU:wwwwwwwww:UUUUUU:::``     ((   )\n")
	p("          ((      '''''''\\uuuuuuuu/``````         ))\n")
	p("           ))            `JJJJJJJJJ`           ((\n")
	p("             ((            LLLLLLLLLLL         ))\n")
	p("               ))         ///|||||||\\\\\\       ((\n")
	p("                 ))      (/(/(/(^)\\)\\)\\)       ((\n")
	p("                  ((                           ))\n")
	p("                    ((                       ((\n")
	p("                      ( )( ))( ( ( ) )( ) (()\n")
	p("")
	p("@@@@@@@   @@@  @@@   @@@@@@   @@@@@@@@  @@@  @@@  @@@  @@@  @@@\n")
	p("@@@@@@@@  @@@  @@@  @@@@@@@@  @@@@@@@@  @@@@ @@@  @@@  @@@  @@@\n")
	p("@@!  @@@  @@!  @@@  @@!  @@@  @@!       @@!@!@@@  @@!  @@!  !@@\n")
	p("!@!  @!@  !@!  @!@  !@!  @!@  !@!       !@!!@!@!  !@!  !@!  @!!\n")
	p("@!@@!@!   @!@!@!@!  @!@  !@!  @!!!:!    @!@ !!@!  !!@   !@@!@!\n")
	p("!!@!!!    !!!@!!!!  !@!  !!!  !!!!!:    !@!  !!!  !!!    @!!!\n")
	p("!!:       !!:  !!!  !!:  !!!  !!:       !!:  !!!  !!:   !: :!!\n")
	p(":!:       :!:  !:!  :!:  !:!  :!:       :!:  !:!  :!:  :!:  !:!\n")
	p(" ::       ::   :::  ::::: ::   :: ::::   ::   ::   ::   ::  :::\n")
	p(" :         :   : :   : :  :   : :: ::   ::    :   :     :   ::\n")
	return b.String()
}
