package main

import (
	"fmt"
	"log"

	mesh "github.com/ulshv/mesh/v1"
)

type testWriter struct{}

func (tw testWriter) Write(data []byte) (int, error) {
	fmt.Println(string(data))
	return 0, nil
}

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

	mesh.SendMessageByUUIDs(node1.UUID, node2.UUID, []byte("hello, world!"))

	tw1 := testWriter{}

	_, err = mesh.TestOne(tw1)

	if err != nil {
		log.Fatal(err)
	}
}
