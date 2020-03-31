package main

import (
	"encoding/json"
	config2 "go_chaos/config"
	"io/ioutil"
	"os"
)

func main() {
	fileLocation := getFileLocationFromArgs()
	// read bytes from file
	data, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		panic(err)
	}
	
	config := &config2.ChaosRouterConfig{}
	if err := json.Unmarshal(data, config); err != nil {
		panic(err)
	}
	
	chaosRouter := config2.CreateChaosRouterFromConfig(*config)
	chaosRouter.Start()
}

func getFileLocationFromArgs() string {
	args := os.Args
	if len(args) < 2 {
		panic("config file location not given as program argument")
	}
	return args[1]
}
