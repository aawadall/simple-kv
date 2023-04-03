package main

import (
	kvserver "github.com/aawadall/simple-kv/kv_server"
)

func main() {
	// TODO - Entry point code  here
	// define server
	config := map[string]string{
		"driver":        "none",
		"file_location": "log.txt",
		"sync_interval": "120",
	}
	server := kvserver.NewKVServer(config)

	server.Start()
}
