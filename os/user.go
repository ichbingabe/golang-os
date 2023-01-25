package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	Name string  `json:"name,omitempty"`
	LastName string `json:"lastname,omitempty"`
	Email string `json:"email,omitempty"`
	Id uint `json:"id,omitempty"`
} 

func marshal(usr UserArray) (error, []byte){
	user := usr

	b, err := json.Marshal(user)
		
	if err != nil {
		fmt.Println(err)
		return err, b
	}


	fmt.Println(string(b))

	return err, b
	
}
