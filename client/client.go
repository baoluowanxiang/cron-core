package main

import (
	"crontab/client/client_wrapper"
	"log"
)

func main() {
	connection := &client_wrapper.ConnectionWrapper{resolve}
	connection.Connect()
}

func resolve(r *client_wrapper.Route) {
	r.Put("test", Test)
	r.Put("example", Example)
	r.Put("example1", Example)
}

func Test(task client_wrapper.TaskInfo) {
	log.Print(task)
}

func Example(task client_wrapper.TaskInfo) {
	log.Print("example: ", task.Params)
}
