package main

import (
	"encoding/json"
	"fmt"
	config2 "go_chaos/config"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	log.Print("Starting server")
	fileLocation := getFileLocationFromArgs()
	log.Print(fmt.Sprintf("config file location: %s", fileLocation))
	// read bytes from file
	data, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		panic(err)
	}
	
	config := &config2.ChaosRouterConfig{}
	if err := json.Unmarshal(data, config); err != nil {
		panic(err)
	}
	log.Print("successfully loaded config")
	
	chaosRouter := config2.CreateChaosRouterFromConfig(*config)
	log.Print(fmt.Sprintf("starting chaose router on port: %d", config.Port))
	chaosRouter.Start()
}

func getFileLocationFromArgs() string {
	args := os.Args
	if len(args) < 2 {
		panic("config file location not given as program argument")
	}
	return args[1]
}
