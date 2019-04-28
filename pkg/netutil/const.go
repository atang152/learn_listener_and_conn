package netutil

// By default bufio.Writer uses 4096 bytes long buffer as per
// https://github.com/golang/go/blob/c6c0f47e92771c9b4fced87b94c04f66e5d6eba5/src/bufio/bufio.go#L18
const (
	defaultBufSize = 4096
)
