package main

import (
	"fmt"
	"log"
	"net/http"
	"web-server/config"
	"web-server/proxy"
)

func init() {
	config.Load()
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/proxy", proxy.ForwardRequest)
	log.Fatal(http.ListenAndServe(":"+config.AllConfig.Port, nil))
}
