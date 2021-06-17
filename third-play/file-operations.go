package main

import (
	"errors"
	"fmt"
	"os"
)

func main(){
	if checkFileExiss("/Users/i323691/work_dir/training/go-training/golang-playbook/third-play/map-string-interface-cars.json") {
		fmt.Println("This file is exists")
	}else {
		fmt.Println("The file is not exists and it is a directory")
	}

}

func checkFileExiss(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil{
		errors.New("Errror happened during openning the target file path")
	}
	if fileInfo.IsDir() == true {
		fmt.Println("It is a directory and it is not a file")
		return false
	}

	//https://golang.org/pkg/os/#IsNotExist
	if os.IsNotExist(err) {
		fmt.Println("file does not exist")
		return false
	}else {
		fmt.Println("Found the target file...")
		fmt.Printf("The size of the target file is: %d\n",fileInfo.Size())
		return true
	}

}