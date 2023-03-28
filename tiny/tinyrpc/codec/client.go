package codec

import (
	"bufio"
	"io"
	"net/rpc"
	"sync"

	"github.com/nicolerobin/tinyrpc/compressor"
	"github.com/nicolerobin/tinyrpc/header"
	"github.com/nicolerobin/tinyrpc/serializer"
)

type clientCodec struct {
	r io.Reader
	w io.Writer
	c io.Closer

	compressor compressor.CompressType
	serializer serializer.Serializer
	response   header.ResponseHeader
	mutex      sync.Mutex
	pending    map[uint64]string
}

// NewClientCodec Create a new client codec
func NewClientCodec(conn io.ReadWriteCloser, compressType compressor.CompressType,
	serializer serializer.Serializer) rpc.ClientCodec {
	return &clientCodec{
		r:          bufio.NewReader(conn),
		w:          bufio.NewWriter(conn),
		c:          conn,
		compressor: compressType,
		serializer: serializer,
		pending:    make(map[uint64]string),
	}
}

func (c *clientCodec) WriteRequest(r *rpc.Request, param any) error {
	return nil
}

func (c *clientCodec) ReadResponseHeader(r *rpc.Response) error {
	return nil
}

func (c *clientCodec) ReadResponseBody(param any) error {
	return nil
}

func (c *clientCodec) Close() error {
	return nil
}
