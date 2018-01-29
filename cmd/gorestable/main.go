package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Configuration struct {
	LogFileName string `json:"logfile"`
}

//readConfig reads the configuration file passed as parameter.
//Please note that the configFile parameter is the complete file path.
func readConfig(configFile string) (Configuration, error) {

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

//setLogger function creates a new logger using the file passed as parameter as logfile.
//Please note that the name parameter is the complete file path.
func setLogger(name string) {

	fileHandle, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(io.Writer(fileHandle))

}

func main() {

	config, err := readConfig("config.json")
	if err != nil {
		//Error management
	}

	setLogger(config.LogFileName)
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
