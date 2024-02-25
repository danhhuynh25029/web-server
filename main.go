package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port string `json:"port"`
}

func main() {

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

	var config Config

	json.Unmarshal(byteValue, &config)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
