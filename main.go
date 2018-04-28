package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

var (
	fileName string
	output   []byte
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		check(err)
	}
}

func writeToFile(path string) {
	err := ioutil.WriteFile(path, output, 0644)
	check(err)
}

func createFolder() {
	createDirIfNotExist("./views")
	outputFileName := strings.Split(fileName, ".")[0]
	if outputFileName == "index" {
		path := "./views/" + outputFileName + ".html"
		writeToFile(path)
	} else {
		dir := "./views/" + outputFileName
		createDirIfNotExist(dir)
		path := dir + "/index.html"
		writeToFile(path)
	}
}

func convertMdToHTML(dat []byte) {
	output = blackfriday.Run(dat)
	createFolder()
}

func readFiles(fullFilePath string) {
	dat, err := ioutil.ReadFile(fullFilePath)
	check(err)
	convertMdToHTML(dat)
}

func compile() {
	files, err := ioutil.ReadDir("./markdown")
	check(err)

	for _, file := range files {
		fileName = file.Name()
		fullFilePath := "./markdown/" + fileName
		readFiles(fullFilePath)
	}
}

func main() {
	compile()
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	router.Run(":3000")
}
