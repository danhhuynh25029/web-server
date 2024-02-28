package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Port      string `json:"port"`
	TargetUrl string `json:"target_url"`
}

var AllConfig Config

func Load() {
	jsonFile, err := os.Open("config.json")
	defer func() {
		fmt.Println("jsonFile is closed")
		jsonFile.Close()
	}()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("success fully open config")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &AllConfig)
}
