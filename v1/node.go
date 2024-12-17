package mesh_v1

import (
	"github.com/gofrs/uuid" // UUID v6 generation library
)

// Node represents a basic mesh node with a unique UUID
type Node struct {
	UUID string `json:"id"`
}

func (n *Node) String() (string, error) {
	str, err := ToJsonStr(n)
	return str, err
}

// MakeNode creates a new Node with a UUID v6 as the ID
func MakeNode() (*Node, error) {
	// Generate a UUID v6
	uuidV6, err := uuid.NewV6()
	if err != nil {
		return nil, err
	}
	// Create a new Node instance with the generated UUID
	node := &Node{
		UUID: uuidV6.String(),
	}
	return node, nil
}
