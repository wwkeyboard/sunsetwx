package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

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
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://sunburst.sunsetwx.com/v1/quality", nil)
	if err != nil {
		fmt.Println("could make request", err)
		os.Exit(1)
	}
	token := "foo"
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error querying api", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", body)
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

	return &Config{}, nil
}
