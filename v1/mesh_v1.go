package mesh_v1

import "fmt"

func log(msg string) {
	fmt.Printf("[mesh_v1]: %s\n", msg)
}

func init() {
	log("init() is called.")
}

// not called:
// func main() {
// 	log("main() is called.")
// }

func Test() {
	log("Test() is called")
}
