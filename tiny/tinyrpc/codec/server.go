package codec

import (
	"bufio"
	"io"
	"net/rpc"
	"sync"

	"github.com/nicolerobin/tinyrpc/header"
	"github.com/nicolerobin/tinyrpc/serializer"
)

type serverCodec struct {
	r io.Reader
	w io.Writer
	c io.Closer

	request    header.ResponseHeader
	serializer serializer.Serializer
	mutex      sync.Mutex
	seq        uint64
	pending    map[uint64]uint64
}

// NewServerCodec Create a new server codec
func NewServerCodec(conn io.ReadWriteCloser, serializer serializer.Serializer) rpc.ServerCodec {
	return &serverCodec{
		r:          bufio.NewReader(conn),
		w:          bufio.NewWriter(conn),
		c:          conn,
		serializer: serializer,
		pending:    make(map[uint64]uint64),
	}
}

// ReadRequestHeader
func (s *serverCodec) ReadRequestHeader(*Request) error {
	return nil
}

// ReadRequestBody
func (s *serverCodec) ReadRequestBody(any) error {
	return nil
}

// WriteResponse
func (s *serverCodec) WriteResponse(*Response, any) error {
	return nil
}

func (s *serverCodec) Close() error {
	return nil
}
