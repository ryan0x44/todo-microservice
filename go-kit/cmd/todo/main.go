package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	// Load config from env
	var c config
	err := envconfig.Process("", &c)
	if err != nil {
		fmt.Println("Error parsing config", err)
		return
	}
	// Connect to DB
	db, err := c.MySQL.DB()
	if err != nil {
		fmt.Println("Error load DB driver")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to DB: ", err)
	}
	defer db.Close()
	fmt.Println("Hello World")
}
