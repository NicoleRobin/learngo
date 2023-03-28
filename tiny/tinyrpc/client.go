package main

import (
	"github.com/nicolerobin/tinyrpc/codec"
	"github.com/nicolerobin/tinyrpc/compressor"
	"github.com/nicolerobin/tinyrpc/serializer"
	"io"
	"net/rpc"
)

type Client struct {
	*rpc.Client
}

// Option provides options for rpc
type Option func(o *options)

type options struct {
	compressType compressor.CompressType
	serializer   serializer.Serializer
}

// WithCompress set client compression format
func WithCompress(c compressor.CompressType) Option {
	return func(o *options) {
		o.compressType = c
	}
}

// WithSerializer set client serializer
func WithSerializer(s serializer.Serializer) Option {
	return func(o *options) {
		o.serializer = s
	}
}

// NewClient
func NewClient(conn io.ReadWriteCloser, opts ...Option) *Client {
	options := options{
		compressType: compressor.Raw,
		serializer:   serializer.Proto,
	}

	for _, opt := range opts {
		opt(&options)
	}

	return &Client{
		rpc.NewClientWithCodec(codec.NewClientCodec(conn, options.compressType, options.serializer)),
	}
}

// Call
func (c *Client) Call(serviceMethod string, args interface{}, reply interface{}) error {
	return c.Client.Call(serviceMethod, args, reply)
}

// AsyncCall
func (c *Client) AsyncCall(serviceMethod, args interface{}, reply interface{}) chan *rpc.Call {
	return c.Go(serviceMethod, args, reply, nil).Done
}
