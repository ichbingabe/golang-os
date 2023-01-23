package main

import (
	"os"
	"log"
)

var Path string

func Mkdir(name string) {
	err := os.Mkdir(name, 0750)
	if err != nil && os.IsExist(err){
		log.Fatal(err)
	}
}

func PythonScript() {
	Path = Name + "/test.py"
	str := 
`def test():
	abc = "testing"
	print(abc)

test()
abc = "test2 funcionando perfeitamente"
def test2(abc):
	res = abc
	print(res)
	
test2(abc)`
	
err := os.WriteFile(Path, []byte(str), 0660)
	if err != nil {
		log.Fatal(err)
	} 
}