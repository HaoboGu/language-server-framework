package server

import "io"

// Connection is the bidirectional connection used by jsonrpc connection
type Connection struct {
	in  io.ReadCloser
	out io.WriteCloser
}

func (c *Connection) Read(p []byte) (n int, err error) {
	return c.in.Read(p)
}

func (c *Connection) Write(p []byte) (n int, err error) {
	return c.out.Write(p)
}

// Close closes the connection
func (c *Connection) Close() error {
	if err := c.in.Close(); err != nil {
		return err
	}
	return c.out.Close()
}
