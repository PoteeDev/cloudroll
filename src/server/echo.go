package server

import (
	"context"

	proto "github.com/PoteeDev/cloudroll/proto"
)

// Implements of EchoServiceServer

type echoServer struct {
	proto.UnimplementedCloudrollServiceServer
}

func newCloudrollServer() proto.CloudrollServiceServer {
	return new(echoServer)
}

func (s *echoServer) Ping(ctx context.Context, req *proto.Empty) (*proto.EchoMessage, error) {

	return &proto.EchoMessage{
		Value: "pong",
	}, nil
}
