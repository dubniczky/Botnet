package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Command struct {
	Id string
	Cmd string
	Time int
	Signature string
}

var commander_url = "http://127.1:3000/command"

func pollCommand() (Command, error) {
	resp, err := http.Get(commander_url)
	if err != nil {
		return Command{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Command{}, err
	}

	var command Command
	err = json.Unmarshal(body, &command)
	if err != nil {
		return Command{}, err
	}

	return command, nil
}

func main() {
	command, err := pollCommand()
	if err != nil {
		log.Println(err)
		return
	}
	
	log.Println(command.Cmd)
}