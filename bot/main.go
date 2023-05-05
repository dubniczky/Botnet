package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Command struct {
	Id string
	Cmd string
	Time int
	Signature string
}

var commander_url = "http://127.1:3000/command"
var bot_name = "testbot"
var pollInterval uint32 = 10

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

func execute() error {
	command, err := pollCommand()
	if err != nil {
		log.Println(err)
		return err
	}
	
	log.Println(command.Cmd)
	return nil
}

func main() {
	log.Printf("Starting polling cycle... (%ds)\n", pollInterval)
	for range time.Tick(time.Second * time.Duration(pollInterval)) {
        execute()
    }
}