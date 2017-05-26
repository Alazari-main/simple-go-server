package main

import (
	"net/http"
	"path/filepath"
	"server"
	"server/config"
	"flag"
	"log"
	"fmt"
)

func main() {

	var wd string
	var port int

	flag.StringVar(
		&wd,
		"wd", 
		"./", 
		"wd flag sets the working directory for serving. Default value is current execution directory.")

	flag.IntVar(
		&port,
		"port", 
		9090, 
		"port flag sets the http port which would be used to listening for incoming requests. Default value is 9090.")

	flag.Parse()

	wd, err := filepath.Abs(wd)
	if err != nil {
		log.Fatal(err.Error())
	}

	config.Instance.WorkingDirectory = http.Dir(wd)
	fmt.Printf("Working directory is set to '%s' \n", wd)

	config.Instance.Port = port
	fmt.Printf("Port is set to '%d' \n", port)

	srv := &server.Server{}

	fmt.Printf("Server will be serving at http://127.0.0.1:%d", port)
	srv.Start()

}
