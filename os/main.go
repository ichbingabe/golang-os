package main

import (
	"log"
	"os"
)

func main() {
	name := "testfile1"
	Mkdir(name)
	PythonScript(name)
	err := os.WriteFile("testdir/testfile1.csv", []byte("Hello, World!"), 0660)
	if err != nil {
		log.Fatal(err)
	}
}
