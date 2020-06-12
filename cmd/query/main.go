package main

import (
	"fmt"
	"os"
	"strings"
)

// Config of the application
type Config struct {
	username string
	password string
}

func main() {
	config, err := GetConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("%#v", config)
}

// GetConfig for this run from the environment
func GetConfig() (*Config, error) {
	var errors []string
	username := os.Getenv("SUNSETWX_USERNAME")
	if len(username) == 0 {
		errors = append(errors, "you must set env var SUNSETWK_USERNAME")
	}

	password := os.Getenv("SUNSETWX_PASSWORD")
	if len(password) == 0 {
		errors = append(errors, "you must set env var SUNSETWX_PASSWORD")
	}
	if len(errors) > 0 {
		return nil, fmt.Errorf(strings.Join(errors, "\n"))
	}

	return &Config{
		username: username,
		password: password,
	}, nil
}
