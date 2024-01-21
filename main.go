package main

import (
	"a2s/pkg/client"
	"fmt"
)

func main() {
	var client, error = client.CreateClient("69.10.61.11", 28020)

	defer client.Close()

	fmt.Println(client, error)
	fmt.Println(client.Read())
}
