package main

import (
	"crontab/client/client_wrapper"
	"log"
)

func main() {
	connection := &client_wrapper.ConnectionWrapper{}
	connection.Connect()
}

func resolve(r client_wrapper.Resolver, task client_wrapper.TaskInfo) {
	r.Resolve("test", func(task client_wrapper.TaskInfo) {
		log.Print(task)
	})
}
