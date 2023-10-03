package main

import "github.com/LucasBelusso1/GoExpert/9-Events/pkg/rabbitmq"

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()
	rabbitmq.Publish(ch, "Hello World!", "amq.direct")
}
