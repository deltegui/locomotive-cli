package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const header string = `
package store

var store map[string]string = map[string]string {
`

const end string = `}`

type dataFile struct {
	file *os.File
}

func newDataFile() dataFile {
	output, err := os.Create("./files/data.go")
	if err != nil {
		log.Fatalf("Cannot create file: %s\n", err)
	}
	return dataFile{output}
}

func (f dataFile) write(str string) {
	if _, err := f.file.Write([]byte(str)); err != nil {
		panic(err)
	}
}

func (f dataFile) close() {
	f.file.Close()
}

func main() {
	data := newDataFile()
	defer data.close()
	data.write(header)
	filepath.Walk("./rawfiles", func(path string, info os.FileInfo, err error) error {
		if path == "./rawfiles" {
			return nil
		}
		log.Println(path)
		if err != nil {
			panic(err)
		}
		content, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		var entry string = fmt.Sprintf("\"%s\": `%s`,\n", info.Name(), content)
		data.write(entry)
		return nil
	})
	data.write(end)
}
