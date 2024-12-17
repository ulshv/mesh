package main

import (
	"log"

	mesh "github.com/ulshv/mesh/v1"
)

func main() {
	mesh.Test()
	node, err := mesh.MakeNode()
	if err != nil {
		log.Fatal(err)
	}
	mesh.LogWithJson("created a node from the app: %v", node)
}
