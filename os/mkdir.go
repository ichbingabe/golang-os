package main

import (
	"os"
	"log"
)

func Mkdir(name string) {
	err := os.Mkdir(name, 0750)
	if err != nil && os.IsExist(err){
		log.Fatal(err)
	}
}

func PythonScript(name string) {
	path := name + "/test.py"
	str := 
	`def test():
	abc = "testing"
	print(abc)`
	err := os.WriteFile(path, []byte(str), 0660)
	if err != nil {
		log.Fatal(err)
	} 
}