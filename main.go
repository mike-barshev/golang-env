package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)


type Student struct {
	port     string
	secrets int `json:"Standard"`
	db_credentials
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var studentDecoder *json.Decoder = json.NewDecoder(file)
	if err != nil {
		log.Fatal(err)
	}

	var studentList []Student

	err = studentDecoder.Decode(&studentList)
	if err != nil {
		log.Fatal(err)
	}

	for i, student := range studentList {
		fmt.Println("Student", i+1)
		fmt.Println("Student name:", student.Name)
		fmt.Println("Student standard:", student.Standard)
	}
}
