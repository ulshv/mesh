package main

import (
	"log"

	mesh "github.com/ulshv/mesh/v1"
)

func main() {
	mesh.Test()
	node1, err := mesh.MakeNode()
	if err != nil {
		log.Fatal(err)
	}
	node2, err := mesh.MakeNode()
	if err != nil {
		log.Fatal(err)
	}
	node3, err := mesh.MakeNode()
	if err != nil {
		log.Fatal(err)
	}
	mesh.Log("created nodes:", []*mesh.Node{node1, node2, node3})
}
