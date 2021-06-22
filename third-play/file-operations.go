package main

import (
	"errors"
	"fmt"
	"os"
)

// Mkdir permission table
// +-----+---+--------------------------+
// | rwx | 7 | Read, write and execute  |
// | rw- | 6 | Read, write              |
// | r-x | 5 | Read, and execute        |
// | r-- | 4 | Read,                    |
// | -wx | 3 | Write and execute        |
// | -w- | 2 | Write                    |
// | --x | 1 | Execute                  |
// | --- | 0 | no permissions           |
// +------------------------------------+

// +------------+------+-------+
// | Permission | Octal| Field |
// +------------+------+-------+
// | rwx------  | 0700 | User  |
// | ---rwx---  | 0070 | Group |
// | ------rwx  | 0007 | Other |
// +------------+------+-------+

// 0755 Commonly used on web servers. The owner can read, write, execute. 
//  Everyone else can read and execute but not modify the file.

// 0777 Everyone can read write and execute. On a web server, 
// 	it is not advisable to use ‘777’ permission for your files and folders, as it allows anyone to add malicious code to your server.

// 0644 Only the owner can read and write. Everyone else can only read. No one can execute the file.

// 0655 Only the owner can read and write, but not execute the file. Everyone else can read and execute, 
//	but cannot modify the file.

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