package main

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	IndentBlank int = iota
	IndentStick
)

func printOneName(out io.Writer, indents []int, isLast bool, fileinfo fs.FileInfo) {
	// сначала выводим отступы
	for _, indType := range indents {
		switch indType {
		case IndentBlank:
			fmt.Fprint(out, "	")
		case IndentStick:
			fmt.Fprint(out, "│	")
		}
	}

	// затем выводим "подход" к имени файла
	if isLast {
		fmt.Fprint(out, "└───")
	} else {
		fmt.Fprint(out, "├───")
	}

	// само имя файла и размер
	if !fileinfo.IsDir() {
		// размер пишем только для файлов
		sizeString := "empty"
		if fileinfo.Size() != 0 {
			sizeString = fmt.Sprint(fileinfo.Size()) + "b"
		}
		fmt.Fprintf(out, "%s (%s)\n", fileinfo.Name(), sizeString)
	} else {
		fmt.Fprintln(out, fileinfo.Name())
	}
}

func printContent(out io.Writer, path string, indents []int) (err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	for i, file := range files {
		isLast := (i == len(files)-1) // идентификатор последнего элемента в папке
		newIndent := IndentStick
		if isLast {
			newIndent = IndentBlank
		}

		printOneName(out, indents, isLast, file)
		if file.IsDir() {
			printContent(out, filepath.Join(path, file.Name()), append(indents, newIndent))
		}
	}

	return
}

func dirTree(out io.Writer, path string) (err error) {
	return printContent(out, path, []int{})
}

func main() {
	out := os.Stdout
	if len(os.Args) != 2 {
		panic("usage go run main.go .")
	}
	path := os.Args[1]
	err := dirTree(out, path)
	if err != nil {
		panic(err.Error())
	}
}
