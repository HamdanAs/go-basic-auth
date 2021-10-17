package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8090", nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth("hamdan", "hamdan21")
	req.Header.Add("Content-Type", "application/json")
	req.Close = true

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if response.StatusCode != http.StatusOK {
		panic("Non 2xx response from server, request" + response.Status)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	log.Print(string(body))
}
