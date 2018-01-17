package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Configuration struct {
	LogFileName string
}

func readConfig(configFile string) (Configuration, error) {
	//filename is the path to the json config file
	file, err := os.Open(configFile)
	if err != nil {
		return Configuration{}, err
	}
	configuration := Configuration{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		return Configuration{}, err
	}

	return configuration, nil
}

func setLogger(name string) {
	//creates a new log file with name as parameter

	fileHandle, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(io.Writer(fileHandle))

}

func main() {

	config, err := readConfig("config.json")
	if err != nil {
		setLogger(config.LogFileName)
	}

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
