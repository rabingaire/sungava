package main

import (
	"io/ioutil"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

var fileName string

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func writeToFile(output []byte) {
	outputFileName := strings.Split(fileName, ".")[0] + ".html"
	err := ioutil.WriteFile("./views/"+outputFileName, output, 0644)
	check(err)
}

func convertMdToHTML(dat []byte) {
	output := blackfriday.Run(dat)
	writeToFile(output)
}

func readFiles(fullFilePath string) {
	dat, err := ioutil.ReadFile(fullFilePath)
	check(err)
	convertMdToHTML(dat)
}

func main() {
	files, err := ioutil.ReadDir("./markdown")
	check(err)

	for _, file := range files {
		fileName = file.Name()
		fullFilePath := "./markdown/" + fileName
		readFiles(fullFilePath)
	}

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	router.Run(":3000")
}
