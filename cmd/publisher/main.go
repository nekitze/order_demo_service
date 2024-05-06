package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	for {
		var input string
		fmt.Println("Path to .json file:")
		fmt.Scanln(&input)

		if input == "q" {
			fmt.Println("Goodbye!")
			break
		}

		bytes, err := os.ReadFile(input)
		if err != nil {
			log.Printf("Wrong file or path: %s\n", err)
			continue
		}

		err = nc.Publish("orders.new", bytes)
		if err != nil {
			log.Printf("Publish error: %s\n", err)
			continue
		}

		fmt.Println("Published!")
	}
}
