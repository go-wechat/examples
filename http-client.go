package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
)

var baseUrl = "http://httpbin.org"

func main() {
	fmt.Println("vim-go")

	var url = "/"
	var reqUrl = baseUrl + url

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {

	}
}

func init() {
	log.SetPrefix("http-client")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags)
}
