package main

import (
	"log"
	"os"
//	"os/exec"
//	"fmt"
)

var Name string
var Url string

func main() {
	Name = "testfile1"
	Url = "a"
	//if !os.IsExist()
//	Mkdir(Name)
//	PythonScript()
	var id uint 

	name := "a"
	lastNAme := ""
	email := ""
	id = 0

	_, b := Marshal(name, lastNAme, email, id) 
	err := os.WriteFile("testdir/output.json", []byte(b), 0660)
	if err != nil {
		log.Fatal(err)
	}

//	c := exec.Command(Path)
//	if err := c.Run(); err != nil {
//		fmt.Println("error: ", err)
//	}
}
