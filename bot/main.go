package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Command struct {
	id string
	cmd string
	time int
	signature string
}

var commander_url = "http://127.1:3000/command"

func pollCommand() {
	resp, err := http.Get(commander_url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func main() {
	
}