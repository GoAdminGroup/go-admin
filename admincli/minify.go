package main

import (
	"bytes"
	"fmt"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/js"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
)

func CSS(inputDir, outputFile string) {
	err := removeOutputFile(outputFile)
	if err != nil {
		log.Panicln("removeOutputFileError", err)
		return
	}

	files, err := getInputFiles(inputDir)
	if err != nil {
		log.Panicln("getInputFilesError", err)
		return
	}

	notMinifiedString, err := combineFiles(files, inputDir)
	if err != nil {
		log.Panicln("combineFilesError", err)
		return
	}

	minifiedString, err := makeMini(notMinifiedString, "text/css")
	if err != nil {
		log.Panicln("doTheMinifyingError", err)
		return
	}

	err = writeOutputFile(minifiedString, outputFile)
	if err != nil {
		log.Panicln("writeOutputFileError", err)
		return
	}
}

func JS(inputDir, outputFile string) {
	err := removeOutputFile(outputFile)
	if err != nil {
		log.Panicln("removeOutputFileError", err)
		return
	}

	files, err := getInputFiles(inputDir)
	if err != nil {
		log.Panicln("getInputFilesError", err)
		return
	}

	var b bytes.Buffer

	for _, name := range files {

		if path.Ext(name) != ".js" {
			continue
		}

		filepath := inputDir + name
		fileTxt, err := ioutil.ReadFile(filepath)
		if err != nil {
			return
		}

		fmt.Println("filepath", filepath)

		m := minify.New()
		m.AddFunc("text/javascript", js.Minify)

		minifiedString, err := m.Bytes("text/javascript", fileTxt)
		if err != nil {
			return
		}

		_, err = b.Write(minifiedString)
		if err != nil {
			return
		}
	}

	err = writeOutputFile(b.String(), outputFile)
	if err != nil {
		checkError(err)
		return
	}
}

func removeOutputFile(outputFile string) error {
	if _, err := os.Stat(outputFile); err == nil {

		if err := os.Remove(outputFile); err != nil {
			return err
		}

		return nil
	}

	return nil
}

func getInputFiles(inputDir string) ([]string, error) {
	filenames := make([]string, 0, 1)

	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		log.Panicln("readInputDirError", err)
		return filenames, err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		filenames = append(filenames, f.Name())
	}

	sort.Strings(filenames)

	return filenames, nil
}

func combineFiles(filenames []string, inputDir string) (string, error) {
	var b bytes.Buffer

	for _, name := range filenames {

		if path.Ext(name) != ".css" {
			continue
		}

		filepath := inputDir + name
		fileTxt, err := ioutil.ReadFile(filepath)
		if err != nil {
			return "", err
		}

		fmt.Println("filepath", filepath)

		_, err = b.Write(fileTxt)
		if err != nil {
			return "", nil
		}
	}

	combinedFiles := b.String()

	return combinedFiles, nil
}

func makeMini(notMinifiedString, fileType string) (string, error) {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/javascript", js.Minify)

	minifiedString, err := m.String(fileType, notMinifiedString)
	if err != nil {
		return "", err
	}

	return minifiedString, nil
}

func writeOutputFile(outputText, outputFile string) error {
	err := ioutil.WriteFile(outputFile, []byte(outputText), 0644)
	if err != nil {
		return err
	}

	return nil
}
