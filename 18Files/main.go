package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in GoLang")

	fileName := "./myTextIOFile.txt"
	content := "This needs to go in file rajathio.in"

	file, err := os.Create(fileName)

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Length is : ", length)

	defer file.Close()

	readFile(fileName)
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)

	checkNilErr(err)

	fmt.Println("File content is : \n", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
