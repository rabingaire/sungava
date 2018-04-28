package main

import (
	"io/ioutil"
	"os"
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

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		check(err)
	}
}

func writeToFile(output []byte) {
	outputFileName := strings.Split(fileName, ".")[0]
	if outputFileName == "index" {
		err := ioutil.WriteFile("./views/"+outputFileName+".html", output, 0644)
		check(err)
	} else {
		dir := "./views/" + outputFileName
		createDirIfNotExist(dir)
		err := ioutil.WriteFile(dir+"/index.html", output, 0644)
		check(err)
	}
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
