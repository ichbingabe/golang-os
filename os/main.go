package main

import (
	"log"
	"os"
	"os/exec"
	"fmt"
	"encoding/csv"
)

var Name string
var Url string

type UserArray struct{
	User []User
}

func main() {
	Name = "testfile1"
	Url = "a"
	//if !os.IsExist()
//	Mkdir(Name)
//	PythonScript()
	user1 := User{
		Name: "",
		LastName: "asdf",
		Email: "email",
		Id: 1,
	}

	user2 := User{
		Name: "atrwtre",
		LastName: "gfdsgfds",
		Email: "fdagfa",
		Id: 2,
   }

   user3 := User{
		Name: "afds",
		LastName: "lafdass",
		Email: "emafdail",
		Id: 3,
	}

	UserArr := []User{user1, user2, user3}
	usr := UserArray{UserArr}

	//fmt.Printf("User %v\\n", usr)
	_, b := marshal(usr) 
	err := os.WriteFile("testdir/output.json", []byte(b), 0660)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("bash", "bash/test.sh")
	if err := cmd.Run(); err != nil {
		fmt.Println("error: ", err)
	}

	cmd = exec.Command("python3", "bash/test.py")
	if err := cmd.Run(); err != nil {
		fmt.Println("error: ", err)
	}


	file, err := os.Create("testdir/output.csv")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, user := range UserArr {
		row := []string{user.Email, user.Name, user.LastName}
		if err := w.Write(row); err != nil{
			log.Fatal(err)
		}
	}
}
