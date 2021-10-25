package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/jmkao/twitch-channel-tool/pkg/core"
)

func main() {
	file, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	// Read configuration
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	config := core.Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()

	core.StartServer(config)
}
