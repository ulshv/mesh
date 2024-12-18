package mesh_v1

import "io"

func TestOne(w io.Writer) (int, error) {
	n, err := w.Write([]byte("hello, test!"))
	return n, err
}
