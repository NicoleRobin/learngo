package main

import (
	"github.com/nicolerobin/tinyrpc/codec"
	"github.com/nicolerobin/tinyrpc/serializer"
	"log"
	"net"
	"net/rpc"
)

type Server struct {
	*rpc.Server
	serializer.Serializer
}

func (s *Server) Serve(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Print("tinyrpc.Serve: accept:", err.Error())
			return
		}
		// one conn one goroutine
		go s.Server.ServeCodec(codec.NewServerCodec(conn, s.Serializer))
	}
}
