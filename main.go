package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Println("here!")
	token := os.Getenv("SUNSETWX_AUTH_TOKEN")
	if len(token) < 10 {
		fmt.Println("you must set the SunsetWX token in SUNSETWX_AUTH_TOKEN")
		os.Exit(1)
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://sunburst.sunsetwx.com/v1/quality", nil)
	if err != nil {
		fmt.Println("could make request", err)
		os.Exit(1)
	}
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
