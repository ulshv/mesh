package mesh_v1

import (
	"sync"

	"github.com/gofrs/uuid" // UUID v6 generation library
)

var nodes struct {
	rwMtx sync.RWMutex
	data  map[string]*Node
}

// Initialize the nodes object in the init function
func init() {
	nodes = struct {
		rwMtx sync.RWMutex
		data  map[string]*Node
	}{
		rwMtx: sync.RWMutex{},
		data:  make(map[string]*Node),
	}
}

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
	AddNode(node)
	return node, nil
}

// AddNode adds a new node to the map
func AddNode(node *Node) {
	nodes.rwMtx.Lock()
	defer nodes.rwMtx.Unlock()
	nodes.data[node.UUID] = node
	Log("AddNode(): node added: UUID=", node.UUID)
}

// getNode retrieves a node by its UUID
func GetNode(uuid string) (*Node, bool) {
	nodes.rwMtx.RLock()
	defer nodes.rwMtx.RUnlock()
	node, ok := nodes.data[uuid]
	return node, ok
}

// listNodes lists all nodes in the map
func ListNodes() {
	nodes.rwMtx.RLock()
	defer nodes.rwMtx.RUnlock()
	Log("All nodes:")
	for uuid, node := range nodes.data {
		Log("UUID: %s, Node: %+v\n", uuid, node)
	}
}

// func (n *Node) SendMessage(toNode *Node) error {

// }
