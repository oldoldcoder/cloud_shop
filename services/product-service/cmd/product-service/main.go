package main

import (
    "fmt"
    _ "github.com/segmentio/kafka-go"
    _ "github.com/rabbitmq/amqp091-go"
)

func main() {
    fmt.Println("product-service started")
}


