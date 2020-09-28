package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

//Config считывает env переменные из файла .vscode\launch.json
type Config struct {
	Port          string
	Secrets       map[string]string
	DBCredentials map[string]string
}

func main() {
	var s Config
	err := envconfig.Process("glebtest", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Port: %s\n", s.Port)
	fmt.Println("Secrets:")
	for k, v := range s.Secrets {
		fmt.Printf("  %s: %s\n", k, v)
	}
	fmt.Println("DBCredentials:")
	for k, v := range s.DBCredentials {
		fmt.Printf("  %s: %s\n", k, v)
	}
}
